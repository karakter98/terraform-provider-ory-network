// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
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
	_ datasource.DataSource              = &ProjectDataSourceProps{}
	_ datasource.DataSourceWithConfigure = &ProjectDataSourceProps{}
)

func ProjectDataSource() datasource.DataSource {
	return &ProjectDataSourceProps{}
}

// ProjectDataSourceProps defines the data source implementation.
type ProjectDataSourceProps struct {
	client *ory.APIClient
}

func (d *ProjectDataSourceProps) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (d *ProjectDataSourceProps) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

func (d *ProjectDataSourceProps) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ProjectDataSourceProps) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ProjectModel

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

	err = data.Deserialize(project, true)
	if err != nil {
		resp.Diagnostics.AddError("Deserialization Error", fmt.Sprintf("Unable to deserialize project, got error: %s", err))
		return
	}

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read project")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
