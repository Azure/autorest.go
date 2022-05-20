// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package complexgroup

import (
	"context"
	"encoding/json"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newPrimitiveClient() *PrimitiveClient {
	pl := runtime.NewPipeline(generatortests.ModuleName, generatortests.ModuleVersion, runtime.PipelineOptions{}, &azcore.ClientOptions{})
	return NewPrimitiveClient(pl)
}

func TestPrimitiveGetInt(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetInt(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.IntWrapper, IntWrapper{Field1: to.Ptr[int32](-1), Field2: to.Ptr[int32](2)}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutInt(t *testing.T) {
	client := newPrimitiveClient()
	a, b := int32(-1), int32(2)
	result, err := client.PutInt(context.Background(), IntWrapper{Field1: &a, Field2: &b}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetLong(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetLong(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.LongWrapper, LongWrapper{
		Field1: to.Ptr[int64](1099511627775),
		Field2: to.Ptr[int64](-999511627788),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutLong(t *testing.T) {
	client := newPrimitiveClient()
	a, b := int64(1099511627775), int64(-999511627788)
	result, err := client.PutLong(context.Background(), LongWrapper{Field1: &a, Field2: &b}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetFloat(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetFloat(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.FloatWrapper, FloatWrapper{
		Field1: to.Ptr[float32](1.05),
		Field2: to.Ptr[float32](-0.003),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutFloat(t *testing.T) {
	client := newPrimitiveClient()
	a, b := float32(1.05), float32(-0.003)
	result, err := client.PutFloat(context.Background(), FloatWrapper{Field1: &a, Field2: &b}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetDouble(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDouble(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DoubleWrapper, DoubleWrapper{
		Field1: to.Ptr(3e-100),
		Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: to.Ptr(-0.000000000000000000000000000000000000000000000000000000005),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDouble(t *testing.T) {
	client := newPrimitiveClient()
	a, b := float64(3e-100), float64(-0.000000000000000000000000000000000000000000000000000000005)
	result, err := client.PutDouble(context.Background(), DoubleWrapper{Field1: &a, Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose: &b}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetBool(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetBool(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.BooleanWrapper, BooleanWrapper{
		FieldFalse: to.Ptr(false),
		FieldTrue:  to.Ptr(true),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitiveGetByte(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetByte(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.ByteWrapper, ByteWrapper{Field: []byte{255, 254, 253, 252, 0, 250, 249, 248, 247, 246}}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutBool(t *testing.T) {
	client := newPrimitiveClient()
	a, b := true, false
	result, err := client.PutBool(context.Background(), BooleanWrapper{FieldTrue: &a, FieldFalse: &b}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestByteWrapperJSONNull(t *testing.T) {
	bw := ByteWrapper{}
	b, err := json.Marshal(bw)
	require.NoError(t, err)
	if string(b) != "{}" {
		t.Fatalf("unexpected value %s", string(b))
	}
	bw.Field = azcore.NullValue[[]byte]()
	b, err = json.Marshal(bw)
	require.NoError(t, err)
	if string(b) != `{"field":null}` {
		t.Fatalf("unexpected value %s", string(b))
	}
}

func TestPrimitivePutByte(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.PutByte(context.Background(), ByteWrapper{Field: []byte{255, 254, 253, 252, 0, 250, 249, 248, 247, 246}}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetString(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetString(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.StringWrapper, StringWrapper{
		Empty: to.Ptr(""),
		Field: to.Ptr("goodrequest"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutString(t *testing.T) {
	client := newPrimitiveClient()
	var c *string
	a, b, c := "goodrequest", "", nil
	result, err := client.PutString(context.Background(), StringWrapper{Field: &a, Empty: &b, Null: c}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetDate(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDate(context.Background(), nil)
	require.NoError(t, err)
	a, err := time.Parse("2006-01-02", "0001-01-01")
	require.NoError(t, err)
	b, err := time.Parse("2006-01-02", "2016-02-29")
	require.NoError(t, err)
	dw := DateWrapper{Field: &a, Leap: &b}
	if r := cmp.Diff(result.DateWrapper, dw); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDate(t *testing.T) {
	client := newPrimitiveClient()
	a, err := time.Parse("2006-01-02", "0001-01-01")
	require.NoError(t, err)
	b, err := time.Parse("2006-01-02", "2016-02-29")
	require.NoError(t, err)
	result, err := client.PutDate(context.Background(), DateWrapper{Field: &a, Leap: &b}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetDuration(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDuration(context.Background(), nil)
	require.NoError(t, err)
	if r := cmp.Diff(result.DurationWrapper, DurationWrapper{
		Field: to.Ptr("P123DT22H14M12.011S"),
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDuration(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.PutDuration(context.Background(), DurationWrapper{Field: to.Ptr("P123DT22H14M12.011S")}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitiveGetDateTime(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDateTime(context.Background(), nil)
	require.NoError(t, err)
	f, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	n, _ := time.Parse(time.RFC3339, "2015-05-18T18:38:00Z")
	if r := cmp.Diff(result.DatetimeWrapper, DatetimeWrapper{
		Field: &f,
		Now:   &n,
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitiveGetDateTimeRFC1123(t *testing.T) {
	client := newPrimitiveClient()
	result, err := client.GetDateTimeRFC1123(context.Background(), nil)
	require.NoError(t, err)
	f, _ := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	n, _ := time.Parse(time.RFC1123, "Mon, 18 May 2015 11:38:00 GMT")
	if r := cmp.Diff(result.Datetimerfc1123Wrapper, Datetimerfc1123Wrapper{
		Field: &f,
		Now:   &n,
	}); r != "" {
		t.Fatal(r)
	}
}

func TestPrimitivePutDateTime(t *testing.T) {
	client := newPrimitiveClient()
	f, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	n, _ := time.Parse(time.RFC3339, "2015-05-18T18:38:00Z")
	result, err := client.PutDateTime(context.Background(), DatetimeWrapper{
		Field: &f,
		Now:   &n,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestPrimitivePutDateTimeRFC1123(t *testing.T) {
	client := newPrimitiveClient()
	f, _ := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	n, _ := time.Parse(time.RFC1123, "Mon, 18 May 2015 11:38:00 GMT")
	result, err := client.PutDateTimeRFC1123(context.Background(), Datetimerfc1123Wrapper{
		Field: &f,
		Now:   &n,
	}, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

func TestDatetimeWrapper(t *testing.T) {
	now := time.Now()
	dtw := DatetimeWrapper{
		Field: azcore.NullValue[*time.Time](),
		Now:   &now,
	}
	b, err := json.Marshal(dtw)
	require.NoError(t, err)
	var dtw2 DatetimeWrapper
	if err = json.Unmarshal(b, &dtw2); err != nil {
		t.Fatal(err)
	}
	if dtw2.Field != nil {
		t.Fatal("expected nil Field")
	}
	if r := cmp.Diff(dtw2.Now, dtw.Now); r != "" {
		t.Fatal(r)
	}
}

func TestDatetimerfc1123Wrapper(t *testing.T) {
	now := time.Now()
	dtw := Datetimerfc1123Wrapper{
		Field: azcore.NullValue[*time.Time](),
		Now:   &now,
	}
	b, err := json.Marshal(dtw)
	require.NoError(t, err)
	var dtw2 Datetimerfc1123Wrapper
	if err = json.Unmarshal(b, &dtw2); err != nil {
		t.Fatal(err)
	}
	if dtw2.Field != nil {
		t.Fatal("expected nil Field")
	}
	if r := cmp.Diff(dtw2.Now.Format(time.RFC1123), dtw.Now.Format(time.RFC1123)); r != "" {
		t.Fatal(r)
	}
}

func TestDateWrapper(t *testing.T) {
	dw := DateWrapper{
		Field: azcore.NullValue[*time.Time](),
		Leap:  to.Ptr(time.Date(2021, 10, 22, 0, 0, 0, 0, time.UTC)),
	}
	b, err := json.Marshal(dw)
	require.NoError(t, err)
	var dw2 DateWrapper
	if err = json.Unmarshal(b, &dw2); err != nil {
		t.Fatal(err)
	}
	if dw2.Field != nil {
		t.Fatal("expected nil Field")
	}
	if r := cmp.Diff(dw2.Leap, dw.Leap); r != "" {
		t.Fatal(r)
	}
}
