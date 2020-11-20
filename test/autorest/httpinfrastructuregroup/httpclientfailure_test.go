// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"testing"
)

func newHTTPClientFailureClient() HTTPClientFailureClient {
	return NewHTTPClientFailureClient(NewDefaultConnection(nil))
}

func TestHTTPClientFailureDelete400(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Delete400(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureDelete407(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Delete407(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureDelete417(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Delete417(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet400(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Get400(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet402(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Get402(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet403(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Get403(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet411(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Get411(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet412(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Get412(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureGet416(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Get416(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead400(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Head400(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead401(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Head401(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead410(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Head410(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureHead429(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Head429(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureOptions400(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Options400(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureOptions403(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Options403(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailureOptions412(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Options412(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePatch400(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Patch400(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePatch405(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Patch405(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePatch414(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Patch414(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePost400(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Post400(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePost406(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Post406(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePost415(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Post415(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut400(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Put400(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut404(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Put404(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut409(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Put409(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}

func TestHTTPClientFailurePut413(t *testing.T) {
	client := newHTTPClientFailureClient()
	result, err := client.Put413(context.Background(), nil)
	if err == nil {
		t.Fatalf("Expected an error but did not receive one")
	}
	if result != nil {
		t.Fatalf("Expected a nil result")
	}
}
