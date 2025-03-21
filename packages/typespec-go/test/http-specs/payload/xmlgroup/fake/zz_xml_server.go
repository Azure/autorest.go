// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
	"sync"
)

// XMLServer is a fake server for instances of the xmlgroup.XMLClient type.
type XMLServer struct {
	// XMLModelWithArrayOfModelValueServer contains the fakes for client XMLModelWithArrayOfModelValueClient
	XMLModelWithArrayOfModelValueServer XMLModelWithArrayOfModelValueServer

	// XMLModelWithAttributesValueServer contains the fakes for client XMLModelWithAttributesValueClient
	XMLModelWithAttributesValueServer XMLModelWithAttributesValueServer

	// XMLModelWithDictionaryValueServer contains the fakes for client XMLModelWithDictionaryValueClient
	XMLModelWithDictionaryValueServer XMLModelWithDictionaryValueServer

	// XMLModelWithEmptyArrayValueServer contains the fakes for client XMLModelWithEmptyArrayValueClient
	XMLModelWithEmptyArrayValueServer XMLModelWithEmptyArrayValueServer

	// XMLModelWithEncodedNamesValueServer contains the fakes for client XMLModelWithEncodedNamesValueClient
	XMLModelWithEncodedNamesValueServer XMLModelWithEncodedNamesValueServer

	// XMLModelWithOptionalFieldValueServer contains the fakes for client XMLModelWithOptionalFieldValueClient
	XMLModelWithOptionalFieldValueServer XMLModelWithOptionalFieldValueServer

	// XMLModelWithRenamedArraysValueServer contains the fakes for client XMLModelWithRenamedArraysValueClient
	XMLModelWithRenamedArraysValueServer XMLModelWithRenamedArraysValueServer

	// XMLModelWithRenamedFieldsValueServer contains the fakes for client XMLModelWithRenamedFieldsValueClient
	XMLModelWithRenamedFieldsValueServer XMLModelWithRenamedFieldsValueServer

	// XMLModelWithSimpleArraysValueServer contains the fakes for client XMLModelWithSimpleArraysValueClient
	XMLModelWithSimpleArraysValueServer XMLModelWithSimpleArraysValueServer

	// XMLModelWithTextValueServer contains the fakes for client XMLModelWithTextValueClient
	XMLModelWithTextValueServer XMLModelWithTextValueServer

	// XMLModelWithUnwrappedArrayValueServer contains the fakes for client XMLModelWithUnwrappedArrayValueClient
	XMLModelWithUnwrappedArrayValueServer XMLModelWithUnwrappedArrayValueServer

	// XMLSimpleModelValueServer contains the fakes for client XMLSimpleModelValueClient
	XMLSimpleModelValueServer XMLSimpleModelValueServer
}

// NewXMLServerTransport creates a new instance of XMLServerTransport with the provided implementation.
// The returned XMLServerTransport instance is connected to an instance of xmlgroup.XMLClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewXMLServerTransport(srv *XMLServer) *XMLServerTransport {
	return &XMLServerTransport{srv: srv}
}

// XMLServerTransport connects instances of xmlgroup.XMLClient to instances of XMLServer.
// Don't use this type directly, use NewXMLServerTransport instead.
type XMLServerTransport struct {
	srv                                     *XMLServer
	trMu                                    sync.Mutex
	trXMLModelWithArrayOfModelValueServer   *XMLModelWithArrayOfModelValueServerTransport
	trXMLModelWithAttributesValueServer     *XMLModelWithAttributesValueServerTransport
	trXMLModelWithDictionaryValueServer     *XMLModelWithDictionaryValueServerTransport
	trXMLModelWithEmptyArrayValueServer     *XMLModelWithEmptyArrayValueServerTransport
	trXMLModelWithEncodedNamesValueServer   *XMLModelWithEncodedNamesValueServerTransport
	trXMLModelWithOptionalFieldValueServer  *XMLModelWithOptionalFieldValueServerTransport
	trXMLModelWithRenamedArraysValueServer  *XMLModelWithRenamedArraysValueServerTransport
	trXMLModelWithRenamedFieldsValueServer  *XMLModelWithRenamedFieldsValueServerTransport
	trXMLModelWithSimpleArraysValueServer   *XMLModelWithSimpleArraysValueServerTransport
	trXMLModelWithTextValueServer           *XMLModelWithTextValueServerTransport
	trXMLModelWithUnwrappedArrayValueServer *XMLModelWithUnwrappedArrayValueServerTransport
	trXMLSimpleModelValueServer             *XMLSimpleModelValueServerTransport
}

// Do implements the policy.Transporter interface for XMLServerTransport.
func (x *XMLServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	return x.dispatchToClientFake(req, method[:strings.Index(method, ".")])
}

func (x *XMLServerTransport) dispatchToClientFake(req *http.Request, client string) (*http.Response, error) {
	var resp *http.Response
	var err error

	switch client {
	case "XMLModelWithArrayOfModelValueClient":
		initServer(&x.trMu, &x.trXMLModelWithArrayOfModelValueServer, func() *XMLModelWithArrayOfModelValueServerTransport {
			return NewXMLModelWithArrayOfModelValueServerTransport(&x.srv.XMLModelWithArrayOfModelValueServer)
		})
		resp, err = x.trXMLModelWithArrayOfModelValueServer.Do(req)
	case "XMLModelWithAttributesValueClient":
		initServer(&x.trMu, &x.trXMLModelWithAttributesValueServer, func() *XMLModelWithAttributesValueServerTransport {
			return NewXMLModelWithAttributesValueServerTransport(&x.srv.XMLModelWithAttributesValueServer)
		})
		resp, err = x.trXMLModelWithAttributesValueServer.Do(req)
	case "XMLModelWithDictionaryValueClient":
		initServer(&x.trMu, &x.trXMLModelWithDictionaryValueServer, func() *XMLModelWithDictionaryValueServerTransport {
			return NewXMLModelWithDictionaryValueServerTransport(&x.srv.XMLModelWithDictionaryValueServer)
		})
		resp, err = x.trXMLModelWithDictionaryValueServer.Do(req)
	case "XMLModelWithEmptyArrayValueClient":
		initServer(&x.trMu, &x.trXMLModelWithEmptyArrayValueServer, func() *XMLModelWithEmptyArrayValueServerTransport {
			return NewXMLModelWithEmptyArrayValueServerTransport(&x.srv.XMLModelWithEmptyArrayValueServer)
		})
		resp, err = x.trXMLModelWithEmptyArrayValueServer.Do(req)
	case "XMLModelWithEncodedNamesValueClient":
		initServer(&x.trMu, &x.trXMLModelWithEncodedNamesValueServer, func() *XMLModelWithEncodedNamesValueServerTransport {
			return NewXMLModelWithEncodedNamesValueServerTransport(&x.srv.XMLModelWithEncodedNamesValueServer)
		})
		resp, err = x.trXMLModelWithEncodedNamesValueServer.Do(req)
	case "XMLModelWithOptionalFieldValueClient":
		initServer(&x.trMu, &x.trXMLModelWithOptionalFieldValueServer, func() *XMLModelWithOptionalFieldValueServerTransport {
			return NewXMLModelWithOptionalFieldValueServerTransport(&x.srv.XMLModelWithOptionalFieldValueServer)
		})
		resp, err = x.trXMLModelWithOptionalFieldValueServer.Do(req)
	case "XMLModelWithRenamedArraysValueClient":
		initServer(&x.trMu, &x.trXMLModelWithRenamedArraysValueServer, func() *XMLModelWithRenamedArraysValueServerTransport {
			return NewXMLModelWithRenamedArraysValueServerTransport(&x.srv.XMLModelWithRenamedArraysValueServer)
		})
		resp, err = x.trXMLModelWithRenamedArraysValueServer.Do(req)
	case "XMLModelWithRenamedFieldsValueClient":
		initServer(&x.trMu, &x.trXMLModelWithRenamedFieldsValueServer, func() *XMLModelWithRenamedFieldsValueServerTransport {
			return NewXMLModelWithRenamedFieldsValueServerTransport(&x.srv.XMLModelWithRenamedFieldsValueServer)
		})
		resp, err = x.trXMLModelWithRenamedFieldsValueServer.Do(req)
	case "XMLModelWithSimpleArraysValueClient":
		initServer(&x.trMu, &x.trXMLModelWithSimpleArraysValueServer, func() *XMLModelWithSimpleArraysValueServerTransport {
			return NewXMLModelWithSimpleArraysValueServerTransport(&x.srv.XMLModelWithSimpleArraysValueServer)
		})
		resp, err = x.trXMLModelWithSimpleArraysValueServer.Do(req)
	case "XMLModelWithTextValueClient":
		initServer(&x.trMu, &x.trXMLModelWithTextValueServer, func() *XMLModelWithTextValueServerTransport {
			return NewXMLModelWithTextValueServerTransport(&x.srv.XMLModelWithTextValueServer)
		})
		resp, err = x.trXMLModelWithTextValueServer.Do(req)
	case "XMLModelWithUnwrappedArrayValueClient":
		initServer(&x.trMu, &x.trXMLModelWithUnwrappedArrayValueServer, func() *XMLModelWithUnwrappedArrayValueServerTransport {
			return NewXMLModelWithUnwrappedArrayValueServerTransport(&x.srv.XMLModelWithUnwrappedArrayValueServer)
		})
		resp, err = x.trXMLModelWithUnwrappedArrayValueServer.Do(req)
	case "XMLSimpleModelValueClient":
		initServer(&x.trMu, &x.trXMLSimpleModelValueServer, func() *XMLSimpleModelValueServerTransport {
			return NewXMLSimpleModelValueServerTransport(&x.srv.XMLSimpleModelValueServer)
		})
		resp, err = x.trXMLSimpleModelValueServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	return resp, err
}

// set this to conditionally intercept incoming requests to XMLServerTransport
var xmlServerTransportInterceptor interface {
	// Do returns true if the server transport should use the returned response/error
	Do(*http.Request) (*http.Response, error, bool)
}
