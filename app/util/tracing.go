package util

import (
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"runtime"
	"time"
)

func Init(service string) (opentracing.Tracer, io.Closer) {
	defcfg := config.Configuration{
		ServiceName: service,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
	}

	cfg, err := defcfg.FromEnv()
	if err != nil {
		panic("Could not parse Jaeger env vars: " + err.Error())
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic("Could not initialize jaeger tracer: " + err.Error())
	}

	return tracer, closer
}

func StartRootSpan(ctx context.Context, name string) (context.Context, io.Closer, opentracing.Span) {
	tracer, closer := Init(name)
	sp := tracer.StartSpan(string("SERVICE"))
	ctx = opentracing.ContextWithSpan(ctx, sp)
	return ctx, closer, sp
}

func CreateRootSpan(ctx context.Context, name string) (opentracing.Span, opentracing.Tracer) {

	parentSpan := opentracing.SpanFromContext(ctx)
	tracer := parentSpan.Tracer()
	parentSpan.SetTag("name", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	parentSpan.SetTag("caller", callerDetails)

	return parentSpan, tracer
}

func CreateChildSpan(ctx context.Context, name string) opentracing.Span {
	parentSpan := opentracing.SpanFromContext(ctx)
	sp := opentracing.StartSpan(
		name,
		opentracing.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	return sp
}

func CreateSubChildSpan(parentSpan opentracing.Span, name string) opentracing.Span {
	sp := opentracing.StartSpan(
		name,
		opentracing.ChildOf(parentSpan.Context()))
	sp.SetTag("name", name)

	// Get caller function name, file and line
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	callerDetails := fmt.Sprintf("%s - %s#%d", frame.Function, frame.File, frame.Line)
	sp.SetTag("caller", callerDetails)

	return sp
}

func LogRequest(sp opentracing.Span, req interface{}) {
	sp.LogFields(log.Object(string("Request"), Stringify(req)))
}

func LogObject(sp opentracing.Span, name string, resp interface{}) {
	sp.LogFields(log.Object(name, Stringify(resp)))
}

func LogResponse(sp opentracing.Span, resp interface{}) {
	sp.LogFields(log.Object(string("Response"), Stringify(resp)))
}

func LogError(sp opentracing.Span, err error) {
	sp.LogFields(log.Object(string("Error"), err))
}
