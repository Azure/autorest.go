// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcompute_test

import (
	"armcompute"
	"armcompute/fake"
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/stretchr/testify/require"
)

func TestFakeClientFactory(t *testing.T) {
	const fakeGalleryID = "/fake/gallery/ID"
	const fakeImageID = "/fake/image/ID"
	server := fake.ServerFactory{
		GalleriesServer: fake.GalleriesServer{
			Get: func(ctx context.Context, resourceGroupName, galleryName string, options *armcompute.GalleriesClientGetOptions) (resp azfake.Responder[armcompute.GalleriesClientGetResponse], errResp azfake.ErrorResponder) {
				resp.SetResponse(http.StatusOK, armcompute.GalleriesClientGetResponse{
					Gallery: armcompute.Gallery{
						ID: to.Ptr(fakeGalleryID),
					},
				}, nil)
				return
			},
		},
		ImagesServer: fake.ImagesServer{
			Get: func(ctx context.Context, resourceGroupName, imageName string, options *armcompute.ImagesClientGetOptions) (resp azfake.Responder[armcompute.ImagesClientGetResponse], errResp azfake.ErrorResponder) {
				resp.SetResponse(http.StatusOK, armcompute.ImagesClientGetResponse{
					Image: armcompute.Image{
						ID: to.Ptr(fakeImageID),
					},
				}, nil)
				return
			},
		},
	}
	factory, err := armcompute.NewClientFactory("fake-subscription-id", &azfake.TokenCredential{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewServerFactoryTransport(&server),
		},
	})
	require.NoError(t, err)

	setsClient := factory.NewAvailabilitySetsClient()
	resp, err := setsClient.Get(context.Background(), "fake-rg", "fake-set-name", nil)
	require.Error(t, err)
	require.Zero(t, resp)

	galleriesClient := factory.NewGalleriesClient()
	galleryResp, err := galleriesClient.Get(context.Background(), "fake-rg", "fake-gallery", nil)
	require.NoError(t, err)
	require.NotNil(t, galleryResp.ID)
	require.EqualValues(t, fakeGalleryID, *galleryResp.ID)

	imagesClient := factory.NewImagesClient()
	imagesResp, err := imagesClient.Get(context.Background(), "fake-rg", "fake-image", nil)
	require.NoError(t, err)
	require.NotNil(t, imagesResp.ID)
	require.EqualValues(t, fakeImageID, *imagesResp.ID)
}

func TestFakeDisksClientBeginDeleteConcurrent(t *testing.T) {
	server := fake.DisksServer{
		BeginDelete: func(ctx context.Context, resourceGroupName, diskName string, options *armcompute.DisksClientBeginDeleteOptions) (resp azfake.PollerResponder[armcompute.DisksClientDeleteResponse], errResp azfake.ErrorResponder) {
			resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.AddNonTerminalResponse(http.StatusOK, nil)
			resp.SetTerminalResponse(http.StatusOK, armcompute.DisksClientDeleteResponse{}, nil)
			return
		},
	}
	client, err := armcompute.NewDisksClient("fake-subscription-id", &azfake.TokenCredential{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewDisksServerTransport(&server),
		},
	})
	require.NoError(t, err)
	poller1, err := client.BeginDelete(context.Background(), "fake-rg", "disk-1", nil)
	require.NoError(t, err)
	poller2, err := client.BeginDelete(context.Background(), "fake-rg", "disk-2", nil)
	require.NoError(t, err)

	poller1Done := make(chan error)
	go func() {
		_, err := poller1.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
			Frequency: time.Second,
		})
		poller1Done <- err
	}()

	poller2Done := make(chan error)
	go func() {
		_, err := poller2.PollUntilDone(context.Background(), &runtime.PollUntilDoneOptions{
			Frequency: time.Second,
		})
		poller2Done <- err
	}()

	require.NoError(t, <-poller1Done)
	require.NoError(t, <-poller2Done)
}

func TestFakeDisksClientNewListByResourceGroupPagerConcurrent(t *testing.T) {
	server := fake.DisksServer{
		NewListByResourceGroupPager: func(resourceGroupName string, options *armcompute.DisksClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.DisksClientListByResourceGroupResponse]) {
			for pageCount := 0; pageCount < 10; pageCount++ {
				resp.AddPage(http.StatusOK, armcompute.DisksClientListByResourceGroupResponse{}, nil)
			}
			return
		},
	}
	client, err := armcompute.NewDisksClient("fake-subscription-id", &azfake.TokenCredential{}, &arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Transport: fake.NewDisksServerTransport(&server),
		},
	})
	require.NoError(t, err)
	pager1 := client.NewListByResourceGroupPager("fake-rg-1", nil)
	pager2 := client.NewListByResourceGroupPager("fake-rg-2", nil)

	pager1Done := make(chan error)
	go func() {
		var err error
		for pager1.More() {
			_, err = pager1.NextPage(context.Background())
			if err != nil {
				break
			}
		}
		pager1Done <- err
	}()

	pager2Done := make(chan error)
	go func() {
		var err error
		for pager2.More() {
			_, err = pager2.NextPage(context.Background())
			if err != nil {
				break
			}
		}
		pager2Done <- err
	}()

	require.NoError(t, <-pager1Done)
	require.NoError(t, <-pager2Done)
}
