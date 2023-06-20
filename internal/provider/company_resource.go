package provider

import (
	"context"
	"fmt"
	"terraform-provider-onsched/onsched"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type companyResource struct {
	client *onsched.Client
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &companyResource{}
	_ resource.ResourceWithConfigure = &companyResource{}
)

// NewcompanyResource is a helper function to simplify the provider implementation.
func NewCompanyResource() resource.Resource {
	return &companyResource{}
}

// Configure adds the provider configured client to the resource.
func (r *companyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *companyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_company"
}

// Schema defines the schema for the resource.
func (r *companyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"registrationDate": schema.StringAttribute{
				Computed: true,
			},
			"registrationEmail": schema.StringAttribute{
				Computed: true,
			},
			"addressLine1": schema.StringAttribute{
				Required: true,
			},
			"addressLine2": schema.StringAttribute{
				Optional: true,
			},
			"city": schema.StringAttribute{
				Optional: true,
			},
			"state": schema.StringAttribute{
				Optional: true,
			},
			"postalCode": schema.StringAttribute{
				Optional: true,
			},
			"country": schema.StringAttribute{
				Optional: true,
			},
			"phone": schema.StringAttribute{
				Optional: true,
			},
			"fax": schema.StringAttribute{
				Optional: true,
			},
			"email": schema.StringAttribute{
				Required: true,
			},
			"website": schema.StringAttribute{
				Optional: true,
			},
			"timezoneId": schema.StringAttribute{
				Required: true,
			},
			"timezoneName": schema.StringAttribute{
				Required: true,
			},
			"bookingWebhookUrl": schema.StringAttribute{
				Optional: true,
			},
			"customerWebhookUrl": schema.StringAttribute{
				Optional: true,
			},
			"reminderWebhookUrl": schema.StringAttribute{
				Optional: true,
			},
			"resourceWebhookUrl": schema.StringAttribute{
				Optional: true,
			},
			"webhookSignatureHash": schema.StringAttribute{
				Optional: true,
			},
			"disableEmailAndSmsNotifications": schema.BoolAttribute{
				Optional: true,
				Default:  booldefault.StaticBool(false),
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *companyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resp.Diagnostics.AddError(
		"Creating a company is not supported",
		"A company is expected to be created by an OnSched representative, use this resource to update the company",
	)
}

// Read refreshes the Terraform state with the latest data.
func (r *companyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state companyResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	c, err := r.client.GetCompany()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading OnSched company",
			err.Error(),
		)
		return
	}

	state.Object = types.StringValue(c.Object)
	state.ID = types.StringValue(c.ID)
	state.Name = types.StringValue(c.Name)
	state.RegistrationDate = types.StringValue(c.RegistrationDate)
	state.RegistrationEmail = types.StringValue(c.RegistrationEmail)
	state.DeletedStatus = types.BoolValue(c.DeletedStatus)
	state.DeletedTime = types.StringValue(c.DeletedTime)
	state.AddressLine1 = types.StringValue(c.AddressLine1)
	state.AddressLine2 = types.StringValue(c.AddressLine2)
	state.City = types.StringValue(c.City)
	state.State = types.StringValue(c.State)
	state.PostalCode = types.StringValue(c.PostalCode)
	state.Country = types.StringValue(c.Country)
	state.Phone = types.StringValue(c.Phone)
	state.Fax = types.StringValue(c.Fax)
	state.Email = types.StringValue(c.Email)
	state.Website = types.StringValue(c.Website)
	state.TimezoneID = types.StringValue(c.TimezoneID)
	state.TimezoneName = types.StringValue(c.TimezoneName)
	state.NotificationFromEmailAddress = types.StringValue(c.NotificationFromEmailAddress)
	state.NotificationFromName = types.StringValue(c.NotificationFromName)
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
func (r *companyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan companyResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	company := onsched.Company{
		Object:                          plan.Object.ValueString(),
		ID:                              plan.ID.ValueString(),
		Name:                            plan.Name.ValueString(),
		RegistrationDate:                plan.RegistrationDate.ValueString(),
		RegistrationEmail:               plan.RegistrationEmail.ValueString(),
		DeletedStatus:                   plan.DeletedStatus.ValueBool(),
		DeletedTime:                     plan.DeletedTime.ValueString(),
		AddressLine1:                    plan.AddressLine1.ValueString(),
		AddressLine2:                    plan.AddressLine2.ValueString(),
		City:                            plan.City.ValueString(),
		State:                           plan.State.ValueString(),
		PostalCode:                      plan.PostalCode.ValueString(),
		Country:                         plan.Country.ValueString(),
		Phone:                           plan.Phone.ValueString(),
		Fax:                             plan.Fax.ValueString(),
		Email:                           plan.Email.ValueString(),
		Website:                         plan.Website.ValueString(),
		TimezoneID:                      plan.TimezoneID.ValueString(),
		TimezoneName:                    plan.TimezoneName.ValueString(),
		NotificationFromEmailAddress:    plan.NotificationFromEmailAddress.ValueString(),
		NotificationFromName:            plan.NotificationFromName.ValueString(),
		BookingWebhookURL:               plan.BookingWebhookURL.ValueString(),
		CustomerWebhookURL:              plan.CustomerWebhookURL.ValueString(),
		ReminderWebhookURL:              plan.ReminderWebhookURL.ValueString(),
		ResourceWebhookURL:              plan.ResourceWebhookURL.ValueString(),
		WebhookSignatureHash:            plan.WebhookSignatureHash.ValueString(),
		DisableEmailAndSmsNotifications: plan.DisableEmailAndSmsNotifications.ValueBool(),
	}

	_, err := r.client.UpdateCompany(company)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating OnSched company",
			err.Error(),
		)
		return
	}

	company, err = r.client.GetCompany()

	plan.Object = types.StringValue(company.Object)
	plan.ID = types.StringValue(company.ID)
	plan.Name = types.StringValue(company.Name)
	plan.RegistrationDate = types.StringValue(company.RegistrationDate)
	plan.RegistrationEmail = types.StringValue(company.RegistrationEmail)
	plan.DeletedStatus = types.BoolValue(company.DeletedStatus)
	plan.DeletedTime = types.StringValue(company.DeletedTime)
	plan.AddressLine1 = types.StringValue(company.AddressLine1)
	plan.AddressLine2 = types.StringValue(company.AddressLine2)
	plan.City = types.StringValue(company.City)
	plan.State = types.StringValue(company.State)
	plan.PostalCode = types.StringValue(company.PostalCode)
	plan.Country = types.StringValue(company.Country)
	plan.Phone = types.StringValue(company.Phone)
	plan.Fax = types.StringValue(company.Fax)
	plan.Email = types.StringValue(company.Email)
	plan.Website = types.StringValue(company.Website)
	plan.TimezoneID = types.StringValue(company.TimezoneID)
	plan.TimezoneName = types.StringValue(company.TimezoneName)
	plan.NotificationFromEmailAddress = types.StringValue(company.NotificationFromEmailAddress)
	plan.NotificationFromName = types.StringValue(company.NotificationFromName)
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
func (r *companyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

type companyResourceModel struct {
	Object                          types.String `tfsdk:"object"`
	ID                              types.String `tfsdk:"id"`
	Name                            types.String `tfsdk:"name"`
	RegistrationDate                types.String `tfsdk:"registrationDate"`
	RegistrationEmail               types.String `tfsdk:"registrationEmail"`
	DeletedStatus                   types.Bool   `tfsdk:"deletedStatus"`
	DeletedTime                     types.String `tfsdk:"deletedTime"`
	AddressLine1                    types.String `tfsdk:"addressLine1"`
	AddressLine2                    types.String `tfsdk:"addressLine2"`
	City                            types.String `tfsdk:"city"`
	State                           types.String `tfsdk:"state"`
	PostalCode                      types.String `tfsdk:"postalCode"`
	Country                         types.String `tfsdk:"country"`
	Phone                           types.String `tfsdk:"phone"`
	Fax                             types.String `tfsdk:"fax"`
	Email                           types.String `tfsdk:"email"`
	Website                         types.String `tfsdk:"website"`
	TimezoneID                      types.String `tfsdk:"timezoneId"`
	TimezoneName                    types.String `tfsdk:"timezoneName"`
	NotificationFromEmailAddress    types.String `tfsdk:"notificationFromEmailAddress"`
	NotificationFromName            types.String `tfsdk:"notificationFromName"`
	BookingWebhookURL               types.String `tfsdk:"bookingWebhookUrl"`
	CustomerWebhookURL              types.String `tfsdk:"customerWebhookUrl"`
	ReminderWebhookURL              types.String `tfsdk:"reminderWebhookUrl"`
	ResourceWebhookURL              types.String `tfsdk:"resourceWebhookUrl"`
	WebhookSignatureHash            types.String `tfsdk:"webhookSignatureHash"`
	DisableEmailAndSmsNotifications types.Bool   `tfsdk:"disableEmailAndSmsNotifications"`
	LastUpdated                     types.String `tfsdk:"last_updated"`
}
