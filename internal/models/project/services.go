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
	Identity   *IdentityType   `tfsdk:"identity"`
}

func NewProjectServicesFromApiRepresentation(apiServices *ory.ProjectServices) *ServicesType {
	permission := NewProjectPermissionFromApiRepresentation(apiServices.Permission)
	identity := NewProjectIdentityFromApiRepresentation(apiServices.Identity)
	return &ServicesType{
		Permission: permission,
		Identity:   identity,
	}
}

func NewProjectServicesFromTerraformRepresentation(objectValue *basetypes.ObjectValue, ctx *context.Context) *ServicesType {
	if objectValue.IsNull() || objectValue.IsUnknown() {
		return nil
	}

	services := &ServicesType{}
	objectValue.As(*ctx, services, basetypes.ObjectAsOptions{UnhandledUnknownAsEmpty: true})
	return services
}

func (services *ServicesType) ToTerraformRepresentation(ctx *context.Context) basetypes.ObjectValue {
	servicesRepresentation, _ := types.ObjectValueFrom(
		*ctx,
		map[string]attr.Type{
			"permission": services.Permission.TerraformType(),
			"identity":   services.Identity.TerraformType(),
		},
		services,
	)
	return servicesRepresentation
}

func (services *ServicesType) MergeWith(other *ServicesType) {
	if services.Permission == nil {
		services.Permission = other.Permission
	}
	if services.Identity == nil {
		services.Identity = other.Identity
	}
}
