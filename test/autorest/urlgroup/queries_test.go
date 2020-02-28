// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgrouptest

import (
	"context"
	"generatortests/autorest/generated/urlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getQueriesClient(t *testing.T) urlgroup.QueriesOperations {
	client, err := urlgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create enum client: %v", err)
	}
	return client.QueriesOperations()
}

func TestArrayStringCsvValid(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.ArrayStringCSVValid(context.Background(), &urlgroup.QueriesArrayStringCSVValidOptions{
		ArrayQuery: &[]string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestArrayStringPipesValid(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.ArrayStringPipesValid(context.Background(), &urlgroup.QueriesArrayStringPipesValidOptions{
		ArrayQuery: &[]string{"ArrayQuery1", "begin!*'();:@ &=+$,/?#[]end", "", ""},
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func toByteSlicePtr(v []byte) *[]byte {
	return &v
}

func TestByteMultiByte(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.ByteMultiByte(context.Background(), &urlgroup.QueriesByteMultiByteOptions{
		ByteQuery: toByteSlicePtr([]byte("啊齄丂狛狜隣郎隣兀﨩")),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestDateTimeValid(t *testing.T) {
	t.Skip("test fails, might be bug in test server")
	client := getQueriesClient(t)
	result, err := client.DateTimeValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestDoubleDecimalNegative(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.DoubleDecimalNegative(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestEnumValid(t *testing.T) {
	t.Skip("test fails, needs investigation")
	client := getQueriesClient(t)
	result, err := client.EnumValid(context.Background(), &urlgroup.QueriesEnumValidOptions{
		EnumQuery: urlgroup.UriColorGreencolor.ToPtr(),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestFloatScientificNegative(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.FloatScientificNegative(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestGetBooleanTrue(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.GetBooleanTrue(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestGetIntOneMillion(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.GetIntOneMillion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestGetTenBillion(t *testing.T) {
	client := getQueriesClient(t)
	result, err := client.GetTenBillion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}

func TestStringUnicode(t *testing.T) {
	t.Skip("test fails, needs investigation")
	client := getQueriesClient(t)
	result, err := client.StringUnicode(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}
