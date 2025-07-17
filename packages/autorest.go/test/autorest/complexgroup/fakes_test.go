// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup_test

import (
	"context"
	"generatortests"
	"generatortests/complexgroup"
	"generatortests/complexgroup/fake"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func TestFakePolymorphismGetValid(t *testing.T) {
	server := fake.PolymorphismServer{
		GetValid: func(ctx context.Context, options *complexgroup.PolymorphismClientGetValidOptions) (resp azfake.Responder[complexgroup.PolymorphismClientGetValidResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, options)
			resp.SetResponse(http.StatusOK, complexgroup.PolymorphismClientGetValidResponse{
				FishClassification: &complexgroup.Cookiecuttershark{
					Birthday: to.Ptr(time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)),
					Length:   to.Ptr[float32](3.14),
					Age:      to.Ptr[int32](42),
					Siblings: []complexgroup.FishClassification{
						&complexgroup.Goblinshark{
							Fishtype: to.Ptr("goblin"),
							Length:   to.Ptr[float32](30),
							Species:  to.Ptr("scary"),
							Age:      to.Ptr[int32](1),
							Birthday: to.Ptr(time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)),
							Color:    to.Ptr(complexgroup.GoblinSharkColor("pinkish-gray")),
							Jawsize:  to.Ptr[int32](5),
						},
					},
					Species: to.Ptr("chocolate chip"),
				},
			}, nil)
			return
		},
	}
	client, err := complexgroup.NewPolymorphismClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewPolymorphismServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.GetValid(context.Background(), nil)
	require.NoError(t, err)
	ccs, ok := resp.FishClassification.(*complexgroup.Cookiecuttershark)
	if !ok {
		t.Fatal("fish wasn't a Cookiecuttershark")
	}
	if r := cmp.Diff(ccs, &complexgroup.Cookiecuttershark{
		Birthday: to.Ptr(time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)),
		Fishtype: to.Ptr("cookiecuttershark"),
		Length:   to.Ptr[float32](3.14),
		Age:      to.Ptr[int32](42),
		Siblings: []complexgroup.FishClassification{
			&complexgroup.Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: to.Ptr(time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)),
				Color:    to.Ptr(complexgroup.GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species: to.Ptr("chocolate chip"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestFakePolymorphismPutValid(t *testing.T) {
	server := fake.PolymorphismServer{
		PutValid: func(ctx context.Context, complexBody complexgroup.FishClassification, options *complexgroup.PolymorphismClientPutValidOptions) (resp azfake.Responder[complexgroup.PolymorphismClientPutValidResponse], errResp azfake.ErrorResponder) {
			require.Nil(t, options)
			if r := cmp.Diff(complexBody, &complexgroup.Cookiecuttershark{
				Birthday: to.Ptr(time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)),
				Fishtype: to.Ptr("cookiecuttershark"),
				Length:   to.Ptr[float32](3.14),
				Age:      to.Ptr[int32](42),
				Siblings: []complexgroup.FishClassification{
					&complexgroup.Goblinshark{
						Fishtype: to.Ptr("goblin"),
						Length:   to.Ptr[float32](30),
						Species:  to.Ptr("scary"),
						Age:      to.Ptr[int32](1),
						Birthday: to.Ptr(time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)),
						Color:    to.Ptr(complexgroup.GoblinSharkColor("pinkish-gray")),
						Jawsize:  to.Ptr[int32](5),
					},
				},
				Species: to.Ptr("chocolate chip"),
			}); r != "" {
				t.Fatal(r)
			}
			resp.SetResponse(http.StatusOK, complexgroup.PolymorphismClientPutValidResponse{}, nil)
			return
		},
	}
	client, err := complexgroup.NewPolymorphismClient(generatortests.Host, &azcore.ClientOptions{
		Transport: fake.NewPolymorphismServerTransport(&server),
	})
	require.NoError(t, err)
	resp, err := client.PutValid(context.Background(), &complexgroup.Cookiecuttershark{
		Birthday: to.Ptr(time.Date(2015, time.August, 8, 0, 0, 0, 0, time.UTC)),
		Fishtype: to.Ptr("cookiecuttershark"),
		Length:   to.Ptr[float32](3.14),
		Age:      to.Ptr[int32](42),
		Siblings: []complexgroup.FishClassification{
			&complexgroup.Goblinshark{
				Fishtype: to.Ptr("goblin"),
				Length:   to.Ptr[float32](30),
				Species:  to.Ptr("scary"),
				Age:      to.Ptr[int32](1),
				Birthday: to.Ptr(time.Date(1900, time.January, 5, 1, 0, 0, 0, time.UTC)),
				Color:    to.Ptr(complexgroup.GoblinSharkColor("pinkish-gray")),
				Jawsize:  to.Ptr[int32](5),
			},
		},
		Species: to.Ptr("chocolate chip"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, resp)
}
