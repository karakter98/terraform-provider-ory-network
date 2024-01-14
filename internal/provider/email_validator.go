package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"net/mail"
)

var _ validator.String = emailValidator{}

type emailValidator struct{}

func (v emailValidator) Description(_ context.Context) string {
	return fmt.Sprintf("value must be a valid email address")
}

func (v emailValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v emailValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() {
		return
	}
	_, err := mail.ParseAddress(request.ConfigValue.ValueString())
	if err != nil {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			v.Description(ctx),
			fmt.Sprintf("%s", request.ConfigValue.ValueString()),
		))
	}
}

func EmailValidator() validator.String {
	return emailValidator{}
}
