// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package datetimerfc1123group

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func newDatetimerfc1123Client() *Datetimerfc1123Client {
	return NewDatetimerfc1123Client(NewDefaultConnection(nil))
}

func TestGetInvalid(t *testing.T) {
	client := newDatetimerfc1123Client()
	_, err := client.GetInvalid(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

func TestGetNull(t *testing.T) {
	client := newDatetimerfc1123Client()
	result, err := client.GetNull(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	var expected *time.Time
	if r := cmp.Diff(result.Value, expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetOverflow(t *testing.T) {
	client := newDatetimerfc1123Client()
	_, err := client.GetOverflow(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

// GetUTCLowercaseMaxDateTime - Get max datetime value fri, 31 dec 9999 23:59:59 gmt
func TestGetUTCLowercaseMaxDateTime(t *testing.T) {
	client := newDatetimerfc1123Client()
	result, err := client.GetUTCLowercaseMaxDateTime(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC1123, "Fri, 31 Dec 9999 23:59:59 GMT")
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

// GetUTCMinDateTime - Get min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func TestGetUTCMinDateTime(t *testing.T) {
	client := newDatetimerfc1123Client()
	result, err := client.GetUTCMinDateTime(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

// GetUTCUppercaseMaxDateTime - Get max datetime value FRI, 31 DEC 9999 23:59:59 GMT
func TestGetUTCUppercaseMaxDateTime(t *testing.T) {
	client := newDatetimerfc1123Client()
	result, err := client.GetUTCUppercaseMaxDateTime(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	expected, err := time.Parse(time.RFC1123, "FRI, 31 DEC 9999 23:59:59 GMT")
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, &expected); r != "" {
		t.Fatal(r)
	}
}

func TestGetUnderflow(t *testing.T) {
	client := newDatetimerfc1123Client()
	_, err := client.GetUnderflow(context.Background(), nil)
	if err == nil {
		t.Fatal("unexpected nil error")
	}
}

// PutUTCMaxDateTime - Put max datetime value Fri, 31 Dec 9999 23:59:59 GMT
func TestPutUTCMaxDateTime(t *testing.T) {
	client := newDatetimerfc1123Client()
	body, err := time.Parse(time.RFC1123, "Fri, 31 Dec 9999 23:59:59 GMT")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutUTCMaxDateTime(context.Background(), body, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}

// PutUTCMinDateTime - Put min datetime value Mon, 1 Jan 0001 00:00:00 GMT
func TestPutUTCMinDateTime(t *testing.T) {
	client := newDatetimerfc1123Client()
	body, err := time.Parse(time.RFC1123, "Mon, 01 Jan 0001 00:00:00 GMT")
	if err != nil {
		t.Fatal(err)
	}
	result, err := client.PutUTCMinDateTime(context.Background(), body, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := result.StatusCode; s != http.StatusOK {
		t.Fatalf("unexpected status code %d", s)
	}
}
