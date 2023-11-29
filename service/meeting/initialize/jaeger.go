package initialize

import (
	"fmt"
	"io"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/palp1tate/brevinect/service/meeting/global"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
	"go.uber.org/zap"
)

func NewJaegerTracer(serviceName string) (tracer opentracing.Tracer, closer io.Closer) {
	j := global.ServerConfig.Jaeger
	cfg := jaegercfg.Configuration{
		ServiceName: serviceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  fmt.Sprintf("%s:%d", j.Host, j.Port),
		},
	}
	tracer, closer, err := cfg.NewTracer(jaegercfg.Metrics(metrics.NullFactory))
	if err != nil {
		zap.S().Errorf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	opentracing.SetGlobalTracer(tracer)
	return
}
