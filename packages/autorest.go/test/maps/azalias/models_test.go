// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azalias

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPolicyAssignmentProperties(t *testing.T) {
	const payload = `{"displayName":"Not allowed resource types - Virtual Machine","metadata":{"one":{"value":{"key":"value"}}},"parameters":{"effect":{"value":"Audit"},"listOfResourceTypesNotAllowed":{"value":["Microsoft.Compute/virtualMachines"]}}}`

	paprops := PolicyAssignmentProperties{}
	if err := json.Unmarshal([]byte(payload), &paprops); err != nil {
		t.Fatal(err)
	}
	var s string
	s, ok := paprops.Parameters["effect"].Value.(string)
	require.True(t, ok)
	if s != "Audit" {
		t.Fatalf("got %s, want Audit", s)
	}
	sl, ok := paprops.Parameters["listOfResourceTypesNotAllowed"].Value.([]any)
	require.True(t, ok)
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
	mm, ok := m.Value.(map[string]any)
	require.True(t, ok)
	if v := mm["key"]; v != "value" {
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

func TestUnmarshalFail(t *testing.T) {
	const data = `{"id": 123}`
	var geo GeoJSONFeature
	err := json.Unmarshal([]byte(data), &geo)
	require.Error(t, err)
	require.Equal(t, "unmarshalling type *azalias.GeoJSONFeature: struct field ID: json: cannot unmarshal number into Go value of type string", err.Error())
}

func TestTimeFormat(t *testing.T) {
	theTime, err := time.Parse(time.TimeOnly, "15:04:05.12345")
	require.NoError(t, err)
	source := TypeWithSliceOfTimes{
		Interval: &theTime,
	}
	data, err := json.Marshal(source)
	require.NoError(t, err)
	require.EqualValues(t, `{"interval":"15:04:05.12345Z","times":[]}`, string(data))
	dest := TypeWithSliceOfTimes{}
	require.NoError(t, json.Unmarshal([]byte(`{"interval": "15:04:05.12345"}`), &dest))
	require.NotNil(t, dest.Interval)
	require.EqualValues(t, theTime, *dest.Interval)
}

func TestDisallowedField(t *testing.T) {
	resp := &AliasesCreateResponse{}
	data := `{"aliasId":"theAlias","unknownField":"value"}`
	require.Error(t, json.Unmarshal([]byte(data), &resp))
}
