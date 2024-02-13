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
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().And(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_As(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().As(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Assert(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Assert(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Async(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Async(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Await(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Await(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Break(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Break(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Class(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Class(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Constructor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Constructor(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Continue(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Continue(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Def(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Def(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Del(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Del(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Elif(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Elif(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Else(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Else(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Except(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Except(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Exec(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Exec(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Finally(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Finally(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_For(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().For(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_From(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().From(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Global(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Global(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_If(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().If(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Import(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Import(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_In(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().In(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Is(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Is(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Lambda(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Lambda(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Not(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Not(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Or(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Or(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Pass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Pass(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Raise(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Raise(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Return(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Return(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Try(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Try(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_While(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().While(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_With(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().With(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Yield(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClient(nil)
	require.NoError(t, err)
	resp, err := client.NewOperationsClient().Yield(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
