package pipeline

import (
	"bytes"
	"fmt"
	"net/http"
	"sort"
)

// The Response interface exposes an http.Response object as it returns through the pipeline of Policy objects.
// This ensures that Policy objects have access to the HTTP response. However, the object this interface encapsulates
// might be a struct with additional fields that is created by a Policy object (typically a method-specific Factory).
// The method that injected the method-specific Factory gets this returned Response and performs a type assertion
// to the expected struct and returns the struct to its caller.
type Response interface {
	Response() *http.Response
}

// This is the default struct that has the http.Response.
// A method can replace this struct with its own struct containing an http.Response
// field and any other additional fields.
type httpResponse struct {
	response *http.Response
}

// NewHTTPResponse is typically called by a Policy object to return a Response object.
func NewHTTPResponse(response *http.Response) Response {
	return &httpResponse{response: response}
}

// This method satisfies the public Response interface's Response method
func (r httpResponse) Response() *http.Response {
	return r.response
}

// WriteRequest appends a formatted HTTP request into a Buffer.
func WriteRequest(b *bytes.Buffer, request *http.Request) {
	fmt.Fprint(b, "=====> REQUEST: "+request.Method+" "+request.URL.String()+"\n")
	writeHeader(b, request.Header)
}

// WriteResponseWithRequest appends a formatted HTTP response with its initiating request into a Buffer.
func WriteResponseWithRequest(b *bytes.Buffer, response *http.Response) {
	WriteRequest(b, response.Request) // Write the request first followed by the response.
	fmt.Fprintf(b, "=====> RESPONSE: Status=%s\n", response.Status)
	writeHeader(b, response.Header)
}

// formatHeaders appends an HTTP request's or response's header into a Buffer.
func writeHeader(b *bytes.Buffer, header map[string][]string) {
	if len(header) == 0 {
		b.WriteString("   (no headers)\n")
		return
	}
	keys := make([]string, 0, len(header))
	// Alphabetize the headers
	for k := range header {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(b, "   %s: %+v\n", k, header[k])
	}
}
