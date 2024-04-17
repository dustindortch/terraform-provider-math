package provider

import (
	"context"
	"math/big"

	"terraform-provider-math/internal/fib"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

// Ensure the implementation satisfies the desired interfaces.
var _ function.Function = &FibFunction{}

type FibFunction struct{}

func NewFibFunction() function.Function {
	return &FibFunction{}
}

func (f *FibFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "fib"
}

func (f *FibFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Fibonacci sequence",
		MarkdownDescription: "Accepts a number and returns the result of the fibonacci sequence at that index.",
		Parameters: []function.Parameter{
			function.NumberParameter{
				Name: "number",
			},
		},
		Return: function.NumberReturn{},
	}
}

func (f *FibFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var number *big.Float
	var result *big.Float

	// Read Terraform argument data into the variables
	resp.Error = function.ConcatFuncErrors(resp.Error, req.Arguments.Get(ctx, &number))

	fibInt, _ := number.Uint64()
	result = big.NewFloat(0).SetUint64(fib.Fib(fibInt))

	resp.Error = function.ConcatFuncErrors(resp.Error, resp.Result.Set(ctx, result))
}
