package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/api/handler"
	"github.com/palp1tate/brevinect/api/middleware"
	"github.com/palp1tate/brevinect/api/router"
)

func Router() *gin.Engine {
	if !global.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		handler.HandleHttpResponse(c, http.StatusOK, "ok", nil, nil)
		return
	})
	r.Use(middleware.Cors())
	ApiGroup := r.Group("/api")
	router.InitThirdPartyRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	router.InitAdminRouter(ApiGroup)
	return r
}
