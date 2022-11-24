package bestirlog

import (
	"strconv"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// See datadog docs: https://docs.datadoghq.com/tracing/connect_logs_and_traces/go/
const (
	// DatadogUITraceKey is the key that datadog expects for the trace id.
	DatadogUITraceKey = "dd.trace_id"
	// DatadogUISpanKey is the key that datadog expects for the span id.
	DatadogUISpanKey = "dd.span_id"
)

// DatadogTraceID is a generic function to extract the trace id
// from a span and return it as a string.
func DatadogTraceID(span tracer.Span) string {
	return strconv.FormatUint(span.Context().TraceID(), 10)
}

// DatadogSpanID is a generic function to extract the span id
// from a span and return it as a string.
func DatadogSpanID(span tracer.Span) string {
	return strconv.FormatUint(span.Context().SpanID(), 10)
}
