package initialize

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/api/handler"
	"github.com/palp1tate/brevinect/api/middleware"
	"github.com/palp1tate/brevinect/api/router"
	sentinel "github.com/sentinel-group/sentinel-go-adapters/gin"
)

func Router() (*gin.Engine, io.Closer) {
	if !global.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	serviceName := global.ServerConfig.Api.Name
	tracer, closer := NewJaegerTracer(serviceName)
	opentracing.SetGlobalTracer(tracer)
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		handler.HandleHttpResponse(c, http.StatusOK, "ok", nil, nil)
		return
	})
	r.Use(
		middleware.Cors(),
		middleware.OpenTracing(),
		sentinel.SentinelMiddleware(sentinel.WithResourceExtractor(middleware.ResourceExtractor)),
	)
	ApiGroup := r.Group("/api")
	router.InitThirdPartyRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	router.InitAdminRouter(ApiGroup)
	router.InitMeetingRouter(ApiGroup)
	return r, closer
}
