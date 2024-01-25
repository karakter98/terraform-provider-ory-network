package project

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ory "github.com/ory/client-go"
)

type ServicesType struct {
	Permission *PermissionType `tfsdk:"permission"`
}

func NewProjectServicesFromApiRepresentation(apiServices *ory.ProjectServices) *ServicesType {
	permission := NewProjectPermissionFromApiRepresentation(apiServices.Permission)

	return &ServicesType{
		Permission: permission,
	}
}

func NewProjectServicesFromTerraformRepresentation(objectValue *basetypes.ObjectValue, ctx *context.Context) *ServicesType {
	if objectValue.IsNull() || objectValue.IsUnknown() {
		return nil
	}

	services := &ServicesType{}
	diags := objectValue.As(*ctx, services, basetypes.ObjectAsOptions{})
	diags.Errors()
	return services
}

func (services *ServicesType) ToTerraformRepresentation(ctx *context.Context) basetypes.ObjectValue {
	servicesRepresentation, _ := types.ObjectValueFrom(
		*ctx,
		map[string]attr.Type{
			"permission": services.Permission.TerraformType(),
		},
		services,
	)
	return servicesRepresentation
}

func (services *ServicesType) ToApiRepresentation() (*ory.ProjectServices, error) {
	oryServices := ory.NewProjectServices()

	oryPermission, err := services.Permission.ToApiRepresentation()
	if err != nil {
		return nil, err
	}

	oryServices.SetPermission(*oryPermission)
	return oryServices, nil
}
