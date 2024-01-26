package api

import (
	"bytes"
	"context"
	"errors"
	ory "github.com/ory/client-go"
	projectmodel "terraform-provider-ory-network/internal/models/project"
)

func CreateProject(c *ory.APIClient, data *projectmodel.ProjectType, ctx *context.Context) (*ory.Project, error) {
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
	project, response, err := c.ProjectAPI.CreateProject(*ctx).CreateProjectBody(*createProjectBody).Execute()

	if err != nil {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(response.Body)
		respBody := buf.String()

		return nil, errors.Join(err, errors.New(respBody))
	}

	return project, nil
}

func UpdateProject(c *ory.APIClient, data *projectmodel.ProjectType, ctx *context.Context) (*ory.Project, error) {
	if data.Name.IsUnknown() || data.Name.IsNull() {
		return nil, errors.New("project name must be set and a known value")
	}
	if data.Id.IsUnknown() || data.Id.IsNull() {
		return nil, errors.New("project ID must be set and a known value")
	}

	adminCors := ory.ProjectCors{}
	newAdminCorsModel := projectmodel.NewProjectCorsFromTerraformRepresentation(&data.CorsAdmin, ctx)

	if newAdminCorsModel != nil {
		adminCors = *newAdminCorsModel.ToApiRepresentation()
	}

	publicCors := ory.ProjectCors{}
	newPublicCorsModel := projectmodel.NewProjectCorsFromTerraformRepresentation(&data.CorsPublic, ctx)

	if newPublicCorsModel != nil {
		publicCors = *newPublicCorsModel.ToApiRepresentation()
	}

	services := ory.ProjectServices{
		Permission: ory.NewProjectServicePermission(make(map[string]interface{})),
	}
	newServices := projectmodel.NewProjectServicesFromTerraformRepresentation(&data.Services, ctx)

	if newServices != nil && newServices.Permission != nil {
		permission, err := newServices.Permission.ToApiRepresentation()
		if err != nil {
			return nil, err
		}
		services.SetPermission(*permission)
	}

	setProjectBody := ory.NewSetProject(
		adminCors,
		publicCors,
		data.Name.ValueString(),
		services,
	)
	setProjectResponse, response, err := c.ProjectAPI.SetProject(*ctx, data.Id.ValueString()).SetProject(*setProjectBody).Execute()

	if err != nil {
		buf := new(bytes.Buffer)
		_, _ = buf.ReadFrom(response.Body)
		respBody := buf.String()

		return nil, errors.Join(err, errors.New(respBody))
	}

	project := setProjectResponse.Project

	return &project, nil
}

func ReadProject(c *ory.APIClient, data *projectmodel.ProjectType, ctx *context.Context) (*ory.Project, error) {
	if data.Id.IsUnknown() || data.Id.IsNull() {
		return nil, errors.New("project ID must be set and a known value")
	}
	project, _, err := c.ProjectAPI.GetProject(*ctx, data.Id.ValueString()).Execute()
	if err != nil {
		return nil, err
	}
	return project, nil
}

func DeleteProject(c *ory.APIClient, data *projectmodel.ProjectType, ctx *context.Context) error {
	if data.Id.IsUnknown() || data.Id.IsNull() {
		return errors.New("project ID must be set and a known value")
	}
	_, err := c.ProjectAPI.PurgeProject(*ctx, data.Id.ValueString()).Execute()
	return err
}
