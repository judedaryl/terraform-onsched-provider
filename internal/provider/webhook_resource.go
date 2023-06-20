package provider

import (
	"context"
	"fmt"
	"terraform-provider-onsched/onsched"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type webhookResource struct {
	client *onsched.Client
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &webhookResource{}
	_ resource.ResourceWithConfigure = &webhookResource{}
)

// NewWebhookResource is a helper function to simplify the provider implementation.
func NewWebhookResource() resource.Resource {
	return &webhookResource{}
}

// Configure adds the provider configured client to the resource.
func (r *webhookResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*onsched.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *onsched.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

// Metadata returns the resource type name.
func (r *webhookResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

// Schema defines the schema for the resource.
func (r *webhookResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Example resource",

		Attributes: map[string]schema.Attribute{
			"booking_webhook_url": schema.StringAttribute{
				MarkdownDescription: "Webhook called when a booking event occurs.",
				Default:             stringdefault.StaticString("SOFT_DELETED"),
				Computed:            true,
				Optional:            true,
			},
			"customer_webhook_url": schema.StringAttribute{
				MarkdownDescription: "Webhook called when a customer event occurs.",
				Default:             stringdefault.StaticString("SOFT_DELETED"),
				Computed:            true,
				Optional:            true,
			},
			"resource_webhook_url": schema.StringAttribute{
				MarkdownDescription: "Webhook called when a resource event occurs.",
				Default:             stringdefault.StaticString("SOFT_DELETED"),
				Computed:            true,
				Optional:            true,
			},
			"reminder_webhook_url": schema.StringAttribute{
				MarkdownDescription: "Webhook called when a reminder event occurs.",
				Default:             stringdefault.StaticString("SOFT_DELETED"),
				Computed:            true,
				Optional:            true,
			},
			"webhook_signature_hash": schema.StringAttribute{
				MarkdownDescription: "Webhook signature hash",
				Default:             stringdefault.StaticString(""),
				Computed:            true,
				Optional:            true,
			},
			"disable_email_and_sms_notifications": schema.BoolAttribute{
				MarkdownDescription: "This will disable all email and sms notifications, webhooks will still be triggered",
				Default:             booldefault.StaticBool(false),
				Optional:            true,
				Computed:            true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *webhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan webhookResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	company, err := r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched webhook",
			err.Error(),
		)
		return
	}

	company.BookingWebhookURL = plan.BookingWebhookURL.ValueString()
	company.CustomerWebhookURL = plan.CustomerWebhookURL.ValueString()
	company.ReminderWebhookURL = plan.ReminderWebhookURL.ValueString()
	company.ResourceWebhookURL = plan.ResourceWebhookURL.ValueString()
	company.WebhookSignatureHash = plan.WebhookSignatureHash.ValueString()
	company.DisableEmailAndSmsNotifications = plan.DisableEmailAndSmsNotifications.ValueBool()

	_, err = r.client.UpdateCompany(company)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched webhook",
			err.Error(),
		)
		return
	}

	company, err = r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched webhook",
			err.Error(),
		)
		return
	}

	plan.BookingWebhookURL = types.StringValue(company.BookingWebhookURL)
	plan.CustomerWebhookURL = types.StringValue(company.CustomerWebhookURL)
	plan.ReminderWebhookURL = types.StringValue(company.ReminderWebhookURL)
	plan.ResourceWebhookURL = types.StringValue(company.ResourceWebhookURL)
	plan.WebhookSignatureHash = types.StringValue(company.WebhookSignatureHash)
	plan.DisableEmailAndSmsNotifications = types.BoolValue(company.DisableEmailAndSmsNotifications)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *webhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state webhookResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading OnSched webhook",
			err.Error(),
		)
		return
	}

	state.BookingWebhookURL = types.StringValue(c.BookingWebhookURL)
	state.CustomerWebhookURL = types.StringValue(c.CustomerWebhookURL)
	state.ReminderWebhookURL = types.StringValue(c.ReminderWebhookURL)
	state.ResourceWebhookURL = types.StringValue(c.ResourceWebhookURL)
	state.WebhookSignatureHash = types.StringValue(c.WebhookSignatureHash)
	state.DisableEmailAndSmsNotifications = types.BoolValue(c.DisableEmailAndSmsNotifications)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}

// Update updates the resource and sets the updated Terraform state on success.
func (r *webhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan webhookResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	company, err := r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched webhook",
			err.Error(),
		)
		return
	}

	company.BookingWebhookURL = plan.BookingWebhookURL.ValueString()
	company.CustomerWebhookURL = plan.CustomerWebhookURL.ValueString()
	company.ReminderWebhookURL = plan.ReminderWebhookURL.ValueString()
	company.ResourceWebhookURL = plan.ResourceWebhookURL.ValueString()
	company.WebhookSignatureHash = plan.WebhookSignatureHash.ValueString()
	company.DisableEmailAndSmsNotifications = plan.DisableEmailAndSmsNotifications.ValueBool()

	_, err = r.client.UpdateCompany(company)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched webhook",
			err.Error(),
		)
		return
	}

	company, err = r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched webhook",
			err.Error(),
		)
		return
	}

	plan.BookingWebhookURL = types.StringValue(company.BookingWebhookURL)
	plan.CustomerWebhookURL = types.StringValue(company.CustomerWebhookURL)
	plan.ReminderWebhookURL = types.StringValue(company.ReminderWebhookURL)
	plan.ResourceWebhookURL = types.StringValue(company.ResourceWebhookURL)
	plan.WebhookSignatureHash = types.StringValue(company.WebhookSignatureHash)
	plan.DisableEmailAndSmsNotifications = types.BoolValue(company.DisableEmailAndSmsNotifications)
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *webhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state webhookResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	company, err := r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting OnSched webhook",
			err.Error(),
		)
		return
	}

	company.BookingWebhookURL = "SOFT_DELETED"
	company.CustomerWebhookURL = "SOFT_DELETED"
	company.ReminderWebhookURL = "SOFT_DELETED"
	company.ResourceWebhookURL = "SOFT_DELETED"

	_, err = r.client.UpdateCompany(company)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting OnSched webhook",
			err.Error(),
		)
		return
	}

	company, err = r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched webhook",
			err.Error(),
		)
		return
	}

	state.BookingWebhookURL = types.StringValue(company.BookingWebhookURL)
	state.CustomerWebhookURL = types.StringValue(company.CustomerWebhookURL)
	state.ReminderWebhookURL = types.StringValue(company.ReminderWebhookURL)
	state.ResourceWebhookURL = types.StringValue(company.ResourceWebhookURL)
	state.WebhookSignatureHash = types.StringValue(company.WebhookSignatureHash)
	state.DisableEmailAndSmsNotifications = types.BoolValue(company.DisableEmailAndSmsNotifications)
	state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

type webhookResourceModel struct {
	BookingWebhookURL               types.String `tfsdk:"booking_webhook_url"`
	CustomerWebhookURL              types.String `tfsdk:"customer_webhook_url"`
	ReminderWebhookURL              types.String `tfsdk:"reminder_webhook_url"`
	ResourceWebhookURL              types.String `tfsdk:"resource_webhook_url"`
	WebhookSignatureHash            types.String `tfsdk:"webhook_signature_hash"`
	DisableEmailAndSmsNotifications types.Bool   `tfsdk:"disable_email_and_sms_notifications"`
	LastUpdated                     types.String `tfsdk:"last_updated"`
}
