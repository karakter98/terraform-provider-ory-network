package provider

import (
	"bytes"
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	ory "github.com/ory/client-go"
)

func signin(c *ory.APIClient, email *string, password *string, ctx *context.Context) (*string, error) {
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
