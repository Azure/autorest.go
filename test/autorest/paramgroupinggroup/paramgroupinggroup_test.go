// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paramgroupinggroup

import (
	"context"
	"net/http"
	"testing"
)

func newParameterGroupingClient() *ParameterGroupingClient {
	return NewParameterGroupingClient(NewDefaultConnection(nil))
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
func TestPostMultiParamGroups(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostMultiParamGroups(context.Background(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PostOptional - Post a bunch of optional parameters grouped
func TestPostOptional(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostOptional(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PostRequired - Post a bunch of required parameters grouped
func TestPostRequired(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostRequired(context.Background(), ParameterGroupingPostRequiredParameters{
		Body:          1234,
		PathParameter: "path",
	})
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
func TestPostSharedParameterGroupObject(t *testing.T) {
	client := newParameterGroupingClient()
	result, err := client.PostSharedParameterGroupObject(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
