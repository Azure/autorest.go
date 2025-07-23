// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimerfc1123group

import (
	"context"
	"generatortests"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

func newDatetimerfc1123Client(t *testing.T) *Datetimerfc1123Client {
	client, err := NewDatetimerfc1123Client(generatortests.Host, &azcore.ClientOptions{
		TracingProvider: generatortests.NewTracingProvider(t),
	})
	require.NoError(t, err)
	return client
}

func TestGetInvalid(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	_, err := client.GetInvalid(context.Background(), nil)
	require.Error(t, err)
}

func TestGetNull(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	result, err := client.GetNull(context.Background(), nil)
	require.NoError(t, err)
	if result.Value != nil {
		t.Fatal("expected nil value")
	}
}

func TestGetOverflow(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	_, err := client.GetOverflow(context.Background(), nil)
	require.Error(t, err)
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
func TestGetUTCLowercaseMaxDateTime(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	result, err := client.GetUTCLowercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC1123, "Fri, 31 Dec 9999 23:59:59 GMT")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func TestGetUTCMinDateTime(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	result, err := client.GetUTCMinDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
func TestGetUTCUppercaseMaxDateTime(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	result, err := client.GetUTCUppercaseMaxDateTime(context.Background(), nil)
	require.NoError(t, err)
	expected, err := time.Parse(time.RFC1123, "FRI, 31 DEC 9999 23:59:59 GMT")
	require.NoError(t, err)
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetUnderflow(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	_, err := client.GetUnderflow(context.Background(), nil)
	require.Error(t, err)
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
func TestPutUTCMaxDateTime(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	body, err := time.Parse(time.RFC1123, "Fri, 31 Dec 9999 23:59:59 GMT")
	require.NoError(t, err)
	result, err := client.PutUTCMaxDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func TestPutUTCMinDateTime(t *testing.T) {
	client := newDatetimerfc1123Client(t)
	body, err := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	require.NoError(t, err)
	result, err := client.PutUTCMinDateTime(context.Background(), body, nil)
	require.NoError(t, err)
	require.Zero(t, result)
}
