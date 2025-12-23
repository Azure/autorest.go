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
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().And(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_As(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().As(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Assert(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Assert(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Async(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Async(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Await(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Await(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Break(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Break(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Class(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Class(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Constructor(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Constructor(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Continue(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Continue(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Def(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Def(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Del(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Del(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Elif(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Elif(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Else(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Else(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Except(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Except(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Exec(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Exec(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Finally(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Finally(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_For(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().For(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_From(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().From(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Global(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Global(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_If(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().If(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Import(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Import(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_In(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().In(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Is(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Is(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Lambda(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Lambda(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Not(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Not(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Or(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Or(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Pass(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Pass(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Raise(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Raise(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Return(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Return(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Try(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Try(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_While(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().While(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_With(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().With(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestOperationsClient_Yield(t *testing.T) {
	client, err := specialwordsgroup.NewSpecialWordsClientWithNoCredential("http://localhost:3000", nil)
	require.NoError(t, err)
	resp, err := client.NewSpecialWordsOperationsClient().Yield(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
