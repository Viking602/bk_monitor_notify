package svc

import (
	"bk_monitor_notify/common/dingtalk"
	"bk_monitor_notify/common/lark"
	"bk_monitor_notify/internal/config"
)

type ServiceContext struct {
	Config           config.Config
	LarkBotModel     lark.BotLark
	DingTalkBotModel dingtalk.BotDingTalk
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		LarkBotModel:     lark.NewLarkBot(),
		DingTalkBotModel: dingtalk.NewDingTalkBot(),
	}
}
