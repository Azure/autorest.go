// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package xmlrouptest

import (
	"context"
	"generatortests/autorest/generated/xmlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getXMLClient(t *testing.T) xmlgroup.XMLOperations {
	client, err := xmlgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create byte client: %v", err)
	}
	return client.XMLOperations()
}

func toTimePtr(layout string, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

func TestGetACLs(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetACLs(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]xmlgroup.SignedIDentifier{
		xmlgroup.SignedIDentifier{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &xmlgroup.AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.SignedIdentifiers, expected)
}

func TestGetComplexTypeRefNoMeta(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetComplexTypeRefNoMeta(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.RootWithRefAndNoMeta{
		RefToModel: &xmlgroup.ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	}
	helpers.DeepEqualOrFatal(t, result.RootWithRefAndNoMeta, expected)
}

func TestGetComplexTypeRefWithMeta(t *testing.T) {
	t.Skip("nyi")
}

func TestGetEmptyChildElement(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetEmptyChildElement(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	}
	helpers.DeepEqualOrFatal(t, result.Banana, expected)
}

func TestGetEmptyList(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetEmptyList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.Slideshow{}
	helpers.DeepEqualOrFatal(t, result.Slideshow, expected)
}

func TestGetEmptyRootList(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetEmptyRootList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	if result.Bananas != nil {
		t.Fatal("expected nil slice")
	}
}

func TestGetEmptyWrappedLists(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetEmptyWrappedLists(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.AppleBarrel{}
	helpers.DeepEqualOrFatal(t, result.AppleBarrel, expected)
}

func TestGetHeaders(t *testing.T) {
	t.Skip("nyi")
}

func TestGetRootList(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetRootList(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]xmlgroup.Banana{
		xmlgroup.Banana{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		xmlgroup.Banana{
			Name:       to.StringPtr("Plantain"),
			Flavor:     to.StringPtr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	helpers.DeepEqualOrFatal(t, result.Bananas, expected)
}

func TestGetRootListSingleItem(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetRootListSingleItem(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &[]xmlgroup.Banana{
		xmlgroup.Banana{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	}
	helpers.DeepEqualOrFatal(t, result.Bananas, expected)
}

func TestGetServiceProperties(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetServiceProperties(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.StorageServiceProperties{
		HourMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(false),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		Logging: &xmlgroup.Logging{
			Version: to.StringPtr("1.0"),
			Delete:  to.BoolPtr(true),
			Read:    to.BoolPtr(false),
			Write:   to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		MinuteMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.StorageServiceProperties, expected)
}

func TestGetSimple(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetSimple(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: &[]xmlgroup.Slide{
			xmlgroup.Slide{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			xmlgroup.Slide{
				Items: &[]string{"Why WonderWidgets are great", "", "Who buys WonderWidgets"},
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.Slideshow, expected)
}

func TestGetWrappedLists(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.GetWrappedLists(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.AppleBarrel{
		BadApples:  &[]string{"Red Delicious"},
		GoodApples: &[]string{"Fuji", "Gala"},
	}
	helpers.DeepEqualOrFatal(t, result.AppleBarrel, expected)
}

func TestJSONInput(t *testing.T) {
	t.Skip("nyi")
}

func TestJSONOutput(t *testing.T) {
	t.Skip("nyi")
}

func TestListBlobs(t *testing.T) {
	t.Skip("nyi")
}

func TestListContainers(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.ListContainers(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
	expected := &xmlgroup.ListContainersResponse{
		ServiceEndpoint: to.StringPtr("https://myaccount.blob.core.windows.net/"),
		MaxResults:      to.Int32Ptr(3),
		NextMarker:      to.StringPtr("video"),
		Containers: &[]xmlgroup.Container{
			xmlgroup.Container{
				Name: to.StringPtr("audio"),
				Properties: &xmlgroup.ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7C6B1B2"),
					PublicAccess: xmlgroup.PublicAccessTypeContainer.ToPtr(),
				},
			},
			xmlgroup.Container{
				Name: to.StringPtr("images"),
				Properties: &xmlgroup.ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7C1EEEC"),
				},
			},
			xmlgroup.Container{
				Name: to.StringPtr("textfiles"),
				Properties: &xmlgroup.ContainerProperties{
					LastModified: toTimePtr(time.RFC1123, "Wed, 26 Oct 2016 20:39:39 GMT"),
					Etag:         to.StringPtr("0x8CACB9BD7BACAC3"),
				},
			},
		},
	}
	helpers.DeepEqualOrFatal(t, result.EnumerationResults, expected)
}

func TestPutACLs(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutACLs(context.Background(), []xmlgroup.SignedIDentifier{
		xmlgroup.SignedIDentifier{
			ID: to.StringPtr("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI="),
			AccessPolicy: &xmlgroup.AccessPolicy{
				Start:      toTimePtr(time.RFC3339Nano, "2009-09-28T08:49:37.123Z"),
				Expiry:     toTimePtr(time.RFC3339Nano, "2009-09-29T08:49:37.123Z"),
				Permission: to.StringPtr("rwd"),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutComplexTypeRefNoMeta(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutComplexTypeRefNoMeta(context.Background(), xmlgroup.RootWithRefAndNoMeta{
		RefToModel: &xmlgroup.ComplexTypeNoMeta{
			ID: to.StringPtr("myid"),
		},
		Something: to.StringPtr("else"),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutComplexTypeRefWithMeta(t *testing.T) {
	t.Skip("nyi")
}

func TestPutEmptyChildElement(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutEmptyChildElement(context.Background(), xmlgroup.Banana{
		Name:       to.StringPtr("Unknown Banana"),
		Expiration: toTimePtr(time.RFC3339Nano, "2012-02-24T00:53:52.789Z"),
		Flavor:     to.StringPtr(""),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutEmptyList(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutEmptyList(context.Background(), xmlgroup.Slideshow{
		Slides: &[]xmlgroup.Slide{},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutEmptyRootList(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutEmptyRootList(context.Background(), []xmlgroup.Banana{})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutEmptyWrappedLists(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutEmptyWrappedLists(context.Background(), xmlgroup.AppleBarrel{
		BadApples:  &[]string{},
		GoodApples: &[]string{},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutRootList(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutRootList(context.Background(), []xmlgroup.Banana{
		xmlgroup.Banana{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
		xmlgroup.Banana{
			Name:       to.StringPtr("Plantain"),
			Flavor:     to.StringPtr("Savory"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutRootListSingleItem(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutRootListSingleItem(context.Background(), []xmlgroup.Banana{
		xmlgroup.Banana{
			Name:       to.StringPtr("Cavendish"),
			Flavor:     to.StringPtr("Sweet"),
			Expiration: toTimePtr(time.RFC3339Nano, "2018-02-28T00:40:00.123Z"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutServiceProperties(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutServiceProperties(context.Background(), xmlgroup.StorageServiceProperties{
		HourMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(false),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		Logging: &xmlgroup.Logging{
			Version: to.StringPtr("1.0"),
			Delete:  to.BoolPtr(true),
			Read:    to.BoolPtr(false),
			Write:   to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
		MinuteMetrics: &xmlgroup.Metrics{
			Version:     to.StringPtr("1.0"),
			Enabled:     to.BoolPtr(true),
			IncludeApIs: to.BoolPtr(true),
			RetentionPolicy: &xmlgroup.RetentionPolicy{
				Enabled: to.BoolPtr(true),
				Days:    to.Int32Ptr(7),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutSimple(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutSimple(context.Background(), xmlgroup.Slideshow{
		Author: to.StringPtr("Yours Truly"),
		Date:   to.StringPtr("Date of publication"),
		Title:  to.StringPtr("Sample Slide Show"),
		Slides: &[]xmlgroup.Slide{
			xmlgroup.Slide{
				Title: to.StringPtr("Wake up to WonderWidgets!"),
				Type:  to.StringPtr("all"),
			},
			xmlgroup.Slide{
				Items: &[]string{"Why WonderWidgets are great", "", "Who buys WonderWidgets"},
				Title: to.StringPtr("Overview"),
				Type:  to.StringPtr("all"),
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}

func TestPutWrappedLists(t *testing.T) {
	client := getXMLClient(t)
	result, err := client.PutWrappedLists(context.Background(), xmlgroup.AppleBarrel{
		BadApples:  &[]string{"Red Delicious"},
		GoodApples: &[]string{"Fuji", "Gala"},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusCreated)
}
