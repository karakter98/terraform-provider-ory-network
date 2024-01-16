package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	ory "github.com/ory/client-go"
)

type ProjectModelCorsType struct {
	Enabled types.Bool     `tfsdk:"enabled"`
	Origins []types.String `tfsdk:"origins"`
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

func (data *ProjectModel) Deserialize(project *ory.Project, overwrite bool) error {
	data.DeserializeComputedAttributes(project)

	data.DeserializeCorsSettings(project, overwrite)

	return data.DeserializeServicesConfig(project, overwrite)
}

func (data *ProjectModel) DeserializeComputedAttributes(project *ory.Project) {
	data.Id = types.StringValue(project.Id)
	data.Name = types.StringValue(project.Name)
	data.Slug = types.StringValue(project.Slug)
	data.RevisionId = types.StringValue(project.RevisionId)
	data.State = types.StringValue(project.State)
	if project.WorkspaceId.Get() != nil {
		data.WorkspaceId = types.StringValue(*project.WorkspaceId.Get())
	}
}

func (data *ProjectModel) DeserializeCorsSettings(project *ory.Project, overwrite bool) {
	if data.CorsAdmin.IsNull() || overwrite {
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
	}

	if data.CorsPublic.IsNull() || overwrite {
		origins := make([]attr.Value, 0)
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
}

func (data *ProjectModel) DeserializeServicesConfig(project *ory.Project, overwrite bool) error {
	identityConfig, err := json.Marshal(project.Services.Identity.Config)
	if err != nil {
		return err
	}
	oauth2Config, err := json.Marshal(project.Services.Oauth2.Config)
	if err != nil {
		return err
	}
	permissionConfig, err := json.Marshal(project.Services.Permission.Config)
	if err != nil {
		return err
	}

	identityAttr := data.Services.Attributes()["identity"]
	if identityAttr != nil && !overwrite {
		identityConfigAttr := identityAttr.(basetypes.ObjectValue).Attributes()["config"]
		if identityConfigAttr != nil {
			identityConfig = []byte(identityConfigAttr.(jsontypes.Normalized).ValueString())
		}
	}
	oauth2Attr := data.Services.Attributes()["oauth2"]
	if oauth2Attr != nil && !overwrite {
		oauth2ConfigAttr := oauth2Attr.(basetypes.ObjectValue).Attributes()["config"]
		if oauth2ConfigAttr != nil {
			oauth2Config = []byte(oauth2ConfigAttr.(jsontypes.Normalized).ValueString())
		}
	}
	permissionAttr := data.Services.Attributes()["permission"]
	if permissionAttr != nil && !overwrite {
		permissionConfigAttr := permissionAttr.(basetypes.ObjectValue).Attributes()["config"]
		if permissionConfigAttr != nil {
			permissionConfig = []byte(permissionConfigAttr.(jsontypes.Normalized).ValueString())
		}
	}

	data.Services = types.ObjectValueMust(
		map[string]attr.Type{
			"identity": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"config": jsontypes.NormalizedType{},
				},
			},
			"oauth2": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"config": jsontypes.NormalizedType{},
				},
			},
			"permission": types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"config": jsontypes.NormalizedType{},
				},
			},
		},
		map[string]attr.Value{
			"identity": types.ObjectValueMust(
				map[string]attr.Type{
					"config": jsontypes.NormalizedType{},
				},
				map[string]attr.Value{
					"config": jsontypes.NewNormalizedValue(string(identityConfig)),
				},
			),
			"oauth2": types.ObjectValueMust(
				map[string]attr.Type{
					"config": jsontypes.NormalizedType{},
				},
				map[string]attr.Value{
					"config": jsontypes.NewNormalizedValue(string(oauth2Config)),
				},
			),
			"permission": types.ObjectValueMust(
				map[string]attr.Type{
					"config": jsontypes.NormalizedType{},
				},
				map[string]attr.Value{
					"config": jsontypes.NewNormalizedValue(string(permissionConfig)),
				},
			),
		},
	)
	return nil
}

func (data *ProjectModel) GetServicesFieldConfig(fieldName string) (map[string]interface{}, error) {
	serviceAttr := data.Services.Attributes()[fieldName]
	if serviceAttr != nil {
		configAttr := serviceAttr.(basetypes.ObjectValue).Attributes()["config"]
		if configAttr != nil {
			config := make(map[string]interface{})
			err := json.Unmarshal([]byte(configAttr.(jsontypes.Normalized).ValueString()), &config)
			if err != nil {
				return nil, err
			}
			return config, nil
		}
	}
	return nil, nil
}
