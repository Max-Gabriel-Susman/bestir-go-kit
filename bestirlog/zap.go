package bestirlog

import (
	"context"

	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

// ZapLogger wraps a standard zap.Logger to automatically add trace and span ids to each log statement.
type ZapLogger struct{ *zap.Logger }

// WrapZap takes an existing zap.Logger and wraps it with the methods defined here.
func WrapZap(z *zap.Logger) *ZapLogger {
	return &ZapLogger{z}

}

// processContext adds the span and trace ids to the list of zap fields given,
// if the context actually contains a span. Otherwise it does nothing.
func (z *ZapLogger) processContext(ctx context.Context, fields []zap.Field) []zap.Field {
	if span, ok := tracer.SpanFromContext(ctx); ok {
		// Add datadog trace and span id's to the log statement.
		// Datadog requires trace and span id keys to be in the form of 'dd.[trace/span]_id' to corelate logs to traces (https://docs.datadoghq.com/tracing/connect_logs_and_traces/go/)
		// But Datadog does not show those fields in the log viewer. The 'visible.dd.[trace/span]_id' keys are to help debug and can/will be removed in the future.
		fields = append(fields,
			zap.String(DatadogUITraceKey, DatadogTraceID(span)),
			zap.String(DatadogUISpanKey, DatadogSpanID(span)),
			// For debugging
			zap.String("visible."+DatadogUITraceKey, DatadogTraceID(span)),
			zap.String("visible."+DatadogUISpanKey, DatadogSpanID(span)),
		)
	}
	return fields
}

// With calls the underlying zap With function but returns a ZapLogger instead of a zap.Logger
func (z *ZapLogger) With(fields ...zap.Field) *ZapLogger {
	return WrapZap(z.Logger.With(fields...))
}

func (z *ZapLogger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	z.Logger.Info(msg, z.processContext(ctx, fields)...)
}

func (z *ZapLogger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	z.Logger.Error(msg, z.processContext(ctx, fields)...)
}

func (z *ZapLogger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	z.Logger.Warn(msg, z.processContext(ctx, fields)...)
}

func (z *ZapLogger) Debug(ctx context.Context, msg string, fields ...zap.Field) {
	z.Logger.Debug(msg, z.processContext(ctx, fields)...)
}
