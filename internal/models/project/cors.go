package project

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ory "github.com/ory/client-go"
)

type CorsType struct {
	Enabled types.Bool     `tfsdk:"enabled"`
	Origins []types.String `tfsdk:"origins"`
}

func NewProjectCorsFromApiRepresentation(apiCors *ory.ProjectCors) *CorsType {
	origins := make([]types.String, 0)
	for _, origin := range apiCors.Origins {
		origins = append(origins, types.StringValue(origin))
	}
	return &CorsType{
		Enabled: types.BoolValue(*apiCors.Enabled),
		Origins: origins,
	}
}

func NewProjectCorsFromTerraformRepresentation(objectValue *basetypes.ObjectValue, ctx *context.Context) *CorsType {
	if objectValue.IsNull() || objectValue.IsUnknown() {
		return nil
	}

	cors := &CorsType{}
	objectValue.As(*ctx, cors, basetypes.ObjectAsOptions{})
	return cors
}

func (cors *CorsType) ToApiRepresentation() *ory.ProjectCors {
	origins := make([]string, 0)
	for _, origin := range cors.Origins {
		origins = append(origins, origin.ValueString())
	}
	return &ory.ProjectCors{
		Enabled: cors.Enabled.ValueBoolPointer(),
		Origins: origins,
	}
}

func (cors *CorsType) ToTerraformRepresentation() basetypes.ObjectValue {
	origins := make([]attr.Value, 0)
	for _, origin := range cors.Origins {
		origins = append(origins, types.StringValue(origin.ValueString()))
	}

	return types.ObjectValueMust(
		map[string]attr.Type{
			"enabled": types.BoolType,
			"origins": types.ListType{ElemType: types.StringType},
		},
		map[string]attr.Value{
			"enabled": types.BoolValue(cors.Enabled.ValueBool()),
			"origins": types.ListValueMust(types.StringType, origins),
		},
	)
}
