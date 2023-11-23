package initialize

import (
	"github.com/palp1tate/brevinect/service/meeting/global"
	"go.uber.org/zap"
)

func InitLogger() {
	logger, _ := zap.NewProduction()
	if global.Debug {
		logger, _ = zap.NewDevelopment()
	}
	zap.ReplaceGlobals(logger)
}
