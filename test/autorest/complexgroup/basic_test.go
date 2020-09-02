// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgrouptest

import (
	"context"
	"generatortests/autorest/generated/complexgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newBasicClient() complexgroup.BasicOperations {
	return complexgroup.NewBasicClient(complexgroup.NewDefaultClient(nil))
}

func TestBasicGetValid(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetValid(context.Background())
	if err != nil {
		t.Fatalf("GetValid: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.Basic, &complexgroup.Basic{ID: to.Int32Ptr(2), Name: to.StringPtr("abc"), Color: complexgroup.CMYKColorsYellow.ToPtr()})
}

func TestBasicPutValid(t *testing.T) {
	client := newBasicClient()
	result, err := client.PutValid(context.Background(), complexgroup.Basic{
		ID:    to.Int32Ptr(2),
		Name:  to.StringPtr("abc"),
		Color: complexgroup.CMYKColorsMagenta.ToPtr(),
	})
	if err != nil {
		t.Fatalf("PutValid: %v", err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

func TestBasicGetInvalid(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetInvalid(context.Background())
	if err == nil {
		t.Fatal("GetInvalid expected an error")
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

func TestBasicGetEmpty(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetEmpty(context.Background())
	if err != nil {
		t.Fatalf("GetEmpty: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.Basic, &complexgroup.Basic{})
}

func TestBasicGetNull(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetNull(context.Background())
	if err != nil {
		t.Fatalf("GetNull: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.Basic, &complexgroup.Basic{})
}

func TestBasicGetNotProvided(t *testing.T) {
	client := newBasicClient()
	result, err := client.GetNotProvided(context.Background())
	if err != nil {
		t.Fatalf("GetNotProvided: %v", err)
	}
	helpers.DeepEqualOrFatal(t, result.Basic, (*complexgroup.Basic)(nil))
}
