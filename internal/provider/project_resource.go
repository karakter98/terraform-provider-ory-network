// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ory "github.com/ory/client-go"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProjectResourceProps{}
var _ resource.ResourceWithConfigure = &ProjectResourceProps{}
var _ resource.ResourceWithImportState = &ProjectResourceProps{}

func ProjectResource() resource.Resource {
	return &ProjectResourceProps{}
}

// ProjectResourceProps defines the resource implementation.
type ProjectResourceProps struct {
	client *ory.APIClient
}

func (r *ProjectResourceProps) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *ProjectResourceProps) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	corsAttributeSchema := schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"origins": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
		},
		Optional: true,
		Computed: true,
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Ory Network Project",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Project identifier (UUID)",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				Optional: true,
			},
			"services": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"permission": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"namespaces": schema.ListNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.Int64Attribute{
													Optional: true,
													Computed: true,
												},
												"name": schema.StringAttribute{
													Optional: true,
													Computed: true,
												},
											},
										},
										Optional: true,
										Computed: true,
									},
								},
								Optional: true,
								Computed: true,
							},
						},
						Optional: true,
						Computed: true,
					},
				},
				Optional: true,
				Computed: true,
			},
		},
	}
}

func (r *ProjectResourceProps) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ory.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ory.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ProjectResourceProps) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var createData ProjectModel
	var updateData ProjectModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &createData)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &updateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, err := createProject(r.client, &createData, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create project, got error: %s", err))
		return
	}

	createData.Deserialize(&ctx, project)
	// Save intermediate data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &createData)...)

	updateData.Id = createData.Id
	project, err = updateProject(r.client, &updateData, &createData, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", fmt.Sprintf("Unable to update project settings, got error: %s", err))
		return
	}
	updateData.Deserialize(&ctx, project)

	// Save final data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updateData)...)
}

func (r *ProjectResourceProps) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ProjectModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, err := readProject(r.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read project, got error: %s", err))
		return
	}

	data.Deserialize(&ctx, project)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResourceProps) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var planData ProjectModel
	var stateData ProjectModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, err := updateProject(r.client, &planData, &stateData, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", fmt.Sprintf("Unable to update project settings, got error: %s", err))
		return
	}

	planData.Deserialize(&ctx, project)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &planData)...)
}

func (r *ProjectResourceProps) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ProjectModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	err := deleteProject(r.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete project, got error: %s", err))
		return
	}
}

func (r *ProjectResourceProps) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
