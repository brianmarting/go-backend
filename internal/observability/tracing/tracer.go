package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	name = "go-backend"
)

func GetTracer() trace.Tracer {
	return otel.GetTracerProvider().Tracer(name)
}
