// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ory "github.com/ory/client-go"
	"os"
)

// Ensure OryNetworkProvider satisfies various provider interfaces.
var _ provider.Provider = &OryNetworkProvider{}

// OryNetworkProvider defines the provider implementation.
type OryNetworkProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// OryNetworkProviderModel describes the provider data model.
type OryNetworkProviderModel struct {
	Email    types.String `tfsdk:"email"`
	Password types.String `tfsdk:"password"`
}

func (p *OryNetworkProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "orynetwork"
	resp.Version = p.version
}

func (p *OryNetworkProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				MarkdownDescription: "Email address used to login to Ory Network",
				Description:         "Email address used to login to Ory Network",
				Required:            true,
				//Optional: true,
				Validators: []validator.String{
					EmailValidator(),
				},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password used to login to Ory Network",
				Description:         "Password used to login to Ory Network",
				Required:            true,
				//Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func (p *OryNetworkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config OryNetworkProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Email.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("email"),
			"Unknown Ory Network Email",
			"The provider cannot create the Ory Network account client as there is an unknown configuration value for the Ory Network account email. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ORY_NETWORK_EMAIL environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Ory Network Password",
			"The provider cannot create the Ory Network account client as there is an unknown configuration value for the Ory Network account password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the ORY_NETWORK_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	email := os.Getenv("ORY_NETWORK_EMAIL")
	password := os.Getenv("ORY_NETWORK_PASSWORD")
	authEndpoint := "https://project.console.ory.sh"
	apiEndpoint := "https://api.console.ory.sh"

	if !config.Email.IsNull() {
		email = config.Email.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if email == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("email"),
			"Missing Ory Network Email",
			"The provider cannot create the Ory Network account client as there is a missing or empty value for the Ory Network account email. "+
				"Set the email value in the configuration or use the ORY_NETWORK_EMAIL environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Ory Network Password",
			"The provider cannot create the Ory Network account client as there is a missing or empty value for the Ory Network account password. "+
				"Set the password value in the configuration or use the ORY_NETWORK_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Create a new Ory Network client using the configuration values
	configuration := ory.NewConfiguration()
	configuration.Servers = ory.ServerConfigurations{{URL: authEndpoint}}
	client := ory.NewAPIClient(configuration)

	sessionToken, err := signin(client, &email, &password, &ctx)

	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Ory Network API Client",
			"An unexpected error occurred when creating the Ory Network API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Ory Network Client Error: "+err.Error(),
		)
		return
	}

	client.GetConfig().AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", *sessionToken))
	client.GetConfig().Servers = ory.ServerConfigurations{{URL: apiEndpoint}}

	// Make the Ory Network client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *OryNetworkProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}

func (p *OryNetworkProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		ProjectDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &OryNetworkProvider{
			version: version,
		}
	}
}
