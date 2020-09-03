// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package httpinfrastructuregroup

import (
	"context"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func newMultipleResponsesClient() MultipleResponsesOperations {
	return NewMultipleResponsesClient(NewDefaultClient(nil))
}

// Get200Model201ModelDefaultError200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGet200Model201ModelDefaultError200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model201ModelDefaultError200Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	switch x := result.(type) {
	case *MyExceptionResponse:
		helpers.DeepEqualOrFatal(t, x.MyException.StatusCode, to.StringPtr("200"))
	case *BResponse:
		helpers.VerifyStatusCode(t, x.RawResponse, http.StatusCreated)
	default:
		t.Fatalf("unhandled response type %v", x)
	}
}

// Get200Model201ModelDefaultError201Valid - Send a 201 response with valid payload: {'statusCode': '201', 'textStatusCode': 'Created'}
func TestGet200Model201ModelDefaultError201Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model201ModelDefaultError201Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.(*BResponse)
	if !ok {
		t.Fatal("unexpected response type")
	}
	helpers.DeepEqualOrFatal(t, r.B, &B{
		MyException: MyException{
			StatusCode: to.StringPtr("201"),
		},
		TextStatusCode: to.StringPtr("Created"),
	})
}

// Get200Model201ModelDefaultError400Valid - Send a 400 response with valid payload: {'code': '400', 'message': 'client error'}
func TestGet200Model201ModelDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model201ModelDefaultError400Valid(context.Background())
	r, ok := err.(Error)
	if !ok {
		t.Fatal("unexpected error type")
	}
	helpers.DeepEqualOrFatal(t, r, Error{
		Message: to.StringPtr("client error"),
		Status:  to.Int32Ptr(400),
	})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// Get200Model204NoModelDefaultError200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGet200Model204NoModelDefaultError200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError200Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.MyException, &MyException{
		StatusCode: to.StringPtr("200"),
	})
}

// Get200Model204NoModelDefaultError201Invalid - Send a 201 response with valid payload: {'statusCode': '201'}
func TestGet200Model204NoModelDefaultError201Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError201Invalid(context.Background())
	r, ok := err.(Error)
	if !ok {
		t.Fatal("unexpected error type")
	}
	helpers.DeepEqualOrFatal(t, r, Error{})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// Get200Model204NoModelDefaultError202None - Send a 202 response with no payload:
func TestGet200Model204NoModelDefaultError202None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError202None(context.Background())
	r, ok := err.(Error)
	if !ok {
		t.Fatal("unexpected error type")
	}
	helpers.DeepEqualOrFatal(t, r, Error{})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// Get200Model204NoModelDefaultError204Valid - Send a 204 response with no payload
func TestGet200Model204NoModelDefaultError204Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError204Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusNoContent)
	if result.MyException != nil {
		t.Fatal("expected nil payload")
	}
}

// Get200Model204NoModelDefaultError400Valid - Send a 400 response with valid error payload: {'status': 400, 'message': 'client error'}
func TestGet200Model204NoModelDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200Model204NoModelDefaultError400Valid(context.Background())
	r, ok := err.(Error)
	if !ok {
		t.Fatal("unexpected error type")
	}
	helpers.DeepEqualOrFatal(t, r, Error{
		Message: to.StringPtr("client error"),
		Status:  to.Int32Ptr(400),
	})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// Get200ModelA200Invalid - Send a 200 response with invalid payload {'statusCodeInvalid': '200'}
func TestGet200ModelA200Invalid(t *testing.T) {
	t.Skip("payload doen't match schema doesn't cause unmarshalling error")
}

// Get200ModelA200None - Send a 200 response with no payload, when a payload is expected - client should return a null object of thde type for model A
func TestGet200ModelA200None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA200None(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if result.MyException != nil {
		t.Fatal("expected nil MyException")
	}
}

// Get200ModelA200Valid - Send a 200 response with payload {'statusCode': '200'}
func TestGet200ModelA200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA200Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.MyException, &MyException{
		StatusCode: to.StringPtr("200"),
	})
}

// Get200ModelA201ModelC404ModelDDefaultError200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGet200ModelA201ModelC404ModelDDefaultError200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError200Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.(*MyExceptionResponse)
	if !ok {
		t.Fatal("unexpected result type")
	}
	helpers.DeepEqualOrFatal(t, r.MyException, &MyException{
		StatusCode: to.StringPtr("200"),
	})
}

// Get200ModelA201ModelC404ModelDDefaultError201Valid - Send a 200 response with valid payload: {'httpCode': '201'}
func TestGet200ModelA201ModelC404ModelDDefaultError201Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError201Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.(*CResponse)
	if !ok {
		t.Fatal("unexpected result type")
	}
	helpers.DeepEqualOrFatal(t, r.C, &C{
		HTTPCode: to.StringPtr("201"),
	})
}

// Get200ModelA201ModelC404ModelDDefaultError400Valid - Send a 400 response with valid payload: {'code': '400', 'message': 'client error'}
func TestGet200ModelA201ModelC404ModelDDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError400Valid(context.Background())
	r, ok := err.(Error)
	if !ok {
		t.Fatal("unexpected error type")
	}
	helpers.DeepEqualOrFatal(t, r, Error{
		Message: to.StringPtr("client error"),
		Status:  to.Int32Ptr(400),
	})
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// Get200ModelA201ModelC404ModelDDefaultError404Valid - Send a 200 response with valid payload: {'httpStatusCode': '404'}
func TestGet200ModelA201ModelC404ModelDDefaultError404Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA201ModelC404ModelDDefaultError404Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	r, ok := result.(*DResponse)
	if !ok {
		t.Fatal("unexpected result type")
	}
	helpers.DeepEqualOrFatal(t, r.D, &D{
		HTTPStatusCode: to.StringPtr("404"),
	})
}

// Get200ModelA202Valid - Send a 202 response with payload {'statusCode': '202'}
func TestGet200ModelA202Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA200Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.MyException, &MyException{
		StatusCode: to.StringPtr("200"),
	})
}

// Get200ModelA400Invalid - Send a 200 response with invalid payload {'statusCodeInvalid': '400'}
func TestGet200ModelA400Invalid(t *testing.T) {
	t.Skip("payload doen't match schema doesn't cause unmarshalling error")
}

// Get200ModelA400None - Send a 400 response with no payload client should treat as an http error with no error model
func TestGet200ModelA400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA400None(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// Get200ModelA400Valid - Send a 200 response with payload {'statusCode': '400'}
func TestGet200ModelA400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get200ModelA400Valid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("expected nil result")
	}
}

// Get202None204NoneDefaultError202None - Send a 202 response with no payload
func TestGet202None204NoneDefaultError202None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultError202None(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusAccepted)
}

// Get202None204NoneDefaultError204None - Send a 204 response with no payload
func TestGet202None204NoneDefaultError204None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultError204None(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNoContent)
}

// Get202None204NoneDefaultError400Valid - Send a 400 response with valid payload: {'code': '400', 'message': 'client error'}
func TestGet202None204NoneDefaultError400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultError400Valid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("unexpected nil response")
	}
}

// Get202None204NoneDefaultNone202Invalid - Send a 202 response with an unexpected payload {'property': 'value'}
func TestGet202None204NoneDefaultNone202Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone202Invalid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusAccepted)
}

// Get202None204NoneDefaultNone204None - Send a 204 response with no payload
func TestGet202None204NoneDefaultNone204None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone204None(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusNoContent)
}

// Get202None204NoneDefaultNone400Invalid - Send a 400 response with an unexpected payload {'property': 'value'}
func TestGet202None204NoneDefaultNone400Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone400Invalid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("unexpected nil response")
	}
}

// Get202None204NoneDefaultNone400None - Send a 400 response with no payload
func TestGet202None204NoneDefaultNone400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.Get202None204NoneDefaultNone400None(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultModelA200None - Send a 200 response with no payload
func TestGetDefaultModelA200None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA200None(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if result.MyException != nil {
		t.Fatal("expected nil MyException")
	}
}

// GetDefaultModelA200Valid - Send a 200 response with valid payload: {'statusCode': '200'}
func TestGetDefaultModelA200Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA200Valid(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.MyException, &MyException{
		StatusCode: to.StringPtr("200"),
	})
}

// GetDefaultModelA400None - Send a 400 response with no payload
func TestGetDefaultModelA400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA400None(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultModelA400Valid - Send a 400 response with valid payload: {'statusCode': '400'}
func TestGetDefaultModelA400Valid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultModelA400Valid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultNone200Invalid - Send a 200 response with invalid payload: {'statusCode': '200'}
func TestGetDefaultNone200Invalid(t *testing.T) {
	t.Skip("payload doen't match schema doesn't cause unmarshalling error")
}

// GetDefaultNone200None - Send a 200 response with no payload
func TestGetDefaultNone200None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultNone200None(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// GetDefaultNone400Invalid - Send a 400 response with valid payload: {'statusCode': '400'}
func TestGetDefaultNone400Invalid(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultNone400Invalid(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("unexpected nil response")
	}
}

// GetDefaultNone400None - Send a 400 response with no payload
func TestGetDefaultNone400None(t *testing.T) {
	client := newMultipleResponsesClient()
	result, err := client.GetDefaultNone400None(context.Background())
	if err == nil {
		t.Fatal("unexpected nil error")
	}
	if result != nil {
		t.Fatal("unexpected nil response")
	}
}
