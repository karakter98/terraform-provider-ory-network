package provider

import (
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ory "github.com/ory/client-go"
	"math"
)

type ProjectModelCorsType struct {
	Enabled types.Bool     `tfsdk:"enabled"`
	Origins []types.String `tfsdk:"origins"`
}

type ProjectModelPermissionNamespaceType struct {
	Id   types.Int64  `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}

type ProjectModelPermissionConfigType struct {
	Namespaces []ProjectModelPermissionNamespaceType `tfsdk:"namespaces" json:"namespaces"`
}

type ProjectModelPermissionType struct {
	Config ProjectModelPermissionConfigType `tfsdk:"config" json:"config"`
}

type ProjectModelServicesType struct {
	Permission ProjectModelPermissionType `tfsdk:"permission"`
}

// ProjectModel describes the resource data model.
type ProjectModel struct {
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

func (data *ProjectModel) Deserialize(ctx *context.Context, project *ory.Project) {
	data.DeserializeBaseAttributes(project)
	data.DeserializeCorsSettings(project)
	data.DeserializeServices(ctx, project)
}

func (data *ProjectModel) DeserializeBaseAttributes(project *ory.Project) {
	data.Id = types.StringValue(project.Id)
	data.Name = types.StringValue(project.Name)
	data.Slug = types.StringValue(project.Slug)
	data.RevisionId = types.StringValue(project.RevisionId)
	data.State = types.StringValue(project.State)
	if project.WorkspaceId.Get() != nil {
		data.WorkspaceId = types.StringValue(*project.WorkspaceId.Get())
	}
}

func (data *ProjectModel) DeserializeCorsSettings(project *ory.Project) {
	origins := make([]attr.Value, 0)
	for _, origin := range project.CorsAdmin.Origins {
		origins = append(origins, types.StringValue(origin))
	}
	data.CorsAdmin = types.ObjectValueMust(
		map[string]attr.Type{
			"enabled": types.BoolType,
			"origins": types.ListType{ElemType: types.StringType},
		},
		map[string]attr.Value{
			"enabled": types.BoolValue(*project.CorsAdmin.Enabled),
			"origins": types.ListValueMust(types.StringType, origins),
		},
	)

	origins = make([]attr.Value, 0)
	for _, origin := range project.CorsPublic.Origins {
		origins = append(origins, types.StringValue(origin))
	}
	data.CorsPublic = types.ObjectValueMust(
		map[string]attr.Type{
			"enabled": types.BoolType,
			"origins": types.ListType{ElemType: types.StringType},
		},
		map[string]attr.Value{
			"enabled": types.BoolValue(*project.CorsPublic.Enabled),
			"origins": types.ListValueMust(types.StringType, origins),
		},
	)
}

func (data *ProjectModel) UnmarshalPermission(project *ory.Project) ProjectModelPermissionType {
	permission := ProjectModelPermissionType{}
	permission.Config = ProjectModelPermissionConfigType{}

	rawPermissionConfig := project.Services.Permission.Config

	namespaces := make([]ProjectModelPermissionNamespaceType, 0)
	for _, rawNamespace := range rawPermissionConfig["namespaces"].([]interface{}) {
		namespace := rawNamespace.(map[string]interface{})
		namespaces = append(namespaces, ProjectModelPermissionNamespaceType{
			// For some reason, Go deserializes JSON integers into float64, we have to convert to int64
			// with Round() to avoid floating point errors
			Id:   types.Int64Value(int64(int(math.Round(namespace["id"].(float64))))),
			Name: types.StringValue(namespace["name"].(string)),
		})
	}

	permission.Config.Namespaces = namespaces

	return permission
}

func (data *ProjectModel) DeserializeServices(ctx *context.Context, project *ory.Project) diag.Diagnostics {
	services := ProjectModelServicesType{
		Permission: data.UnmarshalPermission(project),
	}

	servicesValue, diags := types.ObjectValueFrom(*ctx, map[string]attr.Type{
		"permission": types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"config": types.ObjectType{
					AttrTypes: map[string]attr.Type{
						"namespaces": types.ListType{
							ElemType: types.ObjectType{
								AttrTypes: map[string]attr.Type{
									"id":   types.Int64Type,
									"name": types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}, services)

	data.Services = servicesValue
	return diags
}

func (data *ProjectModel) MarshalPermission(ctx *context.Context) (map[string]interface{}, error) {
	permissionAttr := data.Services.Attributes()["permission"]
	if permissionAttr != nil {
		permission := ProjectModelPermissionType{}
		permissionAttr.(basetypes.ObjectValue).As(*ctx, &permission, basetypes.ObjectAsOptions{})

		jsonEncoding, err := json.Marshal(permission.Config)
		if err != nil {
			return nil, err
		}

		jsonMapDecoding := make(map[string]interface{})
		err = json.Unmarshal(jsonEncoding, &jsonMapDecoding)
		if err != nil {
			return nil, err
		}

		return jsonMapDecoding, nil
	}
	return nil, nil
}

func (namespace *ProjectModelPermissionNamespaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   namespace.Id.ValueInt64(),
		"name": namespace.Name.ValueString(),
	})
}
