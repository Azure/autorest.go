// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup

import (
	"context"
	"encoding/xml"
	"generatortests"
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

func newXMLClient(t *testing.T) *XMLClient {
	options := azcore.ClientOptions{
		Logging: policy.LogOptions{
			IncludeBody: true,
		},
		TracingProvider: generatortests.NewTracingProvider(t),
	}
	client, err := NewXMLClient(&options)
	require.NoError(t, err)
	return client
}

func TestGetACLs(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetACLs(context.Background(), nil)
	require.NoError(t, err)
	expected := []*SignedIdentifier{
		{
			ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.Ptr("rwd"),
			},
		},
	}
	if r := cmp.Diff(result.SignedIdentifiers, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetBytes(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetBytes(context.Background(), nil)
	require.NoError(t, err)
	if string(result.Bytes) != "Hello world" {
		t.Fatalf("unexpected bytes %s", string(result.Bytes))
	}
}

func TestGetComplexTypeRefNoMeta(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetComplexTypeRefNoMeta(context.Background(), nil)
	require.NoError(t, err)
	expected := RootWithRefAndNoMeta{
		RefToModel: &ComplexTypeNoMeta{
			ID: to.Ptr("myid"),
		},
		Something: to.Ptr("else"),
	}
	if r := cmp.Diff(result.RootWithRefAndNoMeta, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetComplexTypeRefWithMeta(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetComplexTypeRefWithMeta(context.Background(), nil)
	require.NoError(t, err)
	expected := RootWithRefAndMeta{
		RefToModel: &ComplexTypeWithMeta{
			ID: to.Ptr("myid"),
		},
		Something: to.Ptr("else"),
	}
	if r := cmp.Diff(result.RootWithRefAndMeta, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetEmptyChildElement(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetEmptyChildElement(context.Background(), nil)
	require.NoError(t, err)
	expected := Banana{
		Name:       to.Ptr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.Ptr(""),
	}
	if r := cmp.Diff(result.Banana, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetEmptyList(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetEmptyList(context.Background(), nil)
	require.NoError(t, err)
	expected := Slideshow{}
	if r := cmp.Diff(result.Slideshow, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetEmptyRootList(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetEmptyRootList(context.Background(), nil)
	require.NoError(t, err)
	if result.Bananas != nil {
		t.Fatal("expected nil slice")
	}
}

func TestGetEmptyWrappedLists(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetEmptyWrappedLists(context.Background(), nil)
	require.NoError(t, err)
	expected := AppleBarrel{}
	if r := cmp.Diff(result.AppleBarrel, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetHeaders(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetHeaders(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.CustomHeader, to.Ptr("custom-value")); r != "" {
		t.Fatal(r)
	}
}

func TestGetRootList(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetRootList(context.Background(), nil)
	require.NoError(t, err)
	expected := []*Banana{
		{
			Name:       to.Ptr("Cavendish"),
			Flavor:     to.Ptr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		{
			Name:       to.Ptr("Plantain"),
			Flavor:     to.Ptr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	if r := cmp.Diff(result.Bananas, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetRootListSingleItem(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetRootListSingleItem(context.Background(), nil)
	require.NoError(t, err)
	expected := []*Banana{
		{
			Name:       to.Ptr("Cavendish"),
			Flavor:     to.Ptr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	if r := cmp.Diff(result.Bananas, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetServiceProperties(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetServiceProperties(context.Background(), nil)
	require.NoError(t, err)
	expected := StorageServiceProperties{
		HourMetrics: &Metrics{
			Version:     to.Ptr("1.0"),
			Enabled:     to.Ptr(true),
			IncludeAPIs: to.Ptr(false),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr[int32](7),
			},
		},
		Logging: &Logging{
			Version: to.Ptr("1.0"),
			Delete:  to.Ptr(true),
			Read:    to.Ptr(false),
			Write:   to.Ptr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr[int32](7),
			},
		},
		MinuteMetrics: &Metrics{
			Version:     to.Ptr("1.0"),
			Enabled:     to.Ptr(true),
			IncludeAPIs: to.Ptr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr[int32](7),
			},
		},
	}
	if r := cmp.Diff(result.StorageServiceProperties, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetSimple(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetSimple(context.Background(), nil)
	require.NoError(t, err)
	expected := Slideshow{
		Author: to.Ptr("Yours Truly"),
		Date:   to.Ptr("Date of publication"),
		Title:  to.Ptr("Sample Slide Show"),
		Slides: []*Slide{
			{
				Title: to.Ptr("Wake up to WonderWidgets!"),
				Type:  to.Ptr("all"),
			},
			{
				Items: to.SliceOfPtrs("Why WonderWidgets are great", "", "Who buys WonderWidgets"),
				Title: to.Ptr("Overview"),
				Type:  to.Ptr("all"),
			},
		},
	}
	if r := cmp.Diff(result.Slideshow, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetWrappedLists(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetWrappedLists(context.Background(), nil)
	require.NoError(t, err)
	expected := AppleBarrel{
		BadApples:  to.SliceOfPtrs("Red Delicious"),
		GoodApples: to.SliceOfPtrs("Fuji", "Gala"),
	}
	if r := cmp.Diff(result.AppleBarrel, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetXMsText(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.GetXMsText(context.Background(), nil)
	require.NoError(t, err)
	expected := ObjectWithXMsTextProperty{
		Content:  to.Ptr("I am text"),
		Language: to.Ptr("english"),
	}
	if r := cmp.Diff(result.ObjectWithXMsTextProperty, expected); r != "" {
		t.Fatal(r)
	}
}

func TestJSONInput(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.JSONInput(context.Background(), JSONInput{
		ID: to.Ptr[int32](42),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestJSONOutput(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.JSONOutput(context.Background(), nil)
	require.NoError(t, err)
	expected := JSONOutput{
		ID: to.Ptr[int32](42),
	}
	if r := cmp.Diff(result.JSONOutput, expected); r != "" {
		t.Fatal(r)
	}
}

func TestListBlobs(t *testing.T) {
	client := newXMLClient(t)
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
						"color":            to.Ptr("blue"),
						"blobnumber":       to.Ptr("01"),
						"somemetadataname": to.Ptr("SomeMetadataValue"),
					},
					Name: to.Ptr("blob1.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.Ptr("0x8CBFF45D8A29A19"),
						ContentLength:   to.Ptr[int64](100),
						ContentType:     to.Ptr("text/html"),
						ContentEncoding: to.Ptr(""),
						ContentLanguage: to.Ptr("en-US"),
						ContentMD5:      to.Ptr(""),
						CacheControl:    to.Ptr("no-cache"),
						BlobType:        to.Ptr(BlobTypeBlockBlob),
						LeaseStatus:     to.Ptr(LeaseStatusTypeUnlocked),
					},
				},
				{
					Metadata: map[string]*string{
						"color":             to.Ptr("green"),
						"blobnumber":        to.Ptr("02"),
						"somemetadataname":  to.Ptr("SomeMetadataValue"),
						"x-ms-invalid-name": to.Ptr("nasdf$@#$$"),
					},
					Name: to.Ptr("blob2.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.Ptr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Ptr[int64](5000),
						ContentType:     to.Ptr("application/octet-stream"),
						ContentEncoding: to.Ptr("gzip"),
						ContentLanguage: to.Ptr(""),
						ContentMD5:      to.Ptr(""),
						CacheControl:    to.Ptr(""),
						BlobType:        to.Ptr(BlobTypeBlockBlob),
					},
					Snapshot: to.Ptr("2009-09-09T09:20:03.0427659Z"),
				},
				{
					Metadata: map[string]*string{
						"color":            to.Ptr("green"),
						"blobnumber":       to.Ptr("02"),
						"somemetadataname": to.Ptr("SomeMetadataValue"),
					},
					Name: to.Ptr("blob2.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.Ptr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Ptr[int64](5000),
						ContentType:     to.Ptr("application/octet-stream"),
						ContentEncoding: to.Ptr("gzip"),
						ContentLanguage: to.Ptr(""),
						ContentMD5:      to.Ptr(""),
						CacheControl:    to.Ptr(""),
						BlobType:        to.Ptr(BlobTypeBlockBlob),
					},
					Snapshot: to.Ptr("2009-09-09T09:20:03.1587543Z"),
				},
				{
					Metadata: map[string]*string{
						"color":            to.Ptr("green"),
						"blobnumber":       to.Ptr("02"),
						"somemetadataname": to.Ptr("SomeMetadataValue"),
					},
					Name: to.Ptr("blob2.txt"),
					Properties: &BlobProperties{
						LastModified:    &blob1LM,
						Etag:            to.Ptr("0x8CBFF45D8B4C212"),
						ContentLength:   to.Ptr[int64](5000),
						ContentType:     to.Ptr("application/octet-stream"),
						ContentEncoding: to.Ptr("gzip"),
						ContentLanguage: to.Ptr(""),
						ContentMD5:      to.Ptr(""),
						CacheControl:    to.Ptr(""),
						BlobType:        to.Ptr(BlobTypeBlockBlob),
						LeaseStatus:     to.Ptr(LeaseStatusTypeUnlocked),
					},
				},
				{
					Metadata: map[string]*string{
						"color":            to.Ptr("yellow"),
						"blobnumber":       to.Ptr("03"),
						"somemetadataname": to.Ptr("SomeMetadataValue"),
					},
					Name: to.Ptr("blob3.txt"),
					Properties: &BlobProperties{
						LastModified:       &blob2LM,
						Etag:               to.Ptr("0x8CBFF45D911FADF"),
						ContentLength:      to.Ptr[int64](16384),
						ContentType:        to.Ptr("image/jpeg"),
						ContentEncoding:    to.Ptr(""),
						ContentLanguage:    to.Ptr(""),
						ContentMD5:         to.Ptr(""),
						CacheControl:       to.Ptr(""),
						BlobSequenceNumber: to.Ptr[int32](3),
						BlobType:           to.Ptr(BlobTypePageBlob),
						LeaseStatus:        to.Ptr(LeaseStatusTypeLocked),
					},
				},
			},
		},
		ContainerName: to.Ptr("https://myaccount.blob.core.windows.net/mycontainer"),
		NextMarker:    to.Ptr(""),
	}
	if r := cmp.Diff(result.ListBlobsResponse, expected); r != "" {
		t.Fatal(r)
	}
}

func TestListContainers(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.ListContainers(context.Background(), nil)
	require.NoError(t, err)
	expected := ListContainersResponse{
		ServiceEndpoint: to.Ptr("https://myaccount.blob.core.windows.net/"),
		MaxResults:      to.Ptr[int32](3),
		NextMarker:      to.Ptr("video"),
		Containers: []*Container{
			{
				Name: to.Ptr("audio"),
				Properties: &ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.Ptr("0x8CACB9BD7C6B1B2"),
					PublicAccess: to.Ptr(PublicAccessTypeContainer),
				},
			},
			{
				Name: to.Ptr("images"),
				Properties: &ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.Ptr("0x8CACB9BD7C1EEEC"),
				},
			},
			{
				Name: to.Ptr("textfiles"),
				Properties: &ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.Ptr("0x8CACB9BD7BACAC3"),
				},
			},
		},
	}
	if r := cmp.Diff(result.ListContainersResponse, expected); r != "" {
		t.Fatal(r)
	}
}

func TestPutACLs(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutACLs(context.Background(), []*SignedIdentifier{
		{
			ID: to.Ptr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.Ptr("rwd"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutBinary(t *testing.T) {
	client := newXMLClient(t)
	_, err := client.PutBinary(context.Background(), ModelWithByteProperty{
		Bytes: []byte("Hello world"),
	}, nil)
	require.NoError(t, err)
}

func TestPutComplexTypeRefNoMeta(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutComplexTypeRefNoMeta(context.Background(), RootWithRefAndNoMeta{
		RefToModel: &ComplexTypeNoMeta{
			ID: to.Ptr("myid"),
		},
		Something: to.Ptr("else"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutComplexTypeRefWithMeta(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutComplexTypeRefWithMeta(context.Background(), RootWithRefAndMeta{
		RefToModel: &ComplexTypeWithMeta{
			ID: to.Ptr("myid"),
		},
		Something: to.Ptr("else"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyChildElement(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutEmptyChildElement(context.Background(), Banana{
		Name:       to.Ptr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.Ptr(""),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyList(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutEmptyList(context.Background(), Slideshow{
		Slides: []*Slide{},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyRootList(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutEmptyRootList(context.Background(), []*Banana{}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutEmptyWrappedLists(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutEmptyWrappedLists(context.Background(), AppleBarrel{
		BadApples:  []*string{},
		GoodApples: []*string{},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutRootList(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutRootList(context.Background(), []*Banana{
		{
			Name:       to.Ptr("Cavendish"),
			Flavor:     to.Ptr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		{
			Name:       to.Ptr("Plantain"),
			Flavor:     to.Ptr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutRootListSingleItem(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutRootListSingleItem(context.Background(), []*Banana{
		{
			Name:       to.Ptr("Cavendish"),
			Flavor:     to.Ptr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutServiceProperties(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutServiceProperties(context.Background(), StorageServiceProperties{
		HourMetrics: &Metrics{
			Version:     to.Ptr("1.0"),
			Enabled:     to.Ptr(true),
			IncludeAPIs: to.Ptr(false),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr[int32](7),
			},
		},
		Logging: &Logging{
			Version: to.Ptr("1.0"),
			Delete:  to.Ptr(true),
			Read:    to.Ptr(false),
			Write:   to.Ptr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr[int32](7),
			},
		},
		MinuteMetrics: &Metrics{
			Version:     to.Ptr("1.0"),
			Enabled:     to.Ptr(true),
			IncludeAPIs: to.Ptr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.Ptr(true),
				Days:    to.Ptr[int32](7),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutSimple(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutSimple(context.Background(), Slideshow{
		Author: to.Ptr("Yours Truly"),
		Date:   to.Ptr("Date of publication"),
		Title:  to.Ptr("Sample Slide Show"),
		Slides: []*Slide{
			{
				Title: to.Ptr("Wake up to WonderWidgets!"),
				Type:  to.Ptr("all"),
			},
			{
				Items: to.SliceOfPtrs("Why WonderWidgets are great", "", "Who buys WonderWidgets"),
				Title: to.Ptr("Overview"),
				Type:  to.Ptr("all"),
			},
		},
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPutWrappedLists(t *testing.T) {
	client := newXMLClient(t)
	result, err := client.PutWrappedLists(context.Background(), AppleBarrel{
		BadApples:  to.SliceOfPtrs("Red Delicious"),
		GoodApples: to.SliceOfPtrs("Fuji", "Gala"),
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestMetadataWithEmptyValue(t *testing.T) {
	const data1 = `<Container><Metadata><key1 /><key2>value2</key2></Metadata></Container>`
	c := Container{}
	err := xml.Unmarshal([]byte(data1), &c)
	require.NoError(t, err)
	require.Len(t, c.Metadata, 2)
	require.Empty(t, *c.Metadata["key1"])
	require.EqualValues(t, "value2", *c.Metadata["key2"])

	const data2 = `<Container><Metadata><key2>value2</key2><key1 /></Metadata></Container>`
	c = Container{}
	err = xml.Unmarshal([]byte(data2), &c)
	require.NoError(t, err)
	require.Len(t, c.Metadata, 2)
	require.Empty(t, *c.Metadata["key1"])
	require.EqualValues(t, "value2", *c.Metadata["key2"])

	const data3 = `<Container><Metadata><key1 /></Metadata></Container>`
	c = Container{}
	err = xml.Unmarshal([]byte(data3), &c)
	require.NoError(t, err)
	require.Len(t, c.Metadata, 1)
	require.Empty(t, *c.Metadata["key1"])
}
