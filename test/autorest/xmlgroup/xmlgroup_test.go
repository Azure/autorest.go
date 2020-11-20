// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func toTimePtr(layout string, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

func newXMLClient() XMLClient {
	return NewXMLClient(NewDefaultConnection(nil))
}

func TestGetACLs(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetACLs(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]SignedIDentifier{
		{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.SignedIdentifiers, expected)
}

func TestGetComplexTypeRefNoMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetComplexTypeRefNoMeta(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &RootWithRefAndNoMeta{
		RefToModel: &ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}
	helpers.DeepEqualOrFatal(t, result.RootWithRefAndNoMeta, expected)
}

func TestGetComplexTypeRefWithMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetComplexTypeRefWithMeta(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	expected := &RootWithRefAndMeta{
		RefToModel: &ComplexTypeWithMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}
	helpers.DeepEqualOrFatal(t, result.RootWithRefAndMeta, expected)
}

func TestGetEmptyChildElement(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyChildElement(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	}
	helpers.DeepEqualOrFatal(t, result.Banana, expected)
}

func TestGetEmptyList(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyList(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &Slideshow{}
	helpers.DeepEqualOrFatal(t, result.Slideshow, expected)
}

func TestGetEmptyRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyRootList(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	if result.Bananas != nil {
		t.Fatal("expected nil slice")
	}
}

func TestGetEmptyWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetEmptyWrappedLists(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &AppleBarrel{}
	helpers.DeepEqualOrFatal(t, result.AppleBarrel, expected)
}

func TestGetHeaders(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetHeaders(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.CustomHeader, to.StringPtr("custom-value"))
}

func TestGetRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetRootList(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]Banana{
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
	helpers.DeepEqualOrFatal(t, result.Bananas, expected)
}

func TestGetRootListSingleItem(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetRootListSingleItem(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	helpers.DeepEqualOrFatal(t, result.Bananas, expected)
}

func TestGetServiceProperties(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetServiceProperties(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &StorageServiceProperties{
		HourMetrics: &Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(false),
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
			IncludeApIs: to.BoolPtr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.StorageServiceProperties, expected)
}

func TestGetSimple(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetSimple(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: &[]Slide{
			{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			{
				Items: &[]string{"Why WonderWidgets are great", "", "Who buys WonderWidgets"},
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.Slideshow, expected)
}

func TestGetWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.GetWrappedLists(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &AppleBarrel{
		BadApples:  &[]string{"Red Delicious"},
		GoodApples: &[]string{"Fuji", "Gala"},
	}
	helpers.DeepEqualOrFatal(t, result.AppleBarrel, expected)
}

func TestJSONInput(t *testing.T) {
	client := newXMLClient()
	result, err := client.JSONInput(context.Background(), JSONInput{
		ID: to.Int32Ptr(42),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestJSONOutput(t *testing.T) {
	client := newXMLClient()
	result, err := client.JSONOutput(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	expected := JSONOutput{
		ID: to.Int32Ptr(42),
	}
	helpers.DeepEqualOrFatal(t, result.JSONOutput, &expected)
}

func TestListBlobs(t *testing.T) {
	client := newXMLClient()
	result, err := client.ListBlobs(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	blob1LM, err := time.Parse(time.RFC1123, "Wed, 09 Sep 2009 09:20:02 GMT")
	if err != nil {
		t.Fatal(err)
	}
	blob2LM, err := time.Parse(time.RFC1123, "Wed, 09 Sep 2009 09:20:03 GMT")
	if err != nil {
		t.Fatal(err)
	}
	expected := ListBlobsResponse{
		Blobs: &Blobs{
			Blob: &[]Blob{
				{
					Metadata: &map[string]string{
						"color":            "blue",
						"blobnumber":       "01",
						"somemetadataname": "SomeMetadataValue",
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
					Metadata: &map[string]string{
						"color":             "green",
						"blobnumber":        "02",
						"somemetadataname":  "SomeMetadataValue",
						"x-ms-invalid-name": "nasdf$@#$$",
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
					Metadata: &map[string]string{
						"color":            "green",
						"blobnumber":       "02",
						"somemetadataname": "SomeMetadataValue",
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
					Metadata: &map[string]string{
						"color":            "green",
						"blobnumber":       "02",
						"somemetadataname": "SomeMetadataValue",
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
					Metadata: &map[string]string{
						"color":            "yellow",
						"blobnumber":       "03",
						"somemetadataname": "SomeMetadataValue",
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
	helpers.DeepEqualOrFatal(t, result.EnumerationResults, &expected)
}

func TestListContainers(t *testing.T) {
	client := newXMLClient()
	result, err := client.ListContainers(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &ListContainersResponse{
		ServiceEndpoint: to.StringPtr("https://myaccount.blob.core.windows.net/"),
		MaxResults:      to.Int32Ptr(3),
		NextMarker:      to.StringPtr("video"),
		Containers: &[]Container{
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
	helpers.DeepEqualOrFatal(t, result.EnumerationResults, expected)
}

func TestPutACLs(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutACLs(context.Background(), []SignedIDentifier{
		{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutComplexTypeRefNoMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutComplexTypeRefNoMeta(context.Background(), RootWithRefAndNoMeta{
		RefToModel: &ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutComplexTypeRefWithMeta(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutComplexTypeRefWithMeta(context.Background(), RootWithRefAndMeta{
		RefToModel: &ComplexTypeWithMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyChildElement(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyChildElement(context.Background(), Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyList(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyList(context.Background(), Slideshow{
		Slides: &[]Slide{},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyRootList(context.Background(), []Banana{}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutEmptyWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutEmptyWrappedLists(context.Background(), AppleBarrel{
		BadApples:  &[]string{},
		GoodApples: &[]string{},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutRootList(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutRootList(context.Background(), []Banana{
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
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutRootListSingleItem(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutRootListSingleItem(context.Background(), []Banana{
		{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutServiceProperties(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutServiceProperties(context.Background(), StorageServiceProperties{
		HourMetrics: &Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(false),
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
			IncludeApIs: to.BoolPtr(true),
			RetentionPolicy: &RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutSimple(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutSimple(context.Background(), Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: &[]Slide{
			{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			{
				Items: &[]string{"Why WonderWidgets are great", "", "Who buys WonderWidgets"},
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}

func TestPutWrappedLists(t *testing.T) {
	client := newXMLClient()
	result, err := client.PutWrappedLists(context.Background(), AppleBarrel{
		BadApples:  &[]string{"Red Delicious"},
		GoodApples: &[]string{"Fuji", "Gala"},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusCreated)
}
