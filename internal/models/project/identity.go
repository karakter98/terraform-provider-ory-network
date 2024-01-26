package project

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ory "github.com/ory/client-go"
	"math"
	"reflect"
)

type IdentityIdentitySchemaType struct {
	Id  types.String `tfsdk:"id" json:"id"`
	Url types.String `tfsdk:"url" json:"url"`
}

type IdentityIdentityType struct {
	DefaultSchemaId types.String                 `tfsdk:"default_schema_id" json:"default_schema_id"`
	Schemas         []IdentityIdentitySchemaType `tfsdk:"schemas" json:"schemas"`
}

type IdentitySelfServiceMethodsLinkConfigType struct {
	Lifespan types.String `tfsdk:"lifespan" json:"lifespan"`
	BaseUrl  types.String `tfsdk:"base_url" json:"base_url"`
}

type IdentitySelfServiceMethodsLinkType struct {
	Config  *IdentitySelfServiceMethodsLinkConfigType `tfsdk:"config" json:"config,omitempty"`
	Enabled types.Bool                                `tfsdk:"enabled" json:"enabled"`
}

type IdentitySelfServiceMethodsCodeConfigType struct {
	Lifespan types.String `tfsdk:"lifespan" json:"lifespan"`
}

type IdentitySelfServiceMethodsCodeType struct {
	PasswordlessLoginFallbackEnabled types.Bool                                `tfsdk:"passwordless_login_fallback_enabled" json:"passwordless_login_fallback_enabled"`
	Enabled                          types.Bool                                `tfsdk:"enabled" json:"enabled"`
	PasswordlessEnabled              types.Bool                                `tfsdk:"passwordless_enabled" json:"passwordless_enabled"`
	Config                           *IdentitySelfServiceMethodsCodeConfigType `tfsdk:"config" json:"config,omitempty"`
}

type IdentitySelfServiceMethodsPasswordConfigType struct {
	HaveIBeenPwnedEnabled            types.Bool   `tfsdk:"haveibeenpwned_enabled" json:"haveibeenpwned_enabled"`
	MaxBreaches                      types.Int64  `tfsdk:"max_breaches" json:"max_breaches"`
	IgnoreNetworkErrors              types.Bool   `tfsdk:"ignore_network_errors" json:"ignore_network_errors"`
	MinPasswordLength                types.Int64  `tfsdk:"min_password_length" json:"min_password_length"`
	IdentifierSimilarityCheckEnabled types.Bool   `tfsdk:"identifier_similarity_check_enabled" json:"identifier_similarity_check_enabled"`
	HaveIBeenPwnedHost               types.String `tfsdk:"haveibeenpwned_host" json:"haveibeenpwned_host"`
}

type IdentitySelfServiceMethodsPasswordType struct {
	Enabled types.Bool                                    `tfsdk:"enabled" json:"enabled"`
	Config  *IdentitySelfServiceMethodsPasswordConfigType `tfsdk:"config" json:"config,omitempty"`
}

type IdentitySelfServiceMethodsTotpConfigType struct {
	Issuer types.String `tfsdk:"issuer" json:"issuer"`
}

type IdentitySelfServiceMethodsTotpType struct {
	Enabled types.Bool                                `tfsdk:"enabled" json:"enabled"`
	Config  *IdentitySelfServiceMethodsTotpConfigType `tfsdk:"config" json:"config,omitempty"`
}

type IdentitySelfServiceMethodsLookupSecretType struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type IdentitySelfServiceMethodsProfileType struct {
	Enabled types.Bool `tfsdk:"enabled" json:"enabled"`
}

type IdentitySelfServiceMethodsWebAuthNConfigRpType struct {
	Id          types.String `tfsdk:"id" json:"id"`
	DisplayName types.String `tfsdk:"display_name" json:"display_name"`
	Icon        types.String `tfsdk:"icon" json:"icon"`
}

type IdentitySelfServiceMethodsWebAuthNConfigType struct {
	Passwordless types.Bool                                      `tfsdk:"passwordless" json:"passwordless"`
	Rp           *IdentitySelfServiceMethodsWebAuthNConfigRpType `tfsdk:"rp" json:"rp,omitempty"`
}

type IdentitySelfServiceMethodsWebAuthNType struct {
	Enabled types.Bool                                    `tfsdk:"enabled" json:"enabled"`
	Config  *IdentitySelfServiceMethodsWebAuthNConfigType `tfsdk:"config" json:"config,omitempty"`
}

type IdentitySelfServiceMethodsOidcConfigProviderRequestedClaimsType struct {
	IdToken []types.String `tfsdk:"id_token" json:"id_token"`
}

type IdentitySelfServiceMethodsOidcConfigProviderType struct {
	Id                         types.String                                                     `tfsdk:"id" json:"id"`
	Provider                   types.String                                                     `tfsdk:"provider" json:"provider"`
	ClientId                   types.String                                                     `tfsdk:"client_id" json:"client_id"`
	MapperUrl                  types.String                                                     `tfsdk:"mapper_url" json:"mapper_url"`
	ClientSecret               types.String                                                     `tfsdk:"client_secret" json:"client_secret"`
	IssuerUrl                  types.String                                                     `tfsdk:"issuer_url" json:"issuer_url"`
	AuthUrl                    types.String                                                     `tfsdk:"auth_url" json:"auth_url"`
	TokenUrl                   types.String                                                     `tfsdk:"token_url" json:"token_url"`
	Scope                      []types.String                                                   `tfsdk:"scope" json:"scope"`
	MicrosoftTenant            types.String                                                     `tfsdk:"microsoft_tenant" json:"microsoft_tenant"`
	SubjectSource              types.String                                                     `tfsdk:"subject_source" json:"subject_source"`
	AppleTeamId                types.String                                                     `tfsdk:"apple_team_id" json:"apple_team_id"`
	ApplePrivateKeyId          types.String                                                     `tfsdk:"apple_private_key_id" json:"apple_private_key_id"`
	ApplePrivateKey            types.String                                                     `tfsdk:"apple_private_key" json:"apple_private_key"`
	RequestedClaims            *IdentitySelfServiceMethodsOidcConfigProviderRequestedClaimsType `tfsdk:"requested_claims" json:"requested_claims,omitempty"`
	OrganizationId             types.String                                                     `tfsdk:"organization_id" json:"organization_id"`
	Label                      types.String                                                     `tfsdk:"label" json:"label"`
	AdditionalIdTokenAudiences []types.String                                                   `tfsdk:"additional_id_token_audiences" json:"additional_id_token_audiences"`
}

type IdentitySelfServiceMethodsOidcConfigType struct {
	BaseRedirectUri types.String                                       `tfsdk:"base_redirect_uri" json:"base_redirect_uri"`
	Providers       []IdentitySelfServiceMethodsOidcConfigProviderType `tfsdk:"providers" json:"providers"`
}

type IdentitySelfServiceMethodsOidcType struct {
	Enabled types.Bool                                `tfsdk:"enabled" json:"enabled"`
	Config  *IdentitySelfServiceMethodsOidcConfigType `tfsdk:"config" json:"config,omitempty"`
}

type IdentitySelfServiceMethodsType struct {
	Link         *IdentitySelfServiceMethodsLinkType         `tfsdk:"link" json:"link,omitempty"`
	Code         *IdentitySelfServiceMethodsCodeType         `tfsdk:"code" json:"code,omitempty"`
	Password     *IdentitySelfServiceMethodsPasswordType     `tfsdk:"password" json:"password,omitempty"`
	Totp         *IdentitySelfServiceMethodsTotpType         `tfsdk:"totp" json:"totp,omitempty"`
	LookupSecret *IdentitySelfServiceMethodsLookupSecretType `tfsdk:"lookup_secret" json:"lookup_secret,omitempty"`
	Profile      *IdentitySelfServiceMethodsProfileType      `tfsdk:"profile" json:"profile,omitempty"`
	WebAuthN     *IdentitySelfServiceMethodsWebAuthNType     `tfsdk:"webauthn" json:"webauthn,omitempty"`
	Oidc         *IdentitySelfServiceMethodsOidcType         `tfsdk:"oidc" json:"oidc,omitempty"`
}

type IdentitySelfServiceFlowsLogoutAfterType struct {
	DefaultBrowserReturnUrl types.String `tfsdk:"default_browser_return_url" json:"default_browser_return_url"`
}

type IdentitySelfServiceFlowsLogoutType struct {
	After *IdentitySelfServiceFlowsLogoutAfterType `tfsdk:"after" json:"after,omitempty"`
}

type IdentitySelfServiceFlowsErrorType struct {
	UiUrl types.String `tfsdk:"ui_url" json:"ui_url"`
}

type IdentitySelfServiceFlowsRegistrationType struct {
	LoginHints types.Bool   `tfsdk:"login_hints" json:"login_hints"`
	UiUrl      types.String `tfsdk:"ui_url" json:"ui_url"`
	Lifespan   types.String `tfsdk:"lifespan" json:"lifespan"`
	Enabled    types.Bool   `tfsdk:"enabled" json:"enabled"`
}

type IdentitySelfServiceFlowsLoginType struct {
	UiUrl    types.String `tfsdk:"ui_url" json:"ui_url"`
	Lifespan types.String `tfsdk:"lifespan" json:"lifespan"`
}

type IdentitySelfServiceFlowsVerificationType struct {
	UiUrl                   types.String `tfsdk:"ui_url" json:"ui_url"`
	Lifespan                types.String `tfsdk:"lifespan" json:"lifespan"`
	Use                     types.String `tfsdk:"use" json:"use"`
	NotifyUnknownRecipients types.Bool   `tfsdk:"notify_unknown_recipients" json:"notify_unknown_recipients"`
	Enabled                 types.Bool   `tfsdk:"enabled" json:"enabled"`
}

type IdentitySelfServiceFlowsRecoveryType struct {
	UiUrl                   types.String `tfsdk:"ui_url" json:"ui_url"`
	Lifespan                types.String `tfsdk:"lifespan" json:"lifespan"`
	Use                     types.String `tfsdk:"use" json:"use"`
	NotifyUnknownRecipients types.Bool   `tfsdk:"notify_unknown_recipients" json:"notify_unknown_recipients"`
	Enabled                 types.Bool   `tfsdk:"enabled" json:"enabled"`
}

type IdentitySelfServiceFlowsSettingsType struct {
	UiUrl                   types.String `tfsdk:"ui_url" json:"ui_url"`
	Lifespan                types.String `tfsdk:"lifespan" json:"lifespan"`
	PrivilegedSessionMaxAge types.String `tfsdk:"privileged_session_max_age" json:"privileged_session_max_age"`
	RequiredAal             types.String `tfsdk:"required_aal" json:"required_aal"`
}

type IdentitySelfServiceFlowsType struct {
	Logout       *IdentitySelfServiceFlowsLogoutType       `tfsdk:"logout" json:"logout,omitempty"`
	Error        *IdentitySelfServiceFlowsErrorType        `tfsdk:"error" json:"error,omitempty"`
	Registration *IdentitySelfServiceFlowsRegistrationType `tfsdk:"registration" json:"registration,omitempty"`
	Login        *IdentitySelfServiceFlowsLoginType        `tfsdk:"login" json:"login,omitempty"`
	Verification *IdentitySelfServiceFlowsVerificationType `tfsdk:"verification" json:"verification,omitempty"`
	Recovery     *IdentitySelfServiceFlowsRecoveryType     `tfsdk:"recovery" json:"recovery,omitempty"`
	Settings     *IdentitySelfServiceFlowsSettingsType     `tfsdk:"settings" json:"settings,omitempty"`
}

type IdentitySelfServiceType struct {
	DefaultBrowserReturnUrl types.String                    `tfsdk:"default_browser_return_url" json:"default_browser_return_url"`
	AllowedReturnUrls       []types.String                  `tfsdk:"allowed_return_urls" json:"allowed_return_urls"`
	Methods                 *IdentitySelfServiceMethodsType `tfsdk:"methods" json:"methods,omitempty"`
	Flows                   *IdentitySelfServiceFlowsType   `tfsdk:"flows" json:"flows,omitempty"`
}

type IdentityConfigType struct {
	Identity    IdentityIdentityType    `tfsdk:"identity" json:"identity"`
	SelfService IdentitySelfServiceType `tfsdk:"selfservice" json:"selfservice"`
}

type IdentityType struct {
	Config IdentityConfigType `tfsdk:"config" json:"config"`
}

func marshalWithoutNulls(jsonObj map[string]interface{}) ([]byte, error) {
	for k, v := range jsonObj {
		// Remove null strings from the payload
		if reflect.TypeOf(v).Kind() == reflect.Pointer && reflect.ValueOf(v).IsNil() {
			delete(jsonObj, k)
		}
	}
	return json.Marshal(jsonObj)
}

// MarshalJSON For json.Marshal compatibility.
func (schema *IdentityIdentitySchemaType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"id":  schema.Id.ValueStringPointer(),
		"url": schema.Url.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (identity *IdentityIdentityType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"default_schema_id": identity.DefaultSchemaId.ValueStringPointer(),
		"schemas":           identity.Schemas,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (config *IdentitySelfServiceMethodsLinkConfigType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"lifespan": config.Lifespan.ValueStringPointer(),
		"base_url": config.BaseUrl.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (link *IdentitySelfServiceMethodsLinkType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"config":  link.Config,
		"enabled": link.Enabled.ValueBool(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (config *IdentitySelfServiceMethodsCodeConfigType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"lifespan": config.Lifespan.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (code *IdentitySelfServiceMethodsCodeType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"passwordless_login_fallback_enabled": code.PasswordlessLoginFallbackEnabled.ValueBool(),
		"enabled":                             code.Enabled.ValueBool(),
		"passwordless_enabled":                code.Enabled.ValueBool(),
		"config":                              code.Config,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (config *IdentitySelfServiceMethodsPasswordConfigType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"haveibeenpwned_enabled":              config.HaveIBeenPwnedEnabled.ValueBool(),
		"max_breaches":                        config.MaxBreaches.ValueInt64Pointer(),
		"ignore_network_errors":               config.IgnoreNetworkErrors.ValueBool(),
		"min_password_length":                 config.MinPasswordLength.ValueInt64Pointer(),
		"identifier_similarity_check_enabled": config.IdentifierSimilarityCheckEnabled.ValueBool(),
		"haveibeenpwned_host":                 config.HaveIBeenPwnedHost.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (password *IdentitySelfServiceMethodsPasswordType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"enabled": password.Enabled.ValueBool(),
		"config":  password.Config,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (config *IdentitySelfServiceMethodsTotpConfigType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"issuer": config.Issuer.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (totp *IdentitySelfServiceMethodsTotpType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"enabled": totp.Enabled.ValueBool(),
		"config":  totp.Config,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (lookupSecret *IdentitySelfServiceMethodsLookupSecretType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"enabled": lookupSecret.Enabled.ValueBool(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (profile *IdentitySelfServiceMethodsProfileType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"enabled": profile.Enabled.ValueBool(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (rp *IdentitySelfServiceMethodsWebAuthNConfigRpType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"id":           rp.Id.ValueStringPointer(),
		"display_name": rp.DisplayName.ValueStringPointer(),
		"icon":         rp.Icon.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (config *IdentitySelfServiceMethodsWebAuthNConfigType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"passwordless": config.Passwordless.ValueBool(),
		"rp":           config.Rp,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (webAuthN *IdentitySelfServiceMethodsWebAuthNType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"enabled": webAuthN.Enabled.ValueBool(),
		"config":  webAuthN.Config,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (requestedClaims *IdentitySelfServiceMethodsOidcConfigProviderRequestedClaimsType) MarshalJSON() ([]byte, error) {
	idTokens := make([]*string, 0)
	for _, idToken := range requestedClaims.IdToken {
		idTokens = append(idTokens, idToken.ValueStringPointer())
	}

	return marshalWithoutNulls(map[string]interface{}{
		"id_token": idTokens,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (provider *IdentitySelfServiceMethodsOidcConfigProviderType) MarshalJSON() ([]byte, error) {
	scopes := make([]*string, 0)
	for _, scope := range provider.Scope {
		scopes = append(scopes, scope.ValueStringPointer())
	}

	additionalTokenAudiences := make([]*string, 0)
	for _, audience := range provider.AdditionalIdTokenAudiences {
		additionalTokenAudiences = append(additionalTokenAudiences, audience.ValueStringPointer())
	}

	return marshalWithoutNulls(map[string]interface{}{
		"id":                            provider.Id.ValueStringPointer(),
		"provider":                      provider.Provider.ValueStringPointer(),
		"client_id":                     provider.ClientId.ValueStringPointer(),
		"mapper_url":                    provider.MapperUrl.ValueStringPointer(),
		"client_secret":                 provider.ClientSecret.ValueStringPointer(),
		"issuer_url":                    provider.IssuerUrl.ValueStringPointer(),
		"auth_url":                      provider.AuthUrl.ValueStringPointer(),
		"token_url":                     provider.TokenUrl.ValueStringPointer(),
		"scope":                         scopes,
		"microsoft_tenant":              provider.MicrosoftTenant.ValueStringPointer(),
		"subject_source":                provider.SubjectSource.ValueStringPointer(),
		"apple_team_id":                 provider.AppleTeamId.ValueStringPointer(),
		"apple_private_key_id":          provider.ApplePrivateKeyId.ValueStringPointer(),
		"apple_private_key":             provider.ApplePrivateKey.ValueStringPointer(),
		"requested_claims":              provider.RequestedClaims,
		"organization_id":               provider.OrganizationId.ValueStringPointer(),
		"label":                         provider.Label.ValueStringPointer(),
		"additional_id_token_audiences": additionalTokenAudiences,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (config *IdentitySelfServiceMethodsOidcConfigType) MarshalJSON() ([]byte, error) {
	jsonObj := map[string]interface{}{
		"base_redirect_uri": config.BaseRedirectUri.ValueStringPointer(),
		"providers":         config.Providers,
	}
	for k, v := range jsonObj {
		if v == nil {
			delete(jsonObj, k)
		}
	}
	return marshalWithoutNulls(jsonObj)
}

// MarshalJSON For json.Marshal compatibility.
func (oidc *IdentitySelfServiceMethodsOidcType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"enabled": oidc.Enabled.ValueBool(),
		"config":  oidc.Config,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (methods *IdentitySelfServiceMethodsType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"link":          methods.Link,
		"code":          methods.Code,
		"password":      methods.Password,
		"totp":          methods.Totp,
		"lookup_secret": methods.LookupSecret,
		"profile":       methods.Profile,
		"webauthn":      methods.WebAuthN,
		"oidc":          methods.Oidc,
	})
}

// MarshalJSON For json.Marshal compatibility.
func (after *IdentitySelfServiceFlowsLogoutAfterType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"default_browser_return_url": after.DefaultBrowserReturnUrl.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (flowsError *IdentitySelfServiceFlowsErrorType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"ui_url": flowsError.UiUrl.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (registration *IdentitySelfServiceFlowsRegistrationType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"login_hints": registration.LoginHints.ValueBool(),
		"ui_url":      registration.UiUrl.ValueStringPointer(),
		"lifespan":    registration.Lifespan.ValueStringPointer(),
		"enabled":     registration.Enabled.ValueBool(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (login *IdentitySelfServiceFlowsLoginType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"ui_url":   login.UiUrl.ValueStringPointer(),
		"lifespan": login.Lifespan.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (verification *IdentitySelfServiceFlowsVerificationType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"ui_url":                    verification.UiUrl.ValueStringPointer(),
		"lifespan":                  verification.Lifespan.ValueStringPointer(),
		"use":                       verification.Use.ValueStringPointer(),
		"notify_unknown_recipients": verification.NotifyUnknownRecipients.ValueBool(),
		"enabled":                   verification.Enabled.ValueBool(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (recovery *IdentitySelfServiceFlowsRecoveryType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"ui_url":                    recovery.UiUrl.ValueStringPointer(),
		"lifespan":                  recovery.Lifespan.ValueStringPointer(),
		"use":                       recovery.Use.ValueStringPointer(),
		"notify_unknown_recipients": recovery.NotifyUnknownRecipients.ValueBool(),
		"enabled":                   recovery.Enabled.ValueBool(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (settings *IdentitySelfServiceFlowsSettingsType) MarshalJSON() ([]byte, error) {
	return marshalWithoutNulls(map[string]interface{}{
		"ui_url":                     settings.UiUrl.ValueStringPointer(),
		"lifespan":                   settings.Lifespan.ValueStringPointer(),
		"privileged_session_max_age": settings.PrivilegedSessionMaxAge.ValueStringPointer(),
		"required_aal":               settings.RequiredAal.ValueStringPointer(),
	})
}

// MarshalJSON For json.Marshal compatibility.
func (selfService *IdentitySelfServiceType) MarshalJSON() ([]byte, error) {
	allowedReturnUrls := make([]*string, 0)
	for _, url := range selfService.AllowedReturnUrls {
		allowedReturnUrls = append(allowedReturnUrls, url.ValueStringPointer())
	}

	value := map[string]interface{}{
		"default_browser_return_url": selfService.DefaultBrowserReturnUrl.ValueStringPointer(),
		"allowed_return_urls":        allowedReturnUrls,
		"methods":                    selfService.Methods,
		"flows":                      selfService.Flows,
	}
	if value["methods"].(*IdentitySelfServiceMethodsType) == nil {
		delete(value, "methods")
	}
	if value["flows"].(*IdentitySelfServiceFlowsType) == nil {
		delete(value, "flows")
	}

	return marshalWithoutNulls(value)
}

func (config *IdentityConfigType) ToApiRepresentation() (map[string]interface{}, error) {
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

func NewProjectIdentityFromApiRepresentation(apiIdentity *ory.ProjectServiceIdentity) *IdentityType {
	rawPermissionConfig := apiIdentity.Config
	rawIdentity := rawPermissionConfig["identity"].(map[string]interface{})
	rawSelfService := rawPermissionConfig["selfservice"].(map[string]interface{})

	identitySchemas := make([]IdentityIdentitySchemaType, 0)
	for _, rawIdentitySchema := range rawIdentity["schemas"].([]interface{}) {
		schema := rawIdentitySchema.(map[string]interface{})
		identitySchemas = append(identitySchemas, IdentityIdentitySchemaType{
			Id:  types.StringValue(schema["id"].(string)),
			Url: types.StringValue(schema["url"].(string)),
		})
	}

	identity := IdentityIdentityType{
		DefaultSchemaId: types.StringValue(rawIdentity["default_schema_id"].(string)),
		Schemas:         identitySchemas,
	}

	selfServiceAllowedReturnUrls := make([]types.String, 0)
	for _, allowedReturnUrl := range rawSelfService["allowed_return_urls"].([]interface{}) {
		selfServiceAllowedReturnUrls = append(selfServiceAllowedReturnUrls, types.StringValue(allowedReturnUrl.(string)))
	}

	rawSelfServiceMethods := rawSelfService["methods"].(map[string]interface{})

	rawSelfServiceMethodsLink := rawSelfServiceMethods["link"].(map[string]interface{})
	rawSelfServiceMethodsLinkConfig := rawSelfServiceMethodsLink["config"].(map[string]interface{})

	rawSelfServiceMethodsCode := rawSelfServiceMethods["code"].(map[string]interface{})
	rawSelfServiceMethodsCodeConfig := rawSelfServiceMethodsCode["config"].(map[string]interface{})

	rawSelfServiceMethodsPassword := rawSelfServiceMethods["password"].(map[string]interface{})
	rawSelfServiceMethodsPasswordConfig := rawSelfServiceMethodsPassword["config"].(map[string]interface{})

	rawSelfServiceMethodsTotp := rawSelfServiceMethods["totp"].(map[string]interface{})
	rawSelfServiceMethodsTotpConfig := rawSelfServiceMethodsTotp["config"].(map[string]interface{})

	rawSelfServiceMethodsLookupSecret := rawSelfServiceMethods["lookup_secret"].(map[string]interface{})

	rawSelfServiceMethodsProfile := rawSelfServiceMethods["profile"].(map[string]interface{})

	rawSelfServiceMethodsWebAuthN := rawSelfServiceMethods["webauthn"].(map[string]interface{})
	rawSelfServiceMethodsWebAuthNConfig := rawSelfServiceMethodsWebAuthN["config"].(map[string]interface{})
	rawSelfServiceMethodsWebAuthNConfigRp := rawSelfServiceMethodsWebAuthNConfig["rp"].(map[string]interface{})

	rawSelfServiceMethodsOidc := rawSelfServiceMethods["oidc"].(map[string]interface{})
	rawSelfServiceMethodsOidcConfig := rawSelfServiceMethodsOidc["config"].(map[string]interface{})
	selfServiceMethodsOidcConfigProviders := make([]IdentitySelfServiceMethodsOidcConfigProviderType, 0)
	for _, rawProvider := range rawSelfServiceMethodsOidcConfig["providers"].([]interface{}) {
		rawProviderTyped := rawProvider.(map[string]interface{})

		scopes := make([]types.String, 0)
		for _, scope := range rawProviderTyped["scope"].([]interface{}) {
			scopes = append(scopes, types.StringValue(scope.(string)))
		}

		requestedClaimsIdToken := make([]types.String, 0)
		if rawProviderTyped["requested_claims"] != nil {
			rawIdTokenClaims := rawProviderTyped["requested_claims"].(map[string]interface{})["id_token"].(map[string]interface{})
			for claim := range rawIdTokenClaims {
				requestedClaimsIdToken = append(requestedClaimsIdToken, types.StringValue(claim))
			}
		}

		additionalIdTokenAudiences := make([]types.String, 0)
		if rawProviderTyped["additional_id_token_audiences"] != nil {
			for _, audience := range rawProviderTyped["additional_id_token_audiences"].([]interface{}) {
				additionalIdTokenAudiences = append(additionalIdTokenAudiences, types.StringValue(audience.(string)))
			}
		}

		clientSecret := types.StringNull()
		if rawProviderTyped["client_secret"] != nil {
			clientSecret = types.StringValue(rawProviderTyped["client_secret"].(string))
		}

		issuerUrl := types.StringNull()
		if rawProviderTyped["issuer_url"] != nil {
			issuerUrl = types.StringValue(rawProviderTyped["issuer_url"].(string))
		}

		authUrl := types.StringNull()
		if rawProviderTyped["auth_url"] != nil {
			authUrl = types.StringValue(rawProviderTyped["auth_url"].(string))
		}

		tokenUrl := types.StringNull()
		if rawProviderTyped["token_url"] != nil {
			tokenUrl = types.StringValue(rawProviderTyped["token_url"].(string))
		}

		microsoftTenant := types.StringNull()
		if rawProviderTyped["microsoft_tenant"] != nil {
			microsoftTenant = types.StringValue(rawProviderTyped["microsoft_tenant"].(string))
		}

		subjectSource := types.StringNull()
		if rawProviderTyped["subject_source"] != nil {
			subjectSource = types.StringValue(rawProviderTyped["subject_source"].(string))
		}

		appleTeamId := types.StringNull()
		if rawProviderTyped["apple_team_id"] != nil {
			appleTeamId = types.StringValue(rawProviderTyped["apple_team_id"].(string))
		}

		applePrivateKeyId := types.StringNull()
		if rawProviderTyped["apple_private_key_id"] != nil {
			applePrivateKeyId = types.StringValue(rawProviderTyped["apple_private_key_id"].(string))
		}

		applePrivateKey := types.StringNull()
		if rawProviderTyped["apple_private_key"] != nil {
			applePrivateKey = types.StringValue(rawProviderTyped["apple_private_key"].(string))
		}

		organizationId := types.StringNull()
		if rawProviderTyped["organization_id"] != nil {
			organizationId = types.StringValue(rawProviderTyped["organization_id"].(string))
		}

		label := types.StringNull()
		if rawProviderTyped["label"] != nil {
			label = types.StringValue(rawProviderTyped["label"].(string))
		}

		selfServiceMethodsOidcConfigProviders = append(
			selfServiceMethodsOidcConfigProviders,
			IdentitySelfServiceMethodsOidcConfigProviderType{
				Id:                types.StringValue(rawProviderTyped["id"].(string)),
				Provider:          types.StringValue(rawProviderTyped["provider"].(string)),
				ClientId:          types.StringValue(rawProviderTyped["client_id"].(string)),
				MapperUrl:         types.StringValue(rawProviderTyped["mapper_url"].(string)),
				ClientSecret:      clientSecret,
				IssuerUrl:         issuerUrl,
				AuthUrl:           authUrl,
				TokenUrl:          tokenUrl,
				Scope:             scopes,
				MicrosoftTenant:   microsoftTenant,
				SubjectSource:     subjectSource,
				AppleTeamId:       appleTeamId,
				ApplePrivateKeyId: applePrivateKeyId,
				ApplePrivateKey:   applePrivateKey,
				RequestedClaims: &IdentitySelfServiceMethodsOidcConfigProviderRequestedClaimsType{
					IdToken: requestedClaimsIdToken,
				},
				OrganizationId:             organizationId,
				Label:                      label,
				AdditionalIdTokenAudiences: additionalIdTokenAudiences,
			},
		)
	}

	selfServiceMethodsLink := &IdentitySelfServiceMethodsLinkType{
		Config: &IdentitySelfServiceMethodsLinkConfigType{},
	}
	if rawSelfServiceMethodsLink["enabled"] != nil {
		selfServiceMethodsLink.Enabled = types.BoolValue(rawSelfServiceMethodsLink["enabled"].(bool))
	}
	if rawSelfServiceMethodsLinkConfig["lifespan"] != nil {
		selfServiceMethodsLink.Config.Lifespan = types.StringValue(rawSelfServiceMethodsLinkConfig["lifespan"].(string))
	}
	if rawSelfServiceMethodsLinkConfig["base_url"] != nil {
		selfServiceMethodsLink.Config.BaseUrl = types.StringValue(rawSelfServiceMethodsLinkConfig["base_url"].(string))
	}

	selfServiceMethodsCode := &IdentitySelfServiceMethodsCodeType{
		Config: &IdentitySelfServiceMethodsCodeConfigType{},
	}
	if rawSelfServiceMethodsCode["passwordless_login_fallback_enabled"] != nil {
		selfServiceMethodsCode.PasswordlessLoginFallbackEnabled = types.BoolValue(rawSelfServiceMethodsCode["passwordless_login_fallback_enabled"].(bool))
	}
	if rawSelfServiceMethodsCode["enabled"] != nil {
		selfServiceMethodsCode.Enabled = types.BoolValue(rawSelfServiceMethodsCode["enabled"].(bool))
	}
	if rawSelfServiceMethodsCode["passwordless_enabled"] != nil {
		selfServiceMethodsCode.PasswordlessEnabled = types.BoolValue(rawSelfServiceMethodsCode["passwordless_enabled"].(bool))
	}
	if rawSelfServiceMethodsCodeConfig["lifespan"] != nil {
		selfServiceMethodsCode.Config.Lifespan = types.StringValue(rawSelfServiceMethodsCodeConfig["lifespan"].(string))
	}

	selfServiceMethodsPassword := &IdentitySelfServiceMethodsPasswordType{
		Config: &IdentitySelfServiceMethodsPasswordConfigType{},
	}
	if rawSelfServiceMethodsPassword["enabled"] != nil {
		selfServiceMethodsPassword.Enabled = types.BoolValue(rawSelfServiceMethodsPassword["enabled"].(bool))
	}
	if rawSelfServiceMethodsPasswordConfig["haveibeenpwned_enabled"] != nil {
		selfServiceMethodsPassword.Config.HaveIBeenPwnedEnabled = types.BoolValue(rawSelfServiceMethodsPasswordConfig["haveibeenpwned_enabled"].(bool))
	}
	if rawSelfServiceMethodsPasswordConfig["max_breaches"] != nil {
		selfServiceMethodsPassword.Config.MaxBreaches = types.Int64Value(int64(int(math.Round(rawSelfServiceMethodsPasswordConfig["max_breaches"].(float64)))))
	}
	if rawSelfServiceMethodsPasswordConfig["ignore_network_errors"] != nil {
		selfServiceMethodsPassword.Config.IgnoreNetworkErrors = types.BoolValue(rawSelfServiceMethodsPasswordConfig["ignore_network_errors"].(bool))
	}
	if rawSelfServiceMethodsPasswordConfig["min_password_length"] != nil {
		selfServiceMethodsPassword.Config.MinPasswordLength = types.Int64Value(int64(int(math.Round(rawSelfServiceMethodsPasswordConfig["min_password_length"].(float64)))))
	}
	if rawSelfServiceMethodsPasswordConfig["identifier_similarity_check_enabled"] != nil {
		selfServiceMethodsPassword.Config.IdentifierSimilarityCheckEnabled = types.BoolValue(rawSelfServiceMethodsPasswordConfig["identifier_similarity_check_enabled"].(bool))
	}
	if rawSelfServiceMethodsPasswordConfig["haveibeenpwned_host"] != nil {
		selfServiceMethodsPassword.Config.HaveIBeenPwnedHost = types.StringValue(rawSelfServiceMethodsPasswordConfig["haveibeenpwned_host"].(string))
	}

	selfServiceMethodsTotp := &IdentitySelfServiceMethodsTotpType{
		Config: &IdentitySelfServiceMethodsTotpConfigType{},
	}
	if rawSelfServiceMethodsTotp["enabled"] != nil {
		selfServiceMethodsTotp.Enabled = types.BoolValue(rawSelfServiceMethodsTotp["enabled"].(bool))
	}
	if rawSelfServiceMethodsTotpConfig["issuer"] != nil {
		selfServiceMethodsTotp.Config.Issuer = types.StringValue(rawSelfServiceMethodsTotpConfig["issuer"].(string))
	}

	selfServiceMethodsLookupSecret := &IdentitySelfServiceMethodsLookupSecretType{}
	if rawSelfServiceMethodsLookupSecret["enabled"] != nil {
		selfServiceMethodsLookupSecret.Enabled = types.BoolValue(rawSelfServiceMethodsLookupSecret["enabled"].(bool))
	}

	selfServiceMethodsProfile := &IdentitySelfServiceMethodsProfileType{}
	if rawSelfServiceMethodsProfile["enabled"] != nil {
		selfServiceMethodsProfile.Enabled = types.BoolValue(rawSelfServiceMethodsProfile["enabled"].(bool))
	}

	selfServiceMethodsWebAuthN := &IdentitySelfServiceMethodsWebAuthNType{
		Config: &IdentitySelfServiceMethodsWebAuthNConfigType{
			Rp: &IdentitySelfServiceMethodsWebAuthNConfigRpType{},
		},
	}
	if rawSelfServiceMethodsWebAuthN["enabled"] != nil {
		selfServiceMethodsWebAuthN.Enabled = types.BoolValue(rawSelfServiceMethodsWebAuthN["enabled"].(bool))
	}
	if rawSelfServiceMethodsWebAuthNConfig["passwordless"] != nil {
		selfServiceMethodsWebAuthN.Config.Passwordless = types.BoolValue(rawSelfServiceMethodsWebAuthNConfig["passwordless"].(bool))
	}
	if rawSelfServiceMethodsWebAuthNConfigRp["id"] != nil {
		selfServiceMethodsWebAuthN.Config.Rp.Id = types.StringValue(rawSelfServiceMethodsWebAuthNConfigRp["id"].(string))
	}
	if rawSelfServiceMethodsWebAuthNConfigRp["display_name"] != nil {
		selfServiceMethodsWebAuthN.Config.Rp.DisplayName = types.StringValue(rawSelfServiceMethodsWebAuthNConfigRp["display_name"].(string))
	}
	if rawSelfServiceMethodsWebAuthNConfigRp["icon"] != nil {
		selfServiceMethodsWebAuthN.Config.Rp.Icon = types.StringValue(rawSelfServiceMethodsWebAuthNConfigRp["icon"].(string))
	}

	selfServiceMethodsOidc := &IdentitySelfServiceMethodsOidcType{
		Config: &IdentitySelfServiceMethodsOidcConfigType{
			Providers: selfServiceMethodsOidcConfigProviders,
		},
	}
	if rawSelfServiceMethodsOidc["enabled"] != nil {
		selfServiceMethodsOidc.Enabled = types.BoolValue(rawSelfServiceMethodsOidc["enabled"].(bool))
	}
	if rawSelfServiceMethodsOidcConfig["base_redirect_uri"] != nil {
		selfServiceMethodsOidc.Config.BaseRedirectUri = types.StringValue(rawSelfServiceMethodsOidcConfig["base_redirect_uri"].(string))
	}

	selfServiceMethods := IdentitySelfServiceMethodsType{
		Link:         selfServiceMethodsLink,
		Code:         selfServiceMethodsCode,
		Password:     selfServiceMethodsPassword,
		Totp:         selfServiceMethodsTotp,
		LookupSecret: selfServiceMethodsLookupSecret,
		Profile:      selfServiceMethodsProfile,
		WebAuthN:     selfServiceMethodsWebAuthN,
		Oidc:         selfServiceMethodsOidc,
	}

	selfService := IdentitySelfServiceType{
		DefaultBrowserReturnUrl: types.StringValue(rawSelfService["default_browser_return_url"].(string)),
		AllowedReturnUrls:       selfServiceAllowedReturnUrls,
		// TODO: Implement these
		Methods: &selfServiceMethods,
		Flows:   nil,
	}

	return &IdentityType{
		Config: IdentityConfigType{
			Identity:    identity,
			SelfService: selfService,
		},
	}
}

func (identity *IdentityType) ToApiRepresentation() (*ory.ProjectServiceIdentity, error) {
	config, err := identity.Config.ToApiRepresentation()
	if err != nil {
		return nil, err
	}
	return ory.NewProjectServiceIdentity(config), nil
}

func (schema *IdentityIdentitySchemaType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":  types.StringType,
			"url": types.StringType,
		},
	}
}

func (identity *IdentityIdentityType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_schema_id": types.StringType,
			"schemas": types.ListType{
				ElemType: (&IdentityIdentitySchemaType{}).TerraformType(),
			},
		},
	}
}

func (config *IdentityConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"identity":    (&IdentityIdentityType{}).TerraformType(),
			"selfservice": (&IdentitySelfServiceType{}).TerraformType(),
		},
	}
}

func (config *IdentitySelfServiceMethodsLinkConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"lifespan": types.StringType,
			"base_url": types.StringType,
		},
	}
}

func (link *IdentitySelfServiceMethodsLinkType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config":  (&IdentitySelfServiceMethodsLinkConfigType{}).TerraformType(),
			"enabled": types.BoolType,
		},
	}
}

func (config *IdentitySelfServiceMethodsCodeConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"lifespan": types.StringType,
		},
	}
}

func (code *IdentitySelfServiceMethodsCodeType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"passwordless_login_fallback_enabled": types.BoolType,
			"enabled":                             types.BoolType,
			"passwordless_enabled":                types.BoolType,
			"config":                              (&IdentitySelfServiceMethodsCodeConfigType{}).TerraformType(),
		},
	}
}

func (config *IdentitySelfServiceMethodsPasswordConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"haveibeenpwned_enabled":              types.BoolType,
			"max_breaches":                        types.Int64Type,
			"ignore_network_errors":               types.BoolType,
			"min_password_length":                 types.Int64Type,
			"identifier_similarity_check_enabled": types.BoolType,
			"haveibeenpwned_host":                 types.StringType,
		},
	}
}

func (password *IdentitySelfServiceMethodsPasswordType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
			"config":  (&IdentitySelfServiceMethodsPasswordConfigType{}).TerraformType(),
		},
	}
}

func (config *IdentitySelfServiceMethodsTotpConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"issuer": types.StringType,
		},
	}
}

func (totp *IdentitySelfServiceMethodsTotpType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
			"config":  (&IdentitySelfServiceMethodsTotpConfigType{}).TerraformType(),
		},
	}
}

func (lookupSecret *IdentitySelfServiceMethodsLookupSecretType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

func (profile *IdentitySelfServiceMethodsProfileType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
		},
	}
}

func (rp *IdentitySelfServiceMethodsWebAuthNConfigRpType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":           types.StringType,
			"display_name": types.StringType,
			"icon":         types.StringType,
		},
	}
}

func (config *IdentitySelfServiceMethodsWebAuthNConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"passwordless": types.BoolType,
			"rp":           (&IdentitySelfServiceMethodsWebAuthNConfigRpType{}).TerraformType(),
		},
	}
}

func (webAuthN *IdentitySelfServiceMethodsWebAuthNType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
			"config":  (&IdentitySelfServiceMethodsWebAuthNConfigType{}).TerraformType(),
		},
	}
}

func (requestedClaims *IdentitySelfServiceMethodsOidcConfigProviderRequestedClaimsType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id_token": types.ListType{
				ElemType: types.StringType,
			},
		},
	}
}

func (provider *IdentitySelfServiceMethodsOidcConfigProviderType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"id":                            types.StringType,
			"provider":                      types.StringType,
			"client_id":                     types.StringType,
			"mapper_url":                    types.StringType,
			"client_secret":                 types.StringType,
			"issuer_url":                    types.StringType,
			"auth_url":                      types.StringType,
			"token_url":                     types.StringType,
			"scope":                         types.ListType{ElemType: types.StringType},
			"microsoft_tenant":              types.StringType,
			"subject_source":                types.StringType,
			"apple_team_id":                 types.StringType,
			"apple_private_key_id":          types.StringType,
			"apple_private_key":             types.StringType,
			"requested_claims":              (&IdentitySelfServiceMethodsOidcConfigProviderRequestedClaimsType{}).TerraformType(),
			"organization_id":               types.StringType,
			"label":                         types.StringType,
			"additional_id_token_audiences": types.ListType{ElemType: types.StringType},
		},
	}
}

func (config *IdentitySelfServiceMethodsOidcConfigType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"base_redirect_uri": types.StringType,
			"providers":         types.ListType{ElemType: (&IdentitySelfServiceMethodsOidcConfigProviderType{}).TerraformType()},
		},
	}
}

func (oidc *IdentitySelfServiceMethodsOidcType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"enabled": types.BoolType,
			"config":  (&IdentitySelfServiceMethodsOidcConfigType{}).TerraformType(),
		},
	}
}

func (methods *IdentitySelfServiceMethodsType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"link":          (&IdentitySelfServiceMethodsLinkType{}).TerraformType(),
			"code":          (&IdentitySelfServiceMethodsCodeType{}).TerraformType(),
			"password":      (&IdentitySelfServiceMethodsPasswordType{}).TerraformType(),
			"totp":          (&IdentitySelfServiceMethodsTotpType{}).TerraformType(),
			"lookup_secret": (&IdentitySelfServiceMethodsLookupSecretType{}).TerraformType(),
			"profile":       (&IdentitySelfServiceMethodsProfileType{}).TerraformType(),
			"webauthn":      (&IdentitySelfServiceMethodsWebAuthNType{}).TerraformType(),
			"oidc":          (&IdentitySelfServiceMethodsOidcType{}).TerraformType(),
		},
	}
}

func (after *IdentitySelfServiceFlowsLogoutAfterType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_browser_return_url": types.StringType,
		},
	}
}

func (logout *IdentitySelfServiceFlowsLogoutType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"after": (&IdentitySelfServiceFlowsLogoutAfterType{}).TerraformType(),
		},
	}
}

func (flowsError *IdentitySelfServiceFlowsErrorType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ui_url": types.StringType,
		},
	}
}

func (registration *IdentitySelfServiceFlowsRegistrationType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"login_hints": types.BoolType,
			"ui_url":      types.StringType,
			"lifespan":    types.StringType,
			"enabled":     types.BoolType,
		},
	}
}

func (login *IdentitySelfServiceFlowsLoginType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ui_url":   types.StringType,
			"lifespan": types.StringType,
		},
	}
}

func (verification *IdentitySelfServiceFlowsVerificationType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ui_url":                    types.StringType,
			"lifespan":                  types.StringType,
			"use":                       types.StringType,
			"notify_unknown_recipients": types.BoolType,
			"enabled":                   types.BoolType,
		},
	}
}

func (recovery *IdentitySelfServiceFlowsRecoveryType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ui_url":                    types.StringType,
			"lifespan":                  types.StringType,
			"use":                       types.StringType,
			"notify_unknown_recipients": types.BoolType,
			"enabled":                   types.BoolType,
		},
	}
}

func (settings *IdentitySelfServiceFlowsSettingsType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"ui_url":                     types.StringType,
			"lifespan":                   types.StringType,
			"privileged_session_max_age": types.StringType,
			"required_aal":               types.StringType,
		},
	}
}

func (flows *IdentitySelfServiceFlowsType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"logout":       (&IdentitySelfServiceFlowsLogoutType{}).TerraformType(),
			"error":        (&IdentitySelfServiceFlowsErrorType{}).TerraformType(),
			"registration": (&IdentitySelfServiceFlowsRegistrationType{}).TerraformType(),
			"login":        (&IdentitySelfServiceFlowsLoginType{}).TerraformType(),
			"verification": (&IdentitySelfServiceFlowsVerificationType{}).TerraformType(),
			"recovery":     (&IdentitySelfServiceFlowsRecoveryType{}).TerraformType(),
			"settings":     (&IdentitySelfServiceFlowsSettingsType{}).TerraformType(),
		},
	}
}

func (selfService *IdentitySelfServiceType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"default_browser_return_url": types.StringType,
			"allowed_return_urls":        types.ListType{ElemType: types.StringType},
			"methods":                    (&IdentitySelfServiceMethodsType{}).TerraformType(),
			"flows":                      (&IdentitySelfServiceFlowsType{}).TerraformType(),
		},
	}
}

func (identity *IdentityType) TerraformType() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"config": (&IdentityConfigType{}).TerraformType(),
		},
	}
}
