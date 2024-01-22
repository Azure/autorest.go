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

func TestOperationsClient_And(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.And(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_As(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.As(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Assert(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Assert(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Async(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Async(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Await(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Await(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Break(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Break(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Class(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Class(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Constructor(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Constructor(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Continue(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Continue(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Def(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Def(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Del(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Del(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Elif(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Elif(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Else(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Else(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Except(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Except(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Exec(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Exec(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Finally(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Finally(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_For(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.For(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_From(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.From(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Global(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Global(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_If(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.If(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Import(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Import(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_In(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.In(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Is(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Is(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Lambda(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Lambda(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Not(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Not(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Or(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Or(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Pass(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Pass(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Raise(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Raise(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Return(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Return(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Try(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Try(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_While(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.While(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_With(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.With(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Yield(t *testing.T) {
	client, err := specialwordsgroup.NewOperationsClient(nil)
	require.NoError(t, err)
	resp, err := client.Yield(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
