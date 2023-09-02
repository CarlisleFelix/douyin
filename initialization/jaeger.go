package initialization

import (
	"douyin/config"
	"douyin/global"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func tracerProvider(jaegerCfg config.Jaeger) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerCfg.Endpoint)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(jaegerCfg.Servicename),
			attribute.String("environment", jaegerCfg.Environment),
			attribute.Int64("ID", jaegerCfg.Id),
		)),
	)
	return tp, nil
}

func InitializeJaeger() {
	jeagerCfg := global.SERVER_CONFIG.Jaeger
	tp, err := tracerProvider(jeagerCfg)
	if err != nil {
		log.Fatal(err)
	}
	global.SERVER_TRACE_PROVIDER = tp

	otel.SetTracerProvider(global.SERVER_TRACE_PROVIDER)

	userTracer := tp.Tracer("User Service")
	global.SERVER_USER_TRACER = userTracer

	videoTracer := tp.Tracer("Video Service")
	global.SERVER_VIDEO_TRACER = videoTracer
	relationTracer := tp.Tracer("Relation Service")
	global.SERVER_RELATION_TRACER = relationTracer
	commentTracer := tp.Tracer("Comment Service")
	global.SERVER_COMMENT_TRACER = commentTracer
	favoriteTracer := tp.Tracer("Favorite Service")
	global.SERVER_FAVORITE_TRACER = favoriteTracer
	messageTracer := tp.Tracer("Messge Service")
	global.SERVER_MESSAGE_TRACER = messageTracer

	//videoTracer.Start()
}
