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
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithAnd(context.Background(), specialwordsgroup.And{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAs(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithAs(context.Background(), specialwordsgroup.As{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAssert(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithAssert(context.Background(), specialwordsgroup.Assert{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAsync(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithAsync(context.Background(), specialwordsgroup.Async{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithAwait(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithAwait(context.Background(), specialwordsgroup.Await{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithBreak(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithBreak(context.Background(), specialwordsgroup.Break{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithClass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithClass(context.Background(), specialwordsgroup.Class{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithConstructor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithConstructor(context.Background(), specialwordsgroup.Constructor{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithContinue(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithContinue(context.Background(), specialwordsgroup.Continue{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithDef(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithDef(context.Background(), specialwordsgroup.Def{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithDel(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithDel(context.Background(), specialwordsgroup.Del{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithElif(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithElif(context.Background(), specialwordsgroup.Elif{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithElse(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithElse(context.Background(), specialwordsgroup.Else{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithExcept(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithExcept(context.Background(), specialwordsgroup.Except{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithExec(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithExec(context.Background(), specialwordsgroup.Exec{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithFinally(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithFinally(context.Background(), specialwordsgroup.Finally{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithFor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithFor(context.Background(), specialwordsgroup.For{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithFrom(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithFrom(context.Background(), specialwordsgroup.From{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithGlobal(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithGlobal(context.Background(), specialwordsgroup.Global{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithIf(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithIf(context.Background(), specialwordsgroup.If{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithImport(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithImport(context.Background(), specialwordsgroup.Import{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithIn(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithIn(context.Background(), specialwordsgroup.In{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithIs(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithIs(context.Background(), specialwordsgroup.Is{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithLambda(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithLambda(context.Background(), specialwordsgroup.Lambda{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithNot(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithNot(context.Background(), specialwordsgroup.Not{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithOr(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithOr(context.Background(), specialwordsgroup.Or{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithPass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithPass(context.Background(), specialwordsgroup.Pass{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithRaise(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithRaise(context.Background(), specialwordsgroup.Raise{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithReturn(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithReturn(context.Background(), specialwordsgroup.Return{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithTry(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithTry(context.Background(), specialwordsgroup.Try{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithWhile(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithWhile(context.Background(), specialwordsgroup.While{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithWith(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithWith(context.Background(), specialwordsgroup.With{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestModelsClient_WithYield(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsModelsClient().WithYield(context.Background(), specialwordsgroup.Yield{
		Name: to.Ptr("ok"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
