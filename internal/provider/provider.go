// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure MathProvider satisfies various provider interfaces.
var _ provider.Provider = &MathProvider{}
var _ provider.ProviderWithFunctions = &MathProvider{}

// MathProvider defines the provider implementation.
type MathProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// MathProviderModel describes the provider data model.
type MathProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *MathProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "math"
	resp.Version = p.version
}

func (p *MathProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Math library provider that offers functions and resources for math operations.",
	}
}

func (p *MathProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *MathProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewFibResource,
	}
}

func (p *MathProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}

func (p *MathProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewFibFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &MathProvider{
			version: version,
		}
	}
}
