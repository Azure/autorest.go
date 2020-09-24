// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"
	"time"
)

func newPathsClient() PathsOperations {
	return NewPathsClient(NewDefaultClient(nil))
}

func TestArrayCSVInPath(t *testing.T) {
	client := newPathsClient()
	result, err := client.ArrayCSVInPath(context.Background(), []string{"ArrayPath1", "begin!*'();:@ &=+$,/?#[]end", "", ""}, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsBase64URL(t *testing.T) {
	client := newPathsClient()
	result, err := client.Base64URL(context.Background(), []byte("lorem"), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsByteEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.ByteEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsByteMultiByte(t *testing.T) {
	client := newPathsClient()
	result, err := client.ByteMultiByte(context.Background(), []byte("啊齄丂狛狜隣郎隣兀﨩"), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// TODO: check
func TestPathsByteNull(t *testing.T) {
	client := newPathsClient()
	_, err := client.ByteNull(context.Background(), nil, nil)
	if err == nil {
		t.Fatalf("Did not receive an error, but expected one")
	}
}

func TestPathsDateNull(t *testing.T) {
	client := newPathsClient()
	var time time.Time
	result, err := client.DateNull(context.Background(), time, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusBadRequest)
}

func TestPathsDateTimeNull(t *testing.T) {
	client := newPathsClient()
	var time time.Time
	result, err := client.DateTimeNull(context.Background(), time, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusBadRequest)
}

func TestPathsDateTimeValid(t *testing.T) {
	client := newPathsClient()
	result, err := client.DateTimeValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsDateValid(t *testing.T) {
	client := newPathsClient()
	result, err := client.DateValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsDoubleDecimalNegative(t *testing.T) {
	client := newPathsClient()
	result, err := client.DoubleDecimalNegative(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsDoubleDecimalPositive(t *testing.T) {
	client := newPathsClient()
	result, err := client.DoubleDecimalPositive(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsEnumNull(t *testing.T) {
	client := newPathsClient()
	var color URIColor
	_, err := client.EnumNull(context.Background(), color, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
}

func TestPathsEnumValid(t *testing.T) {
	client := newPathsClient()
	color := URIColorGreenColor
	result, err := client.EnumValid(context.Background(), color, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsFloatScientificNegative(t *testing.T) {
	client := newPathsClient()
	result, err := client.FloatScientificNegative(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsFloatScientificPositive(t *testing.T) {
	client := newPathsClient()
	result, err := client.FloatScientificPositive(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetBooleanFalse(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetBooleanFalse(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetBooleanTrue(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetBooleanTrue(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetIntNegativeOneMillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetIntNegativeOneMillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetIntOneMillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetIntOneMillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetNegativeTenBillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetNegativeTenBillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsGetTenBillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetTenBillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.StringEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringNull(t *testing.T) {
	client := newPathsClient()
	var s string
	_, err := client.StringNull(context.Background(), s, nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
}

func TestPathsStringURLEncoded(t *testing.T) {
	client := newPathsClient()
	result, err := client.StringURLEncoded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringURLNonEncoded(t *testing.T) {
	client := newPathsClient()
	result, err := client.StringURLNonEncoded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsStringUnicode(t *testing.T) {
	client := newPathsClient()
	result, err := client.StringUnicode(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestPathsUnixTimeURL(t *testing.T) {
	client := newPathsClient()
	d, err := time.Parse("2006-01-02", "2016-04-13")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.UnixTimeURL(context.Background(), d, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}
