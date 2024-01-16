// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
	jsonConfigSchema := schema.ObjectAttribute{
		AttributeTypes: map[string]attr.Type{
			"config": jsontypes.NormalizedType{},
		},
		Optional: true,
		Computed: true,
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Ory Network Project",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "Project identifier",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Project name",
				Required:            true,
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
				Optional: true,
			},
			"services": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"identity":   jsonConfigSchema,
					"oauth2":     jsonConfigSchema,
					"permission": jsonConfigSchema,
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
	var data ProjectModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	createProjectBody := ory.NewCreateProjectBody(data.Name.ValueString())
	if data.WorkspaceId.IsNull() {
		createProjectBody.SetWorkspaceIdNil()
	} else {
		createProjectBody.SetWorkspaceId(data.WorkspaceId.ValueString())
	}
	project, _, err := r.client.ProjectAPI.CreateProject(ctx).CreateProjectBody(*createProjectBody).Execute()

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create project, got error: %s", err))
		return
	}

	err = data.Deserialize(project, false)
	if err != nil {
		resp.Diagnostics.AddError("Deserialization Error", fmt.Sprintf("Unable to deserialize project, got error: %s", err))
		return
	}
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

	err = r.UpdateProjectSettings(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", fmt.Sprintf("Unable to update project settings, got error: %s", err))
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
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
	project, _, err := r.client.ProjectAPI.GetProject(ctx, data.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read project, got error: %s", err))
		return
	}

	err = data.Deserialize(project, true)
	if err != nil {
		resp.Diagnostics.AddError("Deserialization Error", fmt.Sprintf("Unable to deserialize project, got error: %s", err))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResourceProps) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ProjectModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	err := r.UpdateProjectSettings(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", fmt.Sprintf("Unable to update project settings, got error: %s", err))
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
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
	_, err := r.client.ProjectAPI.PurgeProject(ctx, data.Id.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete project, got error: %s", err))
		return
	}
}

func (r *ProjectResourceProps) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ProjectResourceProps) UpdateProjectSettings(ctx context.Context, data *ProjectModel) error {
	services, err := data.SerializeServices()
	if err != nil {
		return err
	}

	adminCors := ProjectModelCorsType{}
	data.CorsAdmin.As(ctx, &adminCors, basetypes.ObjectAsOptions{})
	publicCors := ProjectModelCorsType{}
	data.CorsPublic.As(ctx, &publicCors, basetypes.ObjectAsOptions{})

	setProjectBody := ory.NewSetProject(
		*data.SerializeCorsSettings(&adminCors),
		*data.SerializeCorsSettings(&publicCors),
		data.Name.ValueString(),
		*services,
	)
	setProjectResponse, _, err := r.client.ProjectAPI.SetProject(ctx, data.Id.ValueString()).SetProject(*setProjectBody).Execute()

	if err != nil {
		return err
	}

	project := setProjectResponse.Project

	err = data.Deserialize(&project, true)
	if err != nil {
		return err
	}
	return nil
}
