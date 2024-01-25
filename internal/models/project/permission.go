package project

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ory "github.com/ory/client-go"
	"math"
)

type PermissionNamespaceType struct {
	Id   types.Int64  `tfsdk:"id" json:"id"`
	Name types.String `tfsdk:"name" json:"name"`
}

type PermissionConfigType struct {
	Namespaces []PermissionNamespaceType `tfsdk:"namespaces" json:"namespaces"`
}

type PermissionType struct {
	Config PermissionConfigType `tfsdk:"config" json:"config"`
}

// MarshalJSON For json.Marshal compatibility
func (namespace *PermissionNamespaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   namespace.Id.ValueInt64(),
		"name": namespace.Name.ValueString(),
	})
}

func (config *PermissionConfigType) ToApiRepresentation() (map[string]interface{}, error) {
	jsonEncoding, err := json.Marshal(config)
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

func NewProjectPermissionFromApiRepresentation(apiPermission *ory.ProjectServicePermission) *PermissionType {
	rawPermissionConfig := apiPermission.Config

	namespaces := make([]PermissionNamespaceType, 0)
	for _, rawNamespace := range rawPermissionConfig["namespaces"].([]interface{}) {
		namespace := rawNamespace.(map[string]interface{})
		namespaces = append(namespaces, PermissionNamespaceType{
			// For some reason, Go deserializes JSON integers into float64, we have to convert to int64
			// with Round() to avoid floating point errors
			Id:   types.Int64Value(int64(int(math.Round(namespace["id"].(float64))))),
			Name: types.StringValue(namespace["name"].(string)),
		})
	}

	return &PermissionType{
		Config: PermissionConfigType{
			Namespaces: namespaces,
		},
	}
}

func (permission *PermissionType) ToApiRepresentation() (*ory.ProjectServicePermission, error) {
	config, err := permission.Config.ToApiRepresentation()
	if err != nil {
		return nil, err
	}
	return ory.NewProjectServicePermission(config), nil
}

func (namespace *PermissionNamespaceType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":   types.Int64Type,
			"name": types.StringType,
		},
	}
}

func (config *PermissionConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"namespaces": types.ListType{
				ElemType: (&PermissionNamespaceType{}).TerraformType(),
			},
		},
	}
}

func (permission *PermissionType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": (&PermissionConfigType{}).TerraformType(),
		},
	}
}
