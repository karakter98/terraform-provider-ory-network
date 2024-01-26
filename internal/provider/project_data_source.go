// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	ory "github.com/ory/client-go"
	"terraform-provider-ory-network/internal/api"
	projectmodel "terraform-provider-ory-network/internal/models/project"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = &ProjectDataSourceProps{}
	_ datasource.DataSourceWithConfigure = &ProjectDataSourceProps{}
)

func ProjectDataSource() datasource.DataSource {
	return &ProjectDataSourceProps{}
}

// ProjectDataSourceProps defines the data source implementation.
type ProjectDataSourceProps struct {
	client *ory.APIClient
}

func (d *ProjectDataSourceProps) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

func (d *ProjectDataSourceProps) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	corsAttributeSchema := schema.SingleNestedAttribute{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				Computed: true,
			},
			"origins": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
		},
		Computed: true,
	}

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Ory Network Project",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "Project identifier (UUID)",
				Required:    true,
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
			},
			"cors_admin":  corsAttributeSchema,
			"cors_public": corsAttributeSchema,
			"workspace_id": schema.StringAttribute{
				Computed: true,
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
													Computed: true,
												},
												"name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
										Computed: true,
									},
								},
								Computed: true,
							},
						},
						Computed: true,
					},
					"identity": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"config": schema.SingleNestedAttribute{
								Attributes: map[string]schema.Attribute{
									"selfservice": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"default_browser_return_url": schema.StringAttribute{
												Computed: true,
											},
											"allowed_return_urls": schema.ListAttribute{
												ElementType: types.StringType,
												Computed:    true,
											},
											"methods": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"link": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"lifespan": schema.StringAttribute{
																		Computed: true,
																	},
																	"base_url": schema.StringAttribute{
																		Computed: true,
																	},
																},
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"code": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"passwordless_login_fallback_enabled": schema.BoolAttribute{
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
															"passwordless_enabled": schema.BoolAttribute{
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"lifespan": schema.StringAttribute{
																		Computed: true,
																	},
																},
																Computed: true,
															},
														},
														Computed: true,
													},
													"password": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"haveibeenpwned_enabled": schema.BoolAttribute{
																		Computed: true,
																	},
																	"max_breaches": schema.Int64Attribute{
																		Computed: true,
																	},
																	"ignore_network_errors": schema.BoolAttribute{
																		Computed: true,
																	},
																	"min_password_length": schema.Int64Attribute{
																		Computed: true,
																	},
																	"identifier_similarity_check_enabled": schema.BoolAttribute{
																		Computed: true,
																	},
																	"haveibeenpwned_host": schema.StringAttribute{
																		Computed: true,
																	},
																},
																Computed: true,
															},
														},
														Computed: true,
													},
													"totp": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"issuer": schema.StringAttribute{
																		Computed: true,
																	},
																},
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"lookup_secret": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"profile": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"webauthn": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"passwordless": schema.BoolAttribute{
																		Computed: true,
																	},
																	"rp": schema.SingleNestedAttribute{
																		Attributes: map[string]schema.Attribute{
																			"id": schema.StringAttribute{
																				Computed: true,
																			},
																			"display_name": schema.StringAttribute{
																				Computed: true,
																			},
																			"icon": schema.StringAttribute{
																				Computed: true,
																			},
																		},
																		Computed: true,
																	},
																},
																Computed: true,
															},
														},
														Computed: true,
													},
													"oidc": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
															"config": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"base_redirect_uri": schema.StringAttribute{
																		Computed: true,
																	},
																	"providers": schema.ListNestedAttribute{
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"id": schema.StringAttribute{
																					Computed: true,
																				},
																				"provider": schema.StringAttribute{
																					Computed: true,
																				},
																				"client_id": schema.StringAttribute{
																					Computed: true,
																				},
																				"mapper_url": schema.StringAttribute{
																					Computed: true,
																				},
																				"client_secret": schema.StringAttribute{
																					Computed: true,
																				},
																				"issuer_url": schema.StringAttribute{
																					Computed: true,
																				},
																				"auth_url": schema.StringAttribute{
																					Computed: true,
																				},
																				"token_url": schema.StringAttribute{
																					Computed: true,
																				},
																				"scope": schema.ListAttribute{
																					ElementType: types.StringType,
																					Computed:    true,
																				},
																				"microsoft_tenant": schema.StringAttribute{
																					Computed: true,
																				},
																				"subject_source": schema.StringAttribute{
																					Computed: true,
																				},
																				"apple_team_id": schema.StringAttribute{
																					Computed: true,
																				},
																				"apple_private_key_id": schema.StringAttribute{
																					Computed: true,
																				},
																				"apple_private_key": schema.StringAttribute{
																					Computed:  true,
																					Sensitive: true,
																				},
																				"requested_claims": schema.SingleNestedAttribute{
																					Attributes: map[string]schema.Attribute{
																						"id_token": schema.ListAttribute{
																							ElementType: types.StringType,
																							Computed:    true,
																						},
																					},
																					Computed: true,
																				},
																				"organization_id": schema.StringAttribute{
																					Computed: true,
																				},
																				"label": schema.StringAttribute{
																					Computed: true,
																				},
																				"additional_id_token_audiences": schema.ListAttribute{
																					ElementType: types.StringType,
																					Computed:    true,
																				},
																			},
																		},
																		Computed: true,
																	},
																},
																Computed: true,
															},
														},
														Computed: true,
													},
												},
												Computed: true,
											},
											"flows": schema.SingleNestedAttribute{
												Attributes: map[string]schema.Attribute{
													"logout": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"after": schema.SingleNestedAttribute{
																Attributes: map[string]schema.Attribute{
																	"default_browser_return_url": schema.StringAttribute{
																		Computed: true,
																	},
																},
																Computed: true,
															},
														},
														Computed: true,
													},
													"error": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"registration": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"login_hints": schema.BoolAttribute{
																Computed: true,
															},
															"ui_url": schema.StringAttribute{
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"login": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"verification": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Computed: true,
															},
															"use": schema.StringAttribute{
																Computed: true,
															},
															"notify_unknown_recipients": schema.BoolAttribute{
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"recovery": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Computed: true,
															},
															"use": schema.StringAttribute{
																Computed: true,
															},
															"notify_unknown_recipients": schema.BoolAttribute{
																Computed: true,
															},
															"enabled": schema.BoolAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
													"settings": schema.SingleNestedAttribute{
														Attributes: map[string]schema.Attribute{
															"ui_url": schema.StringAttribute{
																Computed: true,
															},
															"lifespan": schema.StringAttribute{
																Computed: true,
															},
															"privileged_session_max_age": schema.StringAttribute{
																Computed: true,
															},
															"required_aal": schema.StringAttribute{
																Computed: true,
															},
														},
														Computed: true,
													},
												},
												Computed: true,
											},
										},
										Computed: true,
									},
									"identity": schema.SingleNestedAttribute{
										Attributes: map[string]schema.Attribute{
											"default_schema_id": schema.StringAttribute{
												Computed: true,
											},
											"schemas": schema.ListNestedAttribute{
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"id": schema.StringAttribute{
															Computed: true,
														},
														"url": schema.StringAttribute{
															Computed: true,
														},
													},
												},
												Computed: true,
											},
										},
										Computed: true,
									},
								},
								Computed: true,
							},
						},
						Computed: true,
					},
				},
				Computed: true,
			},
		},
	}
}

func (d *ProjectDataSourceProps) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*ory.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *ory.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ProjectDataSourceProps) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data projectmodel.ProjectType

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If applicable, this is a great opportunity to initialize any necessary
	// provider client data and make a call using it.
	project, err := api.ReadProject(d.client, &data, &ctx)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read project, got error: %s", err))
		return
	}

	data = *projectmodel.NewProjectFromApiRepresentation(project, &ctx)

	// Write logs using the tflog package
	// Documentation: https://terraform.io/plugin/log
	tflog.Trace(ctx, "read project")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
