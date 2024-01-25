package project

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ory "github.com/ory/client-go"
)

// ProjectType describes the resource data model.
//
//goland:noinspection GoNameStartsWithPackageName
type ProjectType struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Slug        types.String `tfsdk:"slug"`
	CorsAdmin   types.Object `tfsdk:"cors_admin"`
	CorsPublic  types.Object `tfsdk:"cors_public"`
	RevisionId  types.String `tfsdk:"revision_id"`
	State       types.String `tfsdk:"state"`
	WorkspaceId types.String `tfsdk:"workspace_id"`
	Services    types.Object `tfsdk:"services"`
}

func NewProjectFromApiRepresentation(apiProject *ory.Project, ctx *context.Context) *ProjectType {
	corsAdmin := NewProjectCorsFromApiRepresentation(apiProject.CorsAdmin)
	corsPublic := NewProjectCorsFromApiRepresentation(apiProject.CorsPublic)
	services := NewProjectServicesFromApiRepresentation(&apiProject.Services)

	return &ProjectType{
		Id:          types.StringValue(apiProject.Id),
		Name:        types.StringValue(apiProject.Name),
		Slug:        types.StringValue(apiProject.Slug),
		CorsAdmin:   corsAdmin.ToTerraformRepresentation(),
		CorsPublic:  corsPublic.ToTerraformRepresentation(),
		RevisionId:  types.String{},
		State:       types.String{},
		WorkspaceId: types.String{},
		Services:    services.ToTerraformRepresentation(ctx),
	}
}
