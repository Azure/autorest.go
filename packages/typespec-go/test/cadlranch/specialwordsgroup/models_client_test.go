//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package specialwordsgroup_test

import (
	"context"
	"specialwordsgroup"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestModelsClient_WithAnd(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAnd(context.Background(), specialwordsgroup.And{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAs(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAs(context.Background(), specialwordsgroup.As{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAssert(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAssert(context.Background(), specialwordsgroup.Assert{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAsync(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAsync(context.Background(), specialwordsgroup.Async{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAwait(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAwait(context.Background(), specialwordsgroup.Await{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithBreak(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithBreak(context.Background(), specialwordsgroup.Break{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithClass(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithClass(context.Background(), specialwordsgroup.Class{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithConstructor(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithConstructor(context.Background(), specialwordsgroup.Constructor{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithContinue(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithContinue(context.Background(), specialwordsgroup.Continue{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithDef(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithDef(context.Background(), specialwordsgroup.Def{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithDel(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithDel(context.Background(), specialwordsgroup.Del{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithElif(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithElif(context.Background(), specialwordsgroup.Elif{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithElse(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithElse(context.Background(), specialwordsgroup.Else{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithExcept(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithExcept(context.Background(), specialwordsgroup.Except{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithExec(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithExec(context.Background(), specialwordsgroup.Exec{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithFinally(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithFinally(context.Background(), specialwordsgroup.Finally{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithFor(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithFor(context.Background(), specialwordsgroup.For{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithFrom(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithFrom(context.Background(), specialwordsgroup.From{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithGlobal(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithGlobal(context.Background(), specialwordsgroup.Global{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithIf(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithIf(context.Background(), specialwordsgroup.If{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithImport(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithImport(context.Background(), specialwordsgroup.Import{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithIn(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithIn(context.Background(), specialwordsgroup.In{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithIs(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithIs(context.Background(), specialwordsgroup.Is{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithLambda(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithLambda(context.Background(), specialwordsgroup.Lambda{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithNot(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithNot(context.Background(), specialwordsgroup.Not{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithOr(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithOr(context.Background(), specialwordsgroup.Or{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithPass(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithPass(context.Background(), specialwordsgroup.Pass{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithRaise(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithRaise(context.Background(), specialwordsgroup.Raise{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithReturn(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithReturn(context.Background(), specialwordsgroup.Return{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithTry(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithTry(context.Background(), specialwordsgroup.Try{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithWhile(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithWhile(context.Background(), specialwordsgroup.While{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithWith(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithWith(context.Background(), specialwordsgroup.With{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithYield(t *testing.T) {
	client, err := specialwordsgroup.NewModelsClient(nil)
	require.NoError(t, err)
	resp, err := client.WithYield(context.Background(), specialwordsgroup.Yield{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
