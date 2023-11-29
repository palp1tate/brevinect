package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/palp1tate/brevinect/api/global"
	"github.com/palp1tate/brevinect/api/initialize"
	"github.com/palp1tate/brevinect/util"
	"go.uber.org/zap"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	router, closer := initialize.Router()
	if err := initialize.InitTrans("zh"); err != nil {
		zap.S().Warn("初始化翻译器失败:", err.Error())
		panic(err)
	}
	initialize.InitValidator()
	initialize.InitServiceConn()

	host := global.ServerConfig.Api.Host
	port := flag.Int("p", 0, "端口号")
	flag.Parse()
	if *port == 0 {
		*port, _ = util.GetFreePort()
	}
	zap.S().Info("host: ", host)
	zap.S().Info("port: ", *port)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			zap.S().Panic("网关启动失败:", err.Error())
		}
	}()

	client := initialize.NewRegistryClient(global.ServerConfig.Consul.Host, global.ServerConfig.Consul.Port)
	apiName := global.ServerConfig.Api.Name
	apiTags := global.ServerConfig.Api.Tags
	apiId := util.GenerateUUID()
	err := client.Register(host, *port, apiName, apiTags, apiId)
	if err != nil {
		zap.S().Panic("服务注册失败:", err.Error())
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	defer closer.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		zap.S().Error("Server forced to shutdown:", err)
	}
	if err = client.DeRegister(apiId); err != nil {
		zap.S().Warnf("%s注销失败", apiName)
	} else {
		zap.S().Infof("%s注销成功", apiName)
	}
}
