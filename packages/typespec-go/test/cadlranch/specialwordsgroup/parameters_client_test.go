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
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithAnd(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAs(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithAs(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAssert(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithAssert(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAsync(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithAsync(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAwait(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithAwait(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithBreak(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithBreak(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithCancellationToken(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithCancellationToken(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithClass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithClass(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithConstructor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithConstructor(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithContinue(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithContinue(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithDef(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithDef(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithDel(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithDel(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithElif(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithElif(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithElse(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithElse(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithExcept(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithExcept(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithExec(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithExec(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFinally(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithFinally(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithFor(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFrom(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithFrom(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithGlobal(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithGlobal(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIf(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithIf(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithImport(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithImport(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIn(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithIn(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIs(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithIs(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithLambda(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithLambda(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithNot(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithNot(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithOr(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithOr(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithPass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithPass(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithRaise(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithRaise(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithReturn(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithReturn(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithTry(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithTry(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithWhile(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithWhile(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithWith(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithWith(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithYield(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewParametersClient().WithYield(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
