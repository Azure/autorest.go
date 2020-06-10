// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgrouptest

import (
	"context"
	"generatortests/autorest/generated/urlgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"
)

func getPathsOperations(t *testing.T) urlgroup.PathsOperations {
	client, err := urlgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create enum client: %v", err)
	}
	return client.PathsOperations()
}

func TestArrayCSVInPath(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.ArrayCSVInPath(context.Background(), []string{"ArrayPath1", "begin!*'();:@ &=+$,/?#[]end", "", ""})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsBase64URL(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.Base64URL(context.Background(), []byte("lorem"))
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsByteEmpty(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.ByteEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsByteMultiByte(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.ByteMultiByte(context.Background(), []byte("啊齄丂狛狜隣郎隣兀﨩"))
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// TODO: check
func TestPathsByteNull(t *testing.T) {
	client := getPathsOperations(t)
	_, err := client.ByteNull(context.Background(), nil)
	if err == nil {
		t.Fatalf("Did not receive an error, but expected one")
	}
}

func TestPathsDateNull(t *testing.T) {
	client := getPathsOperations(t)
	var time time.Time
	result, err := client.DateNull(context.Background(), time)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusBadRequest)
}

func TestPathsDateTimeNull(t *testing.T) {
	client := getPathsOperations(t)
	var time time.Time
	result, err := client.DateTimeNull(context.Background(), time)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusBadRequest)
}

func TestPathsDateTimeValid(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.DateTimeValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsDateValid(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.DateValid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsDoubleDecimalNegative(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.DoubleDecimalNegative(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsDoubleDecimalPositive(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.DoubleDecimalPositive(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsEnumNull(t *testing.T) {
	client := getPathsOperations(t)
	var color urlgroup.UriColor
	_, err := client.EnumNull(context.Background(), color)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
}

func TestPathsEnumValid(t *testing.T) {
	client := getPathsOperations(t)
	color := urlgroup.UriColorGreenColor
	result, err := client.EnumValid(context.Background(), color)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsFloatScientificNegative(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.FloatScientificNegative(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsFloatScientificPositive(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.FloatScientificPositive(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetBooleanFalse(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.GetBooleanFalse(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetBooleanTrue(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.GetBooleanTrue(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetIntNegativeOneMillion(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.GetIntNegativeOneMillion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetIntOneMillion(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.GetIntOneMillion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetNegativeTenBillion(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.GetNegativeTenBillion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetTenBillion(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.GetTenBillion(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringEmpty(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.StringEmpty(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringNull(t *testing.T) {
	client := getPathsOperations(t)
	var s string
	_, err := client.StringNull(context.Background(), s)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
}

func TestPathsStringURLEncoded(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.StringURLEncoded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringURLNonEncoded(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.StringURLNonEncoded(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringUnicode(t *testing.T) {
	client := getPathsOperations(t)
	result, err := client.StringUnicode(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsUnixTimeURL(t *testing.T) {
	client := getPathsOperations(t)
	d, err := time.Parse("2006-01-02", "2016-04-13")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.UnixTimeURL(context.Background(), d)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
