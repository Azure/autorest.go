package pipeline

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
)

// The Factory interface represents an object that can create its Policy object. Each HTTP request sent
// requires that this Factory create a new instance of its Policy object.
type Factory interface {
	New(node Node) Policy
}

// The Policy interface represents a mutable Policy object created by a Factory. The object can mutate/process
// the HTTP request and then forward it on to the next Policy object in the linked-list. The returned
// Response goes backward through the linked-list for additional processing.
// NOTE: Request is passed by value so changes do not change the caller's version of
// the request. However, Request has some fields that reference mutable objects (not strings).
// These references are copied; a deep copy is not performed. Specifically, this means that
// you should avoid modifying the objects referred to by these fields: URL, Header, Body,
// GetBody, TransferEncoding, Form, MultipartForm, Trailer, TLS, Cancel, and Response.
type Policy interface {
	Do(ctx context.Context, request Request) (Response, error)
}

// Options configures a Pipeline's behavior.
type Options struct {
	HTTPSender Factory // If sender is nil, then http.DefaultClient is used to send the HTTP requests.
	Log        LogOptions
}

// LogSeverity tells a logger the minimum severity to log. When code reports a log entry,
// the LogSeverity indicates the severity of the log entry. The logger only records entries
// whose severity is at as as severe as what it was told to log. See the Log* constants.
// For example, if a logger is configured with LogError severity, then LogError, LogPanic,
// and LogFatal entries will be logged; less severe entries will be ignored.
type LogSeverity int

const (
	// LogNone tells a logger not to log any entries passed to it.
	LogNone LogSeverity = iota

	// LogFatal tells a logger to log all LogFatal entries passed to it.
	LogFatal

	// LogPanic tells a logger to log all LogPanic and LogFatal entries passed to it.
	LogPanic

	// LogError tells a logger to log all LogError, LogPanic and LogFatal entries passed to it.
	LogError

	// LogWarning tells a logger to log all LogWarning, LogError, LogPanic and LogFatal entries passed to it.
	LogWarning

	// LogInfo tells a logger to log all LogInfo, LogWarning, LogError, LogPanic and LogFatal entries passed to it.
	LogInfo
)

// LogOptions configures the pipeline's logging mechanism & severity filtering.
type LogOptions struct {
	Log            func(level LogSeverity, message string)
	LogMaxSeverity LogSeverity // Defaults to LogNone
}

type pipeline struct {
	factories []Factory
	options   Options
}

// The Pipeline interface represents an ordered list of Factory objects and an object implementing the HTTPSender interface.
// You construct a Pipeline by calling the pipeline.NewPipeline function. To send an HTTP request, call pipeline.NewRequest
// and then call Pipeline's Do method passing a context, the request, and a method-specific Factory (or nil). Passing a
// method-specific Factory allows this one call to Do to inject a Policy into the linked-list. The policy is injected where
// the MethodFactoryMarker (see the pipeline.MethodFactoryMarker function) is in the slice of Factory objects.
//
// When Do is called, the Pipeline object asks each Factory object to construct its Policy object and adds each Policy to a linked-list.
// THen, Do sends the Context and Request through all the Policy objects. The final Policy object sends the request over the network
// (via the HTTPSender object passed to NewPipeline) and the response is returned backwards through all the Policy objects.
// Since Pipeline and Factory objects are goroutine-safe, you typically create 1 Pipeline object and reuse it to make many HTTP requests.
type Pipeline interface {
	Do(ctx context.Context, methodFactory Factory, request Request) (Response, error)
}

// NewPipeline creates a new goroutine-safe Pipeline object from the slice of Factory objects and the specified options.
func NewPipeline(factories []Factory, o Options) Pipeline {
	if o.HTTPSender == nil {
		o.HTTPSender = newDefaultHTTPClientFactory()
	}
	if o.Log.Log == nil {
		o.Log.Log = func(LogSeverity, string) {} // No-op logger
	}
	return &pipeline{factories: factories, options: o}
}

// Do is called for each and every HTTP request. It tells each Factory to create its own (mutable) Policy object
// replacing a MethodFactoryMarker factory (if it exists) with the methodFactory passed in. Then, the Context and Request
// are sent through the pipeline of Policy objects (which can transform the Request's URL/query parameters/headers) and
// ultimately sends the transformed HTTP request over the network.
func (p *pipeline) Do(ctx context.Context, methodFactory Factory, request Request) (Response, error) {
	response, err := p.newPolicies(methodFactory).Do(ctx, request)
	request.close()
	return response, err
}

func (p *pipeline) newPolicies(methodFactory Factory) Policy {
	// The last Policy is the one that actually sends the request over the wire and gets the response.
	// It is overridable via the Options' HTTPSender field.
	node := Node{pipeline: p, next: nil}
	node.next = p.options.HTTPSender.New(node)

	// Walk over the slice of Factory objects
	markers := 0
	for _, factory := range p.factories {
		if _, ok := factory.(methodFactoryMarker); ok {
			markers++
			if markers > 1 {
				panic("MethodFactoryMarker can only appear once in the pipeline")
			}
			if methodFactory != nil {
				// Replace MethodFactoryMarker with passed-in methodFactory
				node.next = methodFactory.New(node)
			}
		} else {
			// Use the slice's Factory to construct its Policy
			node.next = factory.New(node)
		}
	}
	// Each Factory has created its Policy
	if markers == 0 && methodFactory != nil {
		panic("Non-nil methodFactory requires MethodFactoryMarker in the pipeline")
	}
	return node.next // Return head of the Policy object linked-list
}

// A Node represents a node in a linked-list of Policy objects. A Node is passed
// to the Factory's New method which passes to the Policy object it creates. The Policy object
// uses the Node to forward the Context and HTTP request to the next Policy object in the pipeline.
type Node struct {
	pipeline *pipeline
	next     Policy
}

// Do forwards the Context and HTTP request to the next Policy object in the pipeline. The last Policy object
// sends the request over the network via HTTPSender's Do method. The response and error are returned
// back up the pipeline through the Policy objects.
func (n *Node) Do(ctx context.Context, request Request) (Response, error) {
	return n.next.Do(ctx, request)
}

// WouldLog returns true if the specified severity level would be logged.
func (n *Node) WouldLog(severity LogSeverity) bool {
	return severity <= n.pipeline.options.Log.LogMaxSeverity
}

// Log logs a string to the Pipeline's Logger.
func (n *Node) Log(severity LogSeverity, msg string) {
	if !n.WouldLog(severity) {
		return // Short circuit message formatting if we're not logging it
	}
	if len(msg) == 0 || msg[len(msg)-1] != '\n' {
		msg += "\n" // Ensure trailing newline
	}
	defaultLog(severity, msg)
	n.pipeline.options.Log.Log(severity, msg)
	// If logger doesn't handle fatal/panic, we'll do it here.
	if severity == LogFatal {
		os.Exit(1)
	} else if severity == LogPanic {
		panic(msg)
	}
}

// Logf logs a string to the Pipeline's Logger.
func (n *Node) Logf(severity LogSeverity, format string, v ...interface{}) {
	if !n.WouldLog(severity) {
		return // Short circuit message formatting if we're not logging it
	}
	b := &bytes.Buffer{}
	fmt.Fprintf(b, format, v...)
	if b.Len() == 0 || b.Bytes()[b.Len()-1] != '\n' {
		b.WriteRune('\n') // Ensure trailing newline
	}
	n.Log(severity, b.String())
}

// newDefaultHTTPClientFactory creates a DefaultHTTPClientPolicyFactory object that sends HTTP requests to a Go's default http.Client.
func newDefaultHTTPClientFactory() Factory {
	return &defaultHTTPClientPolicyFactory{}
}

type defaultHTTPClientPolicyFactory struct {
}

// Create initializes a logging policy object.
func (f *defaultHTTPClientPolicyFactory) New(node Node) Policy {
	return &defaultHTTPClientPolicy{node: node}
}

type defaultHTTPClientPolicy struct {
	node Node
}

func (p *defaultHTTPClientPolicy) Do(ctx context.Context, request Request) (Response, error) {
	r, err := http.DefaultClient.Do(request.WithContext(ctx))
	if err != nil {
		err = NewError(err, "HTTP request failed")
	}
	return NewHTTPResponse(r), err
}

var mfm = methodFactoryMarker{}

// MethodFactoryMarker returns a special marker Factory object. When Pipeline's Do method is called, any
// MethodMarkerFactory object is replaced with the specified methodFactory object. If nil is passed fro Do's
// methodFactory parameter, then the MethodFactoryMarker is ignored as the linked-list of Policy objects is created.
func MethodFactoryMarker() Factory {
	return mfm
}

type methodFactoryMarker struct {
}

func (mpmf methodFactoryMarker) New(node Node) Policy {
	panic("methodFactoryMarker policy should have been replaced with a method policy")
}
