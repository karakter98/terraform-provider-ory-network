// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	ory "github.com/ory/client-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &ProjectProps{}
	_ datasource.DataSourceWithConfigure = &ProjectProps{}
)

func ProjectDataSource() datasource.DataSource {
	return &ProjectProps{}
}

// ProjectProps defines the data source implementation.
type ProjectProps struct {
	client *ory.APIClient
}

type ProjectDataSourceModelCorsType struct {
	Enabled types.Bool     `tfsdk:"enabled"`
	Origins []types.String `tfsdk:"origins"`
}

type ProjectDataSourceModelJsonConfigType struct {
	Config jsontypes.Normalized `tfsdk:"config"`
}

type ProjectDataSourceModelServicesType struct {
	Identity   *ProjectDataSourceModelJsonConfigType `tfsdk:"identity"`
	Oauth2     *ProjectDataSourceModelJsonConfigType `tfsdk:"oauth2"`
	Permission *ProjectDataSourceModelJsonConfigType `tfsdk:"permission"`
}

// ProjectDataSourceModel describes the data source data model.
type ProjectDataSourceModel struct {
	Id          types.String                        `tfsdk:"id"`
	Name        types.String                        `tfsdk:"name"`
	Slug        types.String                        `tfsdk:"slug"`
	CorsAdmin   *ProjectDataSourceModelCorsType     `tfsdk:"cors_admin"`
	CorsPublic  *ProjectDataSourceModelCorsType     `tfsdk:"cors_public"`
	RevisionId  types.String                        `tfsdk:"revision_id"`
	State       types.String                        `tfsdk:"state"`
	WorkspaceId types.String                        `tfsdk:"workspace_id"`
	Services    *ProjectDataSourceModelServicesType `tfsdk:"services"`
}

func (d *ProjectProps) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (d *ProjectProps) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	corsAttributeSchema := schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"origins": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
		},
		Computed: true,
	}
	jsonConfigSchema := schema.ObjectAttribute{
		AttributeTypes: map[string]attr.Type{
			"config": jsontypes.NormalizedType{},
		},
		Computed: true,
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Ory Network Project",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Project identifier",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Project name",
				Computed:            true,
			},
			"slug": schema.StringAttribute{
				MarkdownDescription: "Project slug",
				Computed:            true,
			},
			"cors_admin":  corsAttributeSchema,
			"cors_public": corsAttributeSchema,
			"revision_id": schema.StringAttribute{
				Computed: true,
			},
			"state": schema.StringAttribute{
				Computed: true,
			},
			"workspace_id": schema.StringAttribute{
				Computed: true,
			},
			"services": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"identity":   jsonConfigSchema,
					"oauth2":     jsonConfigSchema,
					"permission": jsonConfigSchema,
				},
				Computed: true,
			},
		},
	}
}

func (d *ProjectProps) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ory.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *ory.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ProjectProps) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ProjectDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, _, err := d.client.ProjectAPI.GetProject(ctx, data.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read project, got error: %s", err))
		return
	}

	data.Name = types.StringValue(project.Name)
	data.Slug = types.StringValue(project.Slug)
	data.RevisionId = types.StringValue(project.RevisionId)
	data.State = types.StringValue(project.State)
	if project.WorkspaceId.Get() != nil {
		data.WorkspaceId = types.StringValue(*project.WorkspaceId.Get())
	}

	origins := make([]types.String, 0)
	for _, origin := range project.CorsAdmin.Origins {
		origins = append(origins, types.StringValue(origin))
	}
	data.CorsAdmin = &ProjectDataSourceModelCorsType{
		Enabled: types.BoolValue(*project.CorsAdmin.Enabled),
		Origins: origins,
	}

	origins = make([]types.String, 0)
	for _, origin := range project.CorsPublic.Origins {
		origins = append(origins, types.StringValue(origin))
	}
	data.CorsPublic = &ProjectDataSourceModelCorsType{
		Enabled: types.BoolValue(*project.CorsPublic.Enabled),
		Origins: origins,
	}

	identityConfig, err := json.Marshal(project.Services.Identity.Config)
	if err != nil {
		resp.Diagnostics.AddError("JSON Error", fmt.Sprintf("Unable to serialize config, got error: %s", err))
	}
	oauth2Config, err := json.Marshal(project.Services.Oauth2.Config)
	if err != nil {
		resp.Diagnostics.AddError("JSON Error", fmt.Sprintf("Unable to serialize config, got error: %s", err))
	}
	permissionConfig, err := json.Marshal(project.Services.Permission.Config)
	if err != nil {
		resp.Diagnostics.AddError("JSON Error", fmt.Sprintf("Unable to serialize config, got error: %s", err))
	}

	data.Services = &ProjectDataSourceModelServicesType{
		Identity: &ProjectDataSourceModelJsonConfigType{
			Config: jsontypes.NewNormalizedValue(string(identityConfig)),
		},
		Oauth2: &ProjectDataSourceModelJsonConfigType{
			Config: jsontypes.NewNormalizedValue(string(oauth2Config)),
		},
		Permission: &ProjectDataSourceModelJsonConfigType{
			Config: jsontypes.NewNormalizedValue(string(permissionConfig)),
		},
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read project")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
