// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func toTimePtr(layout string, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

func newXMLClient() *XMLClient {
	return NewXMLClient(&azcore.ClientOptions{
		Logging: policy.LogOptions{
			IncludeBody: true,
		},
	})
}

func TestGetACLs(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetACLs(context.Background(), nil)
	require.NoError(t, err)
	expected := []*SignedIdentifier{
		{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	}
	if r := cmp.Diff(result.SignedIdentifiers, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetBytes(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetBytes(context.Background(), nil)
	require.NoError(t, err)
	if string(result.Bytes) != "Hello world" {
		t.Fatalf("unexpected bytes %s", string(result.Bytes))
	}
}

func TestGetComplexTypeRefNoMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetComplexTypeRefNoMeta(context.Background(), nil)
	require.NoError(t, err)
	expected := RootWithRefAndNoMeta{
		RefToModel: &ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}
	if r := cmp.Diff(result.RootWithRefAndNoMeta, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetComplexTypeRefWithMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetComplexTypeRefWithMeta(context.Background(), nil)
	require.NoError(t, err)
	expected := RootWithRefAndMeta{
		RefToModel: &ComplexTypeWithMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}
	if r := cmp.Diff(result.RootWithRefAndMeta, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetEmptyChildElement(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyChildElement(context.Background(), nil)
	require.NoError(t, err)
	expected := Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	}
	if r := cmp.Diff(result.Banana, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetEmptyList(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyList(context.Background(), nil)
	require.NoError(t, err)
	expected := Slideshow{}
	if r := cmp.Diff(result.Slideshow, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetEmptyRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyRootList(context.Background(), nil)
	require.NoError(t, err)
	if result.Bananas != nil {
		t.Fatal("expected nil slice")
	}
}

func TestGetEmptyWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyWrappedLists(context.Background(), nil)
	require.NoError(t, err)
	expected := AppleBarrel{}
	if r := cmp.Diff(result.AppleBarrel, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetHeaders(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetHeaders(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.CustomHeader, to.StringPtr("custom-value")); r != "" {
		t.Fatal(r)
	}
}

func TestGetRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetRootList(context.Background(), nil)
	require.NoError(t, err)
	expected := []*Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		{
			Name:       to.StringPtr("Plantain"),
			Flavor:     to.StringPtr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	if r := cmp.Diff(result.Bananas, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetRootListSingleItem(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetRootListSingleItem(context.Background(), nil)
	require.NoError(t, err)
	expected := []*Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	if r := cmp.Diff(result.Bananas, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetServiceProperties(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetServiceProperties(context.Background(), nil)
	require.NoError(t, err)
	expected := StorageServiceProperties{
		HourMetrics: &Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeAPIs: to.BoolPtr(false),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		Logging: &Logging{
			Version: to.StringPtr("1.0"),
			Delete:  to.BoolPtr(true),
			Read:    to.BoolPtr(false),
			Write:   to.BoolPtr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		MinuteMetrics: &Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeAPIs: to.BoolPtr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	}
	if r := cmp.Diff(result.StorageServiceProperties, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetSimple(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetSimple(context.Background(), nil)
	require.NoError(t, err)
	expected := Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: []*Slide{
			{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			{
				Items: to.StringPtrArray("Why WonderWidgets are great", "", "Who buys WonderWidgets"),
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	}
	if r := cmp.Diff(result.Slideshow, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetWrappedLists(context.Background(), nil)
	require.NoError(t, err)
	expected := AppleBarrel{
		BadApples:  to.StringPtrArray("Red Delicious"),
		GoodApples: to.StringPtrArray("Fuji", "Gala"),
	}
	if r := cmp.Diff(result.AppleBarrel, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetXMsText(t *testing.T) {
	t.Skip("support NYI")
	client := newXMLClient()
	result, err := client.GetXMsText(context.Background(), nil)
	require.NoError(t, err)
	expected := ObjectWithXMsTextProperty{
		Content:  to.StringPtr("I am text"),
		Language: to.StringPtr("english"),
	}
	if r := cmp.Diff(result.ObjectWithXMsTextProperty, expected); r != "" {
		t.Fatal(r)
	}
}

func TestJSONInput(t *testing.T) {
	client := newXMLClient()
	result, err := client.JSONInput(context.Background(), JSONInput{
		ID: to.Int32Ptr(42),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestJSONOutput(t *testing.T) {
	client := newXMLClient()
	result, err := client.JSONOutput(context.Background(), nil)
	require.NoError(t, err)
	expected := JSONOutput{
		ID: to.Int32Ptr(42),
	}
	if r := cmp.Diff(result.JSONOutput, expected); r != "" {
		t.Fatal(r)
	}
}

func TestListBlobs(t *testing.T) {
	client := newXMLClient()
	result, err := client.ListBlobs(context.Background(), nil)
	require.NoError(t, err)
	blob1LM, err := time.Parse(time.RFC1123, "Wed, 09 Sep 2009 09:20:02 GMT")
	require.NoError(t, err)
	blob2LM, err := time.Parse(time.RFC1123, "Wed, 09 Sep 2009 09:20:03 GMT")
	require.NoError(t, err)
	expected := ListBlobsResponse{
		Blobs: &Blobs{
			Blob: []*Blob{
				{
					Metadata: map[string]*string{
						"color":            to.StringPtr("blue"),
						"blobnumber":       to.StringPtr("01"),
						"somemetadataname": to.StringPtr("SomeMetadataValue"),
					},
					Name: to.StringPtr("blob1.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8A29A19"),
						ContentLength:   to.Int64Ptr(100),
						ContentType:     to.StringPtr("text/html"),
						ContentEncoding: to.StringPtr(""),
						ContentLanguage: to.StringPtr("en-US"),
						ContentMD5:      to.StringPtr(""),
						CacheControl:    to.StringPtr("no-cache"),
						BlobType:        BlobTypeBlockBlob.ToPtr(),
						LeaseStatus:     LeaseStatusTypeUnlocked.ToPtr(),
					},
				},
				{
					Metadata: map[string]*string{
						"color":             to.StringPtr("green"),
						"blobnumber":        to.StringPtr("02"),
						"somemetadataname":  to.StringPtr("SomeMetadataValue"),
						"x-ms-invalid-name": to.StringPtr("nasdf$@#$$"),
					},
					Name: to.StringPtr("blob2.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Int64Ptr(5000),
						ContentType:     to.StringPtr("application/octet-stream"),
						ContentEncoding: to.StringPtr("gzip"),
						ContentLanguage: to.StringPtr(""),
						ContentMD5:      to.StringPtr(""),
						CacheControl:    to.StringPtr(""),
						BlobType:        BlobTypeBlockBlob.ToPtr(),
					},
					Snapshot: to.StringPtr("2009-09-09T09:20:03.0427659Z"),
				},
				{
					Metadata: map[string]*string{
						"color":            to.StringPtr("green"),
						"blobnumber":       to.StringPtr("02"),
						"somemetadataname": to.StringPtr("SomeMetadataValue"),
					},
					Name: to.StringPtr("blob2.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Int64Ptr(5000),
						ContentType:     to.StringPtr("application/octet-stream"),
						ContentEncoding: to.StringPtr("gzip"),
						ContentLanguage: to.StringPtr(""),
						ContentMD5:      to.StringPtr(""),
						CacheControl:    to.StringPtr(""),
						BlobType:        BlobTypeBlockBlob.ToPtr(),
					},
					Snapshot: to.StringPtr("2009-09-09T09:20:03.1587543Z"),
				},
				{
					Metadata: map[string]*string{
						"color":            to.StringPtr("green"),
						"blobnumber":       to.StringPtr("02"),
						"somemetadataname": to.StringPtr("SomeMetadataValue"),
					},
					Name: to.StringPtr("blob2.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.StringPtr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Int64Ptr(5000),
						ContentType:     to.StringPtr("application/octet-stream"),
						ContentEncoding: to.StringPtr("gzip"),
						ContentLanguage: to.StringPtr(""),
						ContentMD5:      to.StringPtr(""),
						CacheControl:    to.StringPtr(""),
						BlobType:        BlobTypeBlockBlob.ToPtr(),
						LeaseStatus:     LeaseStatusTypeUnlocked.ToPtr(),
					},
				},
				{
					Metadata: map[string]*string{
						"color":            to.StringPtr("yellow"),
						"blobnumber":       to.StringPtr("03"),
						"somemetadataname": to.StringPtr("SomeMetadataValue"),
					},
					Name: to.StringPtr("blob3.txt"),
					Properties: &BlobProperties{
						LastModified:       &blob2LM,
						Etag:               to.StringPtr("0x8CBFF45D911FADF"),
						ContentLength:      to.Int64Ptr(16384),
						ContentType:        to.StringPtr("image/jpeg"),
						ContentEncoding:    to.StringPtr(""),
						ContentLanguage:    to.StringPtr(""),
						ContentMD5:         to.StringPtr(""),
						CacheControl:       to.StringPtr(""),
						BlobSequenceNumber: to.Int32Ptr(3),
						BlobType:           BlobTypePageBlob.ToPtr(),
						LeaseStatus:        LeaseStatusTypeLocked.ToPtr(),
					},
				},
			},
		},
		ContainerName: to.StringPtr("https://myaccount.blob.core.windows.net/mycontainer"),
		NextMarker:    to.StringPtr(""),
	}
	if r := cmp.Diff(result.ListBlobsResponse, expected); r != "" {
		t.Fatal(r)
	}
}

func TestListContainers(t *testing.T) {
	client := newXMLClient()
	result, err := client.ListContainers(context.Background(), nil)
	require.NoError(t, err)
	expected := ListContainersResponse{
		ServiceEndpoint: to.StringPtr("https://myaccount.blob.core.windows.net/"),
		MaxResults:      to.Int32Ptr(3),
		NextMarker:      to.StringPtr("video"),
		Containers: []*Container{
			{
				Name: to.StringPtr("audio"),
				Properties: &ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7C6B1B2"),
					PublicAccess: PublicAccessTypeContainer.ToPtr(),
				},
			},
			{
				Name: to.StringPtr("images"),
				Properties: &ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7C1EEEC"),
				},
			},
			{
				Name: to.StringPtr("textfiles"),
				Properties: &ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7BACAC3"),
				},
			},
		},
	}
	if r := cmp.Diff(result.ListContainersResponse, expected); r != "" {
		t.Fatal(r)
	}
}

func TestPutACLs(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutACLs(context.Background(), []*SignedIdentifier{
		{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutBinary(t *testing.T) {
	client := newXMLClient()
	_, err := client.PutBinary(context.Background(), ModelWithByteProperty{
		Bytes: []byte("Hello world"),
	}, nil)
	require.NoError(t, err)
}

func TestPutComplexTypeRefNoMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutComplexTypeRefNoMeta(context.Background(), RootWithRefAndNoMeta{
		RefToModel: &ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutComplexTypeRefWithMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutComplexTypeRefWithMeta(context.Background(), RootWithRefAndMeta{
		RefToModel: &ComplexTypeWithMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyChildElement(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyChildElement(context.Background(), Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyList(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyList(context.Background(), Slideshow{
		Slides: []*Slide{},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyRootList(context.Background(), []*Banana{}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyWrappedLists(context.Background(), AppleBarrel{
		BadApples:  []*string{},
		GoodApples: []*string{},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutRootList(context.Background(), []*Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		{
			Name:       to.StringPtr("Plantain"),
			Flavor:     to.StringPtr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutRootListSingleItem(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutRootListSingleItem(context.Background(), []*Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutServiceProperties(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutServiceProperties(context.Background(), StorageServiceProperties{
		HourMetrics: &Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeAPIs: to.BoolPtr(false),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		Logging: &Logging{
			Version: to.StringPtr("1.0"),
			Delete:  to.BoolPtr(true),
			Read:    to.BoolPtr(false),
			Write:   to.BoolPtr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		MinuteMetrics: &Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeAPIs: to.BoolPtr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutSimple(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutSimple(context.Background(), Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: []*Slide{
			{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			{
				Items: to.StringPtrArray("Why WonderWidgets are great", "", "Who buys WonderWidgets"),
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutWrappedLists(context.Background(), AppleBarrel{
		BadApples:  to.StringPtrArray("Red Delicious"),
		GoodApples: to.StringPtrArray("Fuji", "Gala"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
