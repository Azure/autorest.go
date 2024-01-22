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
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAnd(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAs(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAs(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAssert(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAssert(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAsync(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAsync(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithAwait(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithAwait(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithBreak(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithBreak(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithCancellationToken(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithCancellationToken(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithClass(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithClass(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithConstructor(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithConstructor(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithContinue(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithContinue(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithDef(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithDef(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithDel(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithDel(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithElif(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithElif(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithElse(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithElse(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithExcept(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithExcept(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithExec(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithExec(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFinally(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithFinally(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFor(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithFor(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithFrom(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithFrom(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithGlobal(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithGlobal(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIf(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithIf(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithImport(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithImport(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIn(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithIn(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithIs(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithIs(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithLambda(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithLambda(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithNot(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithNot(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithOr(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithOr(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithPass(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithPass(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithRaise(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithRaise(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithReturn(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithReturn(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithTry(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithTry(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithWhile(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithWhile(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithWith(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithWith(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestParametersClient_WithYield(t *testing.T) {
	client, err := specialwordsgroup.NewParametersClient(nil)
	require.NoError(t, err)
	resp, err := client.WithYield(context.Background(), "ok", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
