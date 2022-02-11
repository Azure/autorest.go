// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package urlgroup

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func newPathsClient() *PathsClient {
	return NewPathsClient(nil)
}

func TestArrayCSVInPath(t *testing.T) {
	client := newPathsClient()
	result, err := client.ArrayCSVInPath(context.Background(), []string{"ArrayPath1", "begin!*'();:@ &=+$,/?#[]end", "", ""}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsBase64URL(t *testing.T) {
	client := newPathsClient()
	result, err := client.Base64URL(context.Background(), []byte("lorem"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsByteEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.ByteEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsByteMultiByte(t *testing.T) {
	client := newPathsClient()
	result, err := client.ByteMultiByte(context.Background(), []byte("啊齄丂狛狜隣郎隣兀﨩"), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
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
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsDateTimeNull(t *testing.T) {
	client := newPathsClient()
	var time time.Time
	result, err := client.DateTimeNull(context.Background(), time, nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsDateTimeValid(t *testing.T) {
	client := newPathsClient()
	result, err := client.DateTimeValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsDateValid(t *testing.T) {
	client := newPathsClient()
	result, err := client.DateValid(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsDoubleDecimalNegative(t *testing.T) {
	client := newPathsClient()
	result, err := client.DoubleDecimalNegative(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsDoubleDecimalPositive(t *testing.T) {
	client := newPathsClient()
	result, err := client.DoubleDecimalPositive(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
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
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsFloatScientificNegative(t *testing.T) {
	client := newPathsClient()
	result, err := client.FloatScientificNegative(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsFloatScientificPositive(t *testing.T) {
	client := newPathsClient()
	result, err := client.FloatScientificPositive(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsGetBooleanFalse(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetBooleanFalse(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsGetBooleanTrue(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetBooleanTrue(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsGetIntNegativeOneMillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetIntNegativeOneMillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsGetIntOneMillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetIntOneMillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsGetNegativeTenBillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetNegativeTenBillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsGetTenBillion(t *testing.T) {
	client := newPathsClient()
	result, err := client.GetTenBillion(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsStringEmpty(t *testing.T) {
	client := newPathsClient()
	result, err := client.StringEmpty(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
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
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsStringURLNonEncoded(t *testing.T) {
	client := newPathsClient()
	result, err := client.StringURLNonEncoded(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}

func TestPathsStringUnicode(t *testing.T) {
	client := newPathsClient()
	result, err := client.StringUnicode(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
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
	if !reflect.ValueOf(result).IsZero() {
		t.Fatal("expected zero-value result")
	}
}
