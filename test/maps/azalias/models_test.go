package azalias

import (
	"encoding/json"
	"testing"
)

func TestPolicyAssignmentProperties(t *testing.T) {
	const payload = `{"displayName":"Not allowed resource types - Virtual Machine","metadata":{"one":{"value":{"key":"value"}}},"parameters":{"effect":{"value":"Audit"},"listOfResourceTypesNotAllowed":{"value":["Microsoft.Compute/virtualMachines"]}}}`

	paprops := PolicyAssignmentProperties{}
	if err := json.Unmarshal([]byte(payload), &paprops); err != nil {
		t.Fatal(err)
	}
	s, ok := paprops.Parameters["effect"].Value.(string)
	if !ok {
		t.Fatalf("unexpected type %T", paprops.Parameters["effect"].Value)
	}
	if s != "Audit" {
		t.Fatalf("got %s, want Audit", s)
	}
	sl, ok := paprops.Parameters["listOfResourceTypesNotAllowed"].Value.([]interface{})
	if !ok {
		t.Fatalf("unexpected type %T", paprops.Parameters["listOfResourceTypesNotAllowed"].Value)
	}
	if len(sl) != 1 {
		t.Fatal("unexpected slice len")
	}
	if sl[0] != "Microsoft.Compute/virtualMachines" {
		t.Fatalf("got %s, want Microsoft.Compute/virtualMachines", sl[0])
	}
	m, ok := paprops.Metadata["one"]
	if !ok {
		t.Fatal("missing one")
	}
	if v := m.Value["key"]; v != "value" {
		t.Fatalf("got %s want value", v)
	}
	b, err := json.Marshal(paprops)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != payload {
		t.Fatalf("got %s", string(b))
	}
}
