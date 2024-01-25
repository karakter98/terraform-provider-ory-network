package models

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ory "github.com/ory/client-go"
	"math"
)

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

func (namespace *ProjectModelPermissionNamespaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":   namespace.Id.ValueInt64(),
		"name": namespace.Name.ValueString(),
	})
}

func (config *ProjectModelPermissionConfigType) MarshalMap() (map[string]interface{}, error) {
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

func (permission *ProjectModelPermissionType) Unmarshal(oryPermission *ory.ProjectServicePermission) {
	rawPermissionConfig := oryPermission.Config

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

	permission.Config = ProjectModelPermissionConfigType{
		Namespaces: namespaces,
	}
}

func (permission *ProjectModelPermissionType) Marshal() (*ory.ProjectServicePermission, error) {
	config, err := permission.Config.MarshalMap()
	if err != nil {
		return nil, err
	}
	return ory.NewProjectServicePermission(config), nil
}
