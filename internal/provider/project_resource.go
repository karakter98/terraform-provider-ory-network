// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	ory "github.com/ory/client-go"
	"terraform-provider-ory-network/internal/api"
	projectmodel "terraform-provider-ory-network/internal/models/project"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ProjectResourceProps{}
var _ resource.ResourceWithConfigure = &ProjectResourceProps{}
var _ resource.ResourceWithImportState = &ProjectResourceProps{}

func ProjectResource() resource.Resource {
	return &ProjectResourceProps{}
}

// ProjectResourceProps defines the resource implementation.
type ProjectResourceProps struct {
	client *ory.APIClient
}

func (r *ProjectResourceProps) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (r *ProjectResourceProps) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	corsAttributeSchema := schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Optional: true,
				Computed: true,
			},
			"origins": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Ory Network Project",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Project identifier (UUID)",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cors_admin":  corsAttributeSchema,
			"cors_public": corsAttributeSchema,
			"workspace_id": schema.StringAttribute{
				Optional: true,
			},
			"services": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"permission": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"namespaces": schema.ListNestedAttribute{
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"id": schema.Int64Attribute{
													Optional: true,
													Computed: true,
												},
												"name": schema.StringAttribute{
													Optional: true,
													Computed: true,
												},
											},
										},
										Optional: true,
										Computed: true,
									},
								},
								Required: true,
							},
						},
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
					},
					"identity": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"selfservice": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"default_browser_return_url": schema.StringAttribute{
												Required: true,
											},
											"allowed_return_urls": schema.ListAttribute{
												ElementType: types.StringType,
												Optional:    true,
												Computed:    true,
											},
											"methods": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"link": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"lifespan": schema.StringAttribute{
																		Optional: true,
																		Computed: true,
																	},
																	"base_url": schema.StringAttribute{
																		Optional: true,
																		Computed: true,
																	},
																},
																Optional: true,
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"code": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"passwordless_login_fallback_enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"passwordless_enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"lifespan": schema.StringAttribute{
																		Optional: true,
																		Computed: true,
																	},
																},
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"password": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"haveibeenpwned_enabled": schema.BoolAttribute{
																		Optional: true,
																		Computed: true,
																	},
																	"max_breaches": schema.Int64Attribute{
																		Optional: true,
																		Computed: true,
																	},
																	"ignore_network_errors": schema.BoolAttribute{
																		Optional: true,
																		Computed: true,
																	},
																	"min_password_length": schema.Int64Attribute{
																		Optional: true,
																		Computed: true,
																	},
																	"identifier_similarity_check_enabled": schema.BoolAttribute{
																		Optional: true,
																		Computed: true,
																	},
																	"haveibeenpwned_host": schema.StringAttribute{
																		Optional: true,
																		Computed: true,
																	},
																},
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"totp": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"issuer": schema.StringAttribute{
																		Optional: true,
																		Computed: true,
																	},
																},
																Optional: true,
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"lookup_secret": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"profile": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"webauthn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"passwordless": schema.BoolAttribute{
																		Optional: true,
																		Computed: true,
																	},
																	"rp": schema.SingleNestedAttribute{
																		Attributes: map[string]schema.Attribute{
																			"id": schema.StringAttribute{
																				Required: true,
																			},
																			"display_name": schema.StringAttribute{
																				Required: true,
																			},
																			"icon": schema.StringAttribute{
																				Optional: true,
																				Computed: true,
																			},
																		},
																		Optional: true,
																		Computed: true,
																	},
																},
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"oidc": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"base_redirect_uri": schema.StringAttribute{
																		Optional: true,
																		Computed: true,
																	},
																	"providers": schema.ListNestedAttribute{
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"id": schema.StringAttribute{
																					Required: true,
																				},
																				"provider": schema.StringAttribute{
																					Required: true,
																				},
																				"client_id": schema.StringAttribute{
																					Required: true,
																				},
																				"mapper_url": schema.StringAttribute{
																					Required: true,
																				},
																				"client_secret": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"issuer_url": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"auth_url": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"token_url": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"scope": schema.ListAttribute{
																					ElementType: types.StringType,
																					Optional:    true,
																					Computed:    true,
																				},
																				"microsoft_tenant": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"subject_source": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"apple_team_id": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"apple_private_key_id": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"apple_private_key": schema.StringAttribute{
																					Optional:  true,
																					Computed:  true,
																					Sensitive: true,
																				},
																				"requested_claims": schema.SingleNestedAttribute{
																					Attributes: map[string]schema.Attribute{
																						"id_token": schema.ListAttribute{
																							ElementType: types.StringType,
																							Optional:    true,
																							Computed:    true,
																						},
																					},
																					Optional: true,
																					Computed: true,
																				},
																				"organization_id": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"label": schema.StringAttribute{
																					Optional: true,
																					Computed: true,
																				},
																				"additional_id_token_audiences": schema.ListAttribute{
																					ElementType: types.StringType,
																					Optional:    true,
																					Computed:    true,
																				},
																			},
																		},
																		Optional: true,
																		Computed: true,
																	},
																},
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
												},
												Optional: true,
												Computed: true,
											},
											"flows": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"logout": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"after": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"default_browser_return_url": schema.StringAttribute{
																		Optional: true,
																		Computed: true,
																	},
																},
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"error": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"registration": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"login_hints": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"ui_url": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"login": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"verification": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"use": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"notify_unknown_recipients": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"recovery": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"use": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"notify_unknown_recipients": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
													"settings": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"privileged_session_max_age": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
															"required_aal": schema.StringAttribute{
																Optional: true,
																Computed: true,
															},
														},
														Optional: true,
														Computed: true,
													},
												},
												Optional: true,
												Computed: true,
											},
										},
										Required: true,
									},
									"identity": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"default_schema_id": schema.StringAttribute{
												Required: true,
											},
											"schemas": schema.ListNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"id": schema.StringAttribute{
															Required: true,
														},
														"url": schema.StringAttribute{
															Required: true,
														},
													},
												},
												Required: true,
											},
										},
										Required: true,
									},
								},
								Required: true,
							},
						},
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
					},
				},
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *ProjectResourceProps) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ory.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *ory.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ProjectResourceProps) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data projectmodel.ProjectType

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, err := api.CreateProject(r.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create project, got error: %s", err))
		return
	}

	createData := *projectmodel.NewProjectFromApiRepresentation(project, &ctx)
	// Save intermediate data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &createData)...)

	data.Id = createData.Id
	if data.Services.IsNull() || data.Services.IsUnknown() {
		data.Services = createData.Services
	}
	if data.CorsPublic.IsNull() || data.CorsPublic.IsUnknown() {
		data.CorsPublic = createData.CorsPublic
	}
	if data.CorsAdmin.IsNull() || data.CorsAdmin.IsUnknown() {
		data.CorsAdmin = createData.CorsAdmin
	}

	configuredServices := projectmodel.NewProjectServicesFromTerraformRepresentation(&data.Services, &ctx)
	defaultCreatedServices := projectmodel.NewProjectServicesFromTerraformRepresentation(&createData.Services, &ctx)
	configuredServices.MergeWith(defaultCreatedServices)
	data.Services = configuredServices.ToTerraformRepresentation(&ctx)

	project, err = api.UpdateProject(r.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", fmt.Sprintf("Unable to update project settings, got error: %s", err))
		return
	}
	updateData := *projectmodel.NewProjectFromApiRepresentation(project, &ctx)

	// Save final data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &updateData)...)
}

func (r *ProjectResourceProps) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data projectmodel.ProjectType

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, err := api.ReadProject(r.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read project, got error: %s", err))
		return
	}

	data = *projectmodel.NewProjectFromApiRepresentation(project, &ctx)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResourceProps) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data projectmodel.ProjectType

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, err := api.UpdateProject(r.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Update Error", fmt.Sprintf("Unable to update project settings, got error: %s", err))
		return
	}

	data = *projectmodel.NewProjectFromApiRepresentation(project, &ctx)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ProjectResourceProps) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data projectmodel.ProjectType

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	// httpResp, err := r.client.Do(httpReq)
	err := api.DeleteProject(r.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete project, got error: %s", err))
		return
	}
}

func (r *ProjectResourceProps) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
