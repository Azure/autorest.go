// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paramgroupinggroup_test

import (
	"context"
	"generatortests/paramgroupinggroup"
	"generatortests/paramgroupinggroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakePostMultiParamGroups(t *testing.T) {
	group1 := paramgroupinggroup.FirstParameterGroup{
		HeaderOne: to.Ptr("header1"),
		QueryOne:  to.Ptr[int32](123),
	}
	group2 := paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup{
		HeaderTwo: to.Ptr("header2"),
		QueryTwo:  to.Ptr[int32](456),
	}
	group3 := paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsOptions{}
	server := fake.ParameterGroupingServer{
		PostMultiParamGroups: func(ctx context.Context, firstParameterGroup *paramgroupinggroup.FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup, options *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsResponse], errResp azfake.ErrorResponder) {
			require.NotNil(t, firstParameterGroup)
			require.NotNil(t, parameterGroupingClientPostMultiParamGroupsSecondParamGroup)
			require.Nil(t, options)
			require.EqualValues(t, group1, *firstParameterGroup)
			require.EqualValues(t, group2, *parameterGroupingClientPostMultiParamGroupsSecondParamGroup)
			resp.SetResponse(http.StatusOK, paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsResponse{}, nil)
			return
		},
	}
	client, err := paramgroupinggroup.NewParameterGroupingClient(&azcore.ClientOptions{
		Transport: fake.NewParameterGroupingServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PostMultiParamGroups(context.Background(), &group1, &group2, &group3)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakePostMultiParamGroupsEmpty(t *testing.T) {
	group1 := paramgroupinggroup.FirstParameterGroup{}
	group2 := paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup{}
	group3 := paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsOptions{}
	server := fake.ParameterGroupingServer{
		PostMultiParamGroups: func(ctx context.Context, firstParameterGroup *paramgroupinggroup.FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup, options *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, firstParameterGroup)
			require.Nil(t, parameterGroupingClientPostMultiParamGroupsSecondParamGroup)
			require.Nil(t, options)
			resp.SetResponse(http.StatusOK, paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsResponse{}, nil)
			return
		},
	}
	client, err := paramgroupinggroup.NewParameterGroupingClient(&azcore.ClientOptions{
		Transport: fake.NewParameterGroupingServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PostMultiParamGroups(context.Background(), &group1, &group2, &group3)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakePostMultiParamGroupsNil(t *testing.T) {
	server := fake.ParameterGroupingServer{
		PostMultiParamGroups: func(ctx context.Context, firstParameterGroup *paramgroupinggroup.FirstParameterGroup, parameterGroupingClientPostMultiParamGroupsSecondParamGroup *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsSecondParamGroup, options *paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsOptions) (resp azfake.Responder[paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, firstParameterGroup)
			require.Nil(t, parameterGroupingClientPostMultiParamGroupsSecondParamGroup)
			require.Nil(t, options)
			resp.SetResponse(http.StatusOK, paramgroupinggroup.ParameterGroupingClientPostMultiParamGroupsResponse{}, nil)
			return
		},
	}
	client, err := paramgroupinggroup.NewParameterGroupingClient(&azcore.ClientOptions{
		Transport: fake.NewParameterGroupingServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PostMultiParamGroups(context.Background(), nil, nil, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
