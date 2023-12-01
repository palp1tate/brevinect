package initialize

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/hotspot"
	"github.com/palp1tate/brevinect/consts"
	"go.uber.org/zap"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		zap.S().Fatalf("初始化sentinel 异常: %v", err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               consts.UserResource,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              3000,
			StatIntervalInMs:       1000,
		},
		{
			Resource:               consts.AdminResource,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              4000,
			StatIntervalInMs:       1000,
		},
		{
			Resource:               consts.MeetingResource,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              5000,
			StatIntervalInMs:       1000,
		},
		{
			Resource:               consts.ThirdResource,
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject,
			Threshold:              2000,
			StatIntervalInMs:       1000,
		},
	})

	if err != nil {
		zap.S().Fatalf("加载流控规则失败: %v", err)
	}
	_, err = hotspot.LoadRules([]*hotspot.Rule{
		{
			Resource:        consts.SMSResource,
			MetricType:      hotspot.QPS,
			ControlBehavior: hotspot.Reject,
			ParamIndex:      0, // 根据手机号限流
			Threshold:       1,
			DurationInSec:   60,
		},
		{
			Resource:        consts.BookResource,
			MetricType:      hotspot.QPS,
			ControlBehavior: hotspot.Reject,
			ParamIndex:      0,
			Threshold:       3,
			DurationInSec:   300,
		},
	})

	if err != nil {
		zap.S().Fatalf("加载热点规则失败: %v", err)
	}
}
