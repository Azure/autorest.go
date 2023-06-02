// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package generatortests

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func NewTracingProvider(t *testing.T) tracing.Provider {
	return tracing.NewProvider(func(name, version string) tracing.Tracer {
		tt := testTracer{
			name:    name,
			version: version,
		}

		t.Cleanup(func() {
			tt.Verify(t)
		})

		return tracing.NewTracer(func(ctx context.Context, spanName string, options *tracing.SpanOptions) (context.Context, tracing.Span) {
			kind := tracing.SpanKindInternal
			if options != nil {
				kind = options.Kind
			}
			return tt.Start(ctx, spanName, kind, t)
		}, nil)
	}, nil)
}

type testTracer struct {
	name    string
	version string
	spans   []*testSpan
}

func (tt *testTracer) Start(ctx context.Context, spanName string, kind tracing.SpanKind, t *testing.T) (context.Context, tracing.Span) {
	ts := testSpan{
		name: spanName,
	}

	tt.spans = append(tt.spans, &ts)
	return ctx, tracing.NewSpan(tracing.SpanImpl{
		End: ts.End,
	})
}

func (tt *testTracer) Verify(t *testing.T) {
	assert.NotEmpty(t, tt.name)
	assert.NotEmpty(t, tt.version)
	require.NotEmpty(t, tt.spans)
	for _, span := range tt.spans {
		span.Verify(t)
	}
}

type testSpan struct {
	name  string
	ended bool
}

func (ts *testSpan) End() {
	ts.ended = true
}

func (ts *testSpan) Verify(t *testing.T) {
	assert.NotEmpty(t, ts.name)
	assert.True(t, ts.ended)
}
