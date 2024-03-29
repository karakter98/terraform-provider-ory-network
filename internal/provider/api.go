package provider

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	ory "github.com/ory/client-go"
)

func getSessionToken(c *ory.APIClient, email *string, password *string, ctx *context.Context) (*string, error) {
	req := c.FrontendAPI.CreateNativeLoginFlow(*ctx)

	flow, response, err := req.Execute()

	if err != nil {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(response.Body)
		respBody := buf.String()

		tflog.Error(*ctx, fmt.Sprintf("Could not create Ory Network flow: %s", respBody))
		return nil, err
	}

	body := ory.UpdateLoginFlowBody{
		UpdateLoginFlowWithPasswordMethod: &ory.UpdateLoginFlowWithPasswordMethod{
			Identifier: *email,
			Password:   *password,
			Method:     "password",
		},
	}

	login, response, err := c.FrontendAPI.UpdateLoginFlow(*ctx).Flow(flow.Id).UpdateLoginFlowBody(body).Execute()
	if err != nil {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(response.Body)
		respBody := buf.String()

		tflog.Error(*ctx, fmt.Sprintf("Could not complete Ory Network login: %s", respBody))
		return nil, err
	}

	tflog.Debug(*ctx, fmt.Sprintf("Received Ory Network Session Token %s", *login.SessionToken))

	sessionToken := *login.SessionToken
	return &sessionToken, nil
}

func createProject(c *ory.APIClient, data *ProjectModel, ctx *context.Context) (*ory.Project, error) {
	if data.Name.IsUnknown() || data.Name.IsNull() {
		return nil, errors.New("project name must be set and a known value")
	}
	if data.WorkspaceId.IsUnknown() {
		return nil, errors.New("workspace ID must be a known value")
	}

	createProjectBody := ory.NewCreateProjectBody(data.Name.ValueString())
	if data.WorkspaceId.IsNull() {
		createProjectBody.SetWorkspaceIdNil()
	} else {
		createProjectBody.SetWorkspaceId(data.WorkspaceId.ValueString())
	}
	project, _, err := c.ProjectAPI.CreateProject(*ctx).CreateProjectBody(*createProjectBody).Execute()

	if err != nil {
		return nil, err
	}

	return project, nil
}

func updateProject(c *ory.APIClient, newData *ProjectModel, oldData *ProjectModel, ctx *context.Context) (*ory.Project, error) {
	if newData.Name.IsUnknown() || newData.Name.IsNull() {
		return nil, errors.New("project name must be set and a known value")
	}
	if newData.Id.IsUnknown() || newData.Id.IsNull() {
		return nil, errors.New("project ID must be set and a known value")
	}

	adminCorsModel := ProjectModelCorsType{}
	if !newData.CorsAdmin.IsNull() && !newData.CorsAdmin.IsUnknown() {
		newData.CorsAdmin.As(*ctx, &adminCorsModel, basetypes.ObjectAsOptions{})
	} else if oldData != nil && !oldData.CorsAdmin.IsNull() && !oldData.CorsAdmin.IsUnknown() {
		oldData.CorsAdmin.As(*ctx, &adminCorsModel, basetypes.ObjectAsOptions{})
	}
	var corsAdminOrigins []string
	for _, origin := range adminCorsModel.Origins {
		corsAdminOrigins = append(corsAdminOrigins, origin.ValueString())
	}
	adminCors := ory.ProjectCors{
		Enabled: adminCorsModel.Enabled.ValueBoolPointer(),
		Origins: corsAdminOrigins,
	}

	publicCorsModel := ProjectModelCorsType{}
	if !newData.CorsPublic.IsNull() && !newData.CorsPublic.IsUnknown() {
		newData.CorsPublic.As(*ctx, &publicCorsModel, basetypes.ObjectAsOptions{})
	} else if oldData != nil && !oldData.CorsPublic.IsNull() && !oldData.CorsPublic.IsUnknown() {
		oldData.CorsPublic.As(*ctx, &publicCorsModel, basetypes.ObjectAsOptions{})
	}
	var corsPublicOrigins []string
	for _, origin := range publicCorsModel.Origins {
		corsPublicOrigins = append(corsPublicOrigins, origin.ValueString())
	}
	publicCors := ory.ProjectCors{
		Enabled: publicCorsModel.Enabled.ValueBoolPointer(),
		Origins: corsPublicOrigins,
	}

	projectServices := ory.NewProjectServices()

	identityConfigMap, err := newData.GetServicesFieldConfig("identity")
	if err != nil {
		return nil, err
	}
	if identityConfigMap != nil {
		projectServices.SetIdentity(ory.ProjectServiceIdentity{
			Config: identityConfigMap,
		})
	}

	oauth2ConfigMap, err := newData.GetServicesFieldConfig("oauth2")
	if err != nil {
		return nil, err
	}
	if oauth2ConfigMap != nil {
		projectServices.SetOauth2(ory.ProjectServiceOAuth2{
			Config: oauth2ConfigMap,
		})
	}

	permissionConfigMap, err := newData.GetServicesFieldConfig("permission")
	if err != nil {
		return nil, err
	}
	if permissionConfigMap != nil {
		projectServices.SetPermission(ory.ProjectServicePermission{
			Config: permissionConfigMap,
		})
	}

	setProjectBody := ory.NewSetProject(
		adminCors,
		publicCors,
		newData.Name.ValueString(),
		*projectServices,
	)
	setProjectResponse, _, err := c.ProjectAPI.SetProject(*ctx, newData.Id.ValueString()).SetProject(*setProjectBody).Execute()

	if err != nil {
		return nil, err
	}

	project := setProjectResponse.Project

	return &project, nil
}

func readProject(c *ory.APIClient, data *ProjectModel, ctx *context.Context) (*ory.Project, error) {
	if data.Id.IsUnknown() || data.Id.IsNull() {
		return nil, errors.New("project ID must be set and a known value")
	}
	project, _, err := c.ProjectAPI.GetProject(*ctx, data.Id.ValueString()).Execute()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func deleteProject(c *ory.APIClient, data *ProjectModel, ctx *context.Context) error {
	if data.Id.IsUnknown() || data.Id.IsNull() {
		return errors.New("project ID must be set and a known value")
	}
	_, err := c.ProjectAPI.PurgeProject(*ctx, data.Id.ValueString()).Execute()
	return err
}
