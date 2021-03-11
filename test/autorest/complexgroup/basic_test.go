// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
	"github.com/google/go-cmp/cmp"
)

func newBasicClient() *BasicClient {
	return NewBasicClient(NewDefaultConnection(nil))
}

func TestBasicGetValid(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetValid(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	if r := cmp.Diff(result.Basic, &Basic{ID: to.Int32Ptr(2), Name: to.StringPtr("abc"), Color: CMYKColorsYELLOW.ToPtr()}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicPutValid(t *testing.T) {
	client := newBasicClient()
	result, err := client.PutValid(context.Background(), Basic{
		ID:    to.Int32Ptr(2),
		Name:  to.StringPtr("abc"),
		Color: CMYKColorsMagenta.ToPtr(),
	}, nil)
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

func TestBasicGetInvalid(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetInvalid(context.Background(), nil)
	if err == nil {
		t.Fatal("GetInvalid expected an error")
	}
	if r := cmp.Diff(result, BasicResponse{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetEmpty(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetEmpty(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	if r := cmp.Diff(result.Basic, &Basic{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetNull(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	if r := cmp.Diff(result.Basic, &Basic{}); r != "" {
		t.Fatal(r)
	}
}

func TestBasicGetNotProvided(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetNotProvided(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	if r := cmp.Diff(result.Basic, (*Basic)(nil)); r != "" {
		t.Fatal(r)
	}
}
