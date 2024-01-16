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

func (data *ProjectModel) SerializeServices() (*ory.ProjectServices, error) {
	projectServices := ory.NewProjectServices()

	identityAttr := data.Services.Attributes()["identity"]
	if identityAttr != nil {
		identityConfigAttr := identityAttr.(basetypes.ObjectValue).Attributes()["config"]
		if identityConfigAttr != nil {
			identityConfigRaw := identityConfigAttr.(jsontypes.Normalized)
			identityConfig := make(map[string]interface{})
			err := json.Unmarshal([]byte(identityConfigRaw.StringValue.ValueString()), &identityConfig)
			if err != nil {
				return projectServices, err
			}
			projectServices.SetIdentity(ory.ProjectServiceIdentity{
				Config: identityConfig,
			})
		}
	}
	if !projectServices.HasIdentity() {
		projectServices.SetIdentity(ory.ProjectServiceIdentity{
			Config: make(map[string]interface{}),
		})
	}

	oauth2Attr := data.Services.Attributes()["oauth2"]
	if oauth2Attr != nil {
		oauth2ConfigAttr := oauth2Attr.(basetypes.ObjectValue).Attributes()["config"]
		if oauth2ConfigAttr != nil {
			oauth2ConfigRaw := oauth2ConfigAttr.(jsontypes.Normalized)
			oauth2Config := make(map[string]interface{})
			err := json.Unmarshal([]byte(oauth2ConfigRaw.StringValue.ValueString()), &oauth2Config)
			if err != nil {
				return projectServices, err
			}
			projectServices.SetOauth2(ory.ProjectServiceOAuth2{
				Config: oauth2Config,
			})
		}
	}
	if !projectServices.HasOauth2() {
		projectServices.SetOauth2(ory.ProjectServiceOAuth2{
			Config: make(map[string]interface{}),
		})
	}

	permissionAttr := data.Services.Attributes()["permission"]
	if permissionAttr != nil {
		permissionConfigAttr := permissionAttr.(basetypes.ObjectValue).Attributes()["config"]
		if permissionConfigAttr != nil {
			permissionConfigRaw := permissionConfigAttr.(jsontypes.Normalized)
			permissionConfig := make(map[string]interface{})
			err := json.Unmarshal([]byte(permissionConfigRaw.StringValue.ValueString()), &permissionConfig)
			if err != nil {
				return projectServices, err
			}
			projectServices.SetPermission(ory.ProjectServicePermission{
				Config: permissionConfig,
			})
		}
	}
	if !projectServices.HasPermission() {
		projectServices.SetPermission(ory.ProjectServicePermission{
			Config: make(map[string]interface{}),
		})
	}

	return projectServices, nil
}

func (data *ProjectModel) SerializeCorsSettings(cors *ProjectModelCorsType) *ory.ProjectCors {
	var corsOrigins []string
	for _, origin := range cors.Origins {
		corsOrigins = append(corsOrigins, origin.ValueString())
	}
	return &ory.ProjectCors{
		Enabled: cors.Enabled.ValueBoolPointer(),
		Origins: corsOrigins,
	}
}
