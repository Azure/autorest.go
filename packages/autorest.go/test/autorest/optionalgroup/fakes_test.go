// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package optionalgroup_test

import (
	"context"
	"generatortests/optionalgroup"
	"generatortests/optionalgroup/fake"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/stretchr/testify/require"
)

func TestFakePostOptionalStringProperty(t *testing.T) {
	value := "optional-value"
	server := fake.ExplicitServer{
		PostOptionalStringProperty: func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalStringPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalStringPropertyResponse], errResp azfake.ErrorResponder) {
			require.NotNil(t, options)
			require.EqualValues(t, value, *options.BodyParameter.Value)
			resp.SetResponse(http.StatusOK, optionalgroup.ExplicitClientPostOptionalStringPropertyResponse{}, nil)
			return
		},
	}
	client, err := optionalgroup.NewExplicitClient(&azcore.ClientOptions{
		Transport: fake.NewExplicitServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PostOptionalStringProperty(context.Background(), &optionalgroup.ExplicitClientPostOptionalStringPropertyOptions{
		BodyParameter: &optionalgroup.StringOptionalWrapper{
			Value: &value,
		},
	})
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakePostOptionalStringPropertyNil(t *testing.T) {
	server := fake.ExplicitServer{
		PostOptionalStringProperty: func(ctx context.Context, options *optionalgroup.ExplicitClientPostOptionalStringPropertyOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostOptionalStringPropertyResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, options)
			resp.SetResponse(http.StatusOK, optionalgroup.ExplicitClientPostOptionalStringPropertyResponse{}, nil)
			return
		},
	}
	client, err := optionalgroup.NewExplicitClient(&azcore.ClientOptions{
		Transport: fake.NewExplicitServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PostOptionalStringProperty(context.Background(), nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}

func TestFakePostRequiredStringParameter(t *testing.T) {
	server := fake.ExplicitServer{
		PostRequiredStringParameter: func(ctx context.Context, bodyParameter string, options *optionalgroup.ExplicitClientPostRequiredStringParameterOptions) (resp azfake.Responder[optionalgroup.ExplicitClientPostRequiredStringParameterResponse], errResp azfake.ErrorResponder) {
			require.EqualValues(t, "the body", bodyParameter)
			resp.SetResponse(http.StatusOK, optionalgroup.ExplicitClientPostRequiredStringParameterResponse{}, nil)
			return
		},
	}
	client, err := optionalgroup.NewExplicitClient(&azcore.ClientOptions{
		Transport: fake.NewExplicitServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PostRequiredStringParameter(context.Background(), "the body", nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
