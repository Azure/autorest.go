//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package specialwordsgroup_test

import (
	"context"
	"specialwordsgroup"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParametersClient_WithAnd(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithAnd(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAs(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithAs(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAssert(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithAssert(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAsync(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithAsync(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAwait(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithAwait(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithBreak(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithBreak(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithCancellationToken(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithCancellationToken(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithClass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithClass(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithConstructor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithConstructor(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithContinue(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithContinue(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithDef(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithDef(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithDel(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithDel(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithElif(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithElif(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithElse(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithElse(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithExcept(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithExcept(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithExec(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithExec(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFinally(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithFinally(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithFor(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFrom(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithFrom(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithGlobal(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithGlobal(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIf(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithIf(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithImport(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithImport(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIn(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithIn(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIs(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithIs(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithLambda(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithLambda(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithNot(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithNot(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithOr(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithOr(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithPass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithPass(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithRaise(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithRaise(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithReturn(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithReturn(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithTry(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithTry(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithWhile(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithWhile(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithWith(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithWith(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithYield(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsParametersClient().WithYield(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
