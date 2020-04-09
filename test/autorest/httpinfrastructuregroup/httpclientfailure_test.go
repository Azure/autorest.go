// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregrouptest

import (
	"context"
	"generatortests/autorest/generated/httpinfrastructuregroup"
	"testing"
)

func getHTTPClientFailureOperations(t *testing.T) httpinfrastructuregroup.HTTPClientFailureOperations {
	client, err := httpinfrastructuregroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create HTTPClientFailure client: %v", err)
	}
	return client.HTTPClientFailureOperations()
}

func TestHTTPClientFailureDelete400(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Delete400(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureDelete407(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Delete407(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureDelete417(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Delete417(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet400(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Get400(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet402(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Get402(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet403(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Get403(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet411(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Get411(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet412(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Get412(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet416(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Get416(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead400(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Head400(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead401(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Head401(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead410(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Head410(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead429(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Head429(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureOptions400(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Options400(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureOptions403(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Options403(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureOptions412(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Options412(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePatch400(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Patch400(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePatch405(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Patch405(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePatch414(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Patch414(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePost400(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Post400(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePost406(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Post406(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePost415(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Post415(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut400(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Put400(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut404(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Put404(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut409(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Put409(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut413(t *testing.T) {
	client := getHTTPClientFailureOperations(t)
	result, err := client.Put413(context.Background())
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}
