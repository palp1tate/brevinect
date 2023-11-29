package initialize

import (
	"github.com/importcjj/sensitive"
	"github.com/palp1tate/brevinect/api/global"
	"go.uber.org/zap"
)

func InitFilter() {
	filter := sensitive.New()
	err := filter.LoadWordDict("key.txt")
	if err != nil {
		zap.S().Panic("加载敏感词词库失败:", err.Error())
	}
	global.Filter = filter
}
