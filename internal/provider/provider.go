package provider

import (
	"context"
	"os"
	"terraform-provider-onsched/onsched"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &OnSchedProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &OnSchedProvider{
			version: version,
		}
	}
}

type OnSchedProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

type Environment string

const (
	Sandbox Environment = "sandbox"
	Prod    Environment = "prod"
)

type onschedProviderModel struct {
	Env types.String `tfsdk:"env"`
}

// Metadata returns the provider type name.
func (p *OnSchedProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "onsched"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *OnSchedProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"env": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf([]string{"sandbox", "prod"}...),
				},
			},
		},
	}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *OnSchedProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config onschedProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var env onsched.Environment

	if config.Env.IsNull() || config.Env.ValueString() == string(Sandbox) {
		env = onsched.Sandbox
	} else {
		env = onsched.Prod
	}

	client_id := os.Getenv("ONSCHED_CLIENT_ID")
	client_secret := os.Getenv("ONSCHED_CLIENT_SECRET")

	if client_id == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_id"),
			"Missing ONSCHED_CLIENT_ID",
			"Set the ONSCHED_CLIENT_ID environment variable.",
		)
	}

	if client_secret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("client_secret"),
			"Missing ONSCHED_CLIENT_SECRET",
			"Set the ONSCHED_CLIENT_SECRET environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Creating OnSched client")
	client := onsched.NewClient(env, client_id, client_secret)

	resp.ResourceData = client
	tflog.Info(ctx, "Configured OnSched client")
}

// DataSources defines the data sources implemented in the provider.
func (p *OnSchedProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *OnSchedProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewWebhookResource,
	}
}
