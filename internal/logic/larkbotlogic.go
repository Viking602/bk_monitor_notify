package logic

import (
	"bk_monitor_notify/common/lark"
	"bk_monitor_notify/common/utils"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"strings"
	"time"

	"bk_monitor_notify/internal/svc"
	"bk_monitor_notify/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeiShuBotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLarkBotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeiShuBotLogic {
	return &FeiShuBotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeiShuBotLogic) LarkBot(req *types.BKCallBack) error {
	//JSON格式化请求参数
	toString, err := jsonx.MarshalToString(req)
	if err != nil {
		return err
	}
	// 记录请求参数
	l.Infof(toString)
	//获取时间戳
	timestamp := time.Now().Unix()
	// Template颜色
	var tempColor string
	// 异常消息
	var anomalyMsg string
	// 判断异常消息
	if req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field1.AnomalyMessage != "" {
		// 告警等级为1=致命则为红色
		tempColor = "red"
		anomalyMsg = req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field1.AnomalyMessage
	} else if req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field2.AnomalyMessage != "" {
		// 告警等级为2=预警则为橙色
		tempColor = "orange"
		anomalyMsg = req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field2.AnomalyMessage
	} else if req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field3.AnomalyMessage != "" {
		// 告警等级为3=提醒则为黄色
		tempColor = "yellow"
		anomalyMsg = req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field3.AnomalyMessage
	}
	// 定义标题
	var title string
	if req.Type == "RECOVERY_NOTICE" {
		// 告警恢复则标题为绿色
		tempColor = "green"
		title = fmt.Sprintf("告警恢复 - %s业务", req.BkBizName)
	} else {
		// 否则按异常消息等级分配颜色并且命名标题
		title = fmt.Sprintf("%d级告警 - %s业务", req.Event.Level, req.BkBizName)
	}
	elements := make([]interface{}, 0)

	elements1 := &lark.CardElements{
		Fields: []*lark.CardFields{
			{
				IsShort: true,
				Text: &lark.CardText{
					Content: "**🚨告警级别：**\n" + req.Event.LevelName,
					Tag:     "lark_md",
				},
			},
			{
				IsShort: true,
				Text: &lark.CardText{
					Content: "**🗒告警策略：**\n" + req.Strategy.Name,
					Tag:     "lark_md",
				},
			},
		},
		Tag: "div",
	}

	elements = append(elements, elements1)

	elements2 := &lark.CardElements{
		Tag: "hr",
	}
	elements = append(elements, elements2)

	var bkTopNodes []string

	for _, displayValue := range req.Event.DimensionTranslation.BktopoNode.DisplayValue {
		instName := displayValue.InstName
		objName := displayValue.ObjName
		bkTopNode := fmt.Sprintf("**%s:**%s\n", objName, instName)
		bkTopNodes = append(bkTopNodes, bkTopNode)
	}

	BkUrl := l.svcCtx.Config.BKUrl

	eventCenterUrl := fmt.Sprintf("%s/o/bk_monitorv3/?bizId=%d#/event-center/detail/%d", BkUrl, req.BkBizId, req.Event.Id)

	var createTime string
	var beginTime string
	var endTime string

	if req.TimeLocal == 1 {
		// 转换时区为CST
		createTime = time.Unix(utils.TimeToTimestampAdd8(req.Event.CreateTime), 0).In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")
		beginTime = time.Unix(utils.TimeToTimestampAdd8(req.Event.BeginTime), 0).In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")
		if req.Event.EndTime != "1980-01-01 08:00:00" {
			endTime = time.Unix(utils.TimeToTimestampAdd8(req.Event.EndTime), 0).In(time.FixedZone("CST", 8*3600)).Format("2006-01-02 15:04:05")
		} else {
			endTime = "暂无"
		}

	} else {
		createTime = req.Event.CreateTime
		beginTime = req.Event.BeginTime
		if req.Event.EndTime != "1980-01-01 08:00:00" {
			endTime = req.Event.EndTime
		} else {
			endTime = "暂无"
		}
	}

	eventData := fmt.Sprintf("**开始时间:**%s \n**告警时间:**%s \n**结束时间:**%s \n**事件ID:**%d", beginTime, createTime, endTime, req.Event.Id)

	msg := fmt.Sprintf("**%s:**%s \n**%s:**%s \n**异常事件:**%s\n%s\n%s\n**详细情况点击[事件中心](%s)**", req.Event.DimensionTranslation.BkTargetCloudId.DisplayName, req.Event.DimensionTranslation.BkTargetCloudId.DisplayValue, req.Event.DimensionTranslation.BkTargetIp.DisplayName, req.Event.DimensionTranslation.BkTargetIp.Value, anomalyMsg, eventData, strings.Join(bkTopNodes, ""), eventCenterUrl)

	elements3 := &lark.CardElements{
		Tag:     "markdown",
		Content: msg,
	}
	elements = append(elements, elements3)

	cardContent := lark.Card{
		Config:   lark.CardConfig{WideScreenMode: true},
		Elements: elements,
		Header: lark.CardHeader{Template: tempColor, Title: &lark.CardText{
			Content: title,
			Tag:     "plain_text",
		}},
	}
	// 计算飞书机器人签名
	sign, err := utils.GenSign(req.Secret, timestamp)
	if err != nil {
		l.Errorw("生成签名错误", logx.Field("error", err))
		return err
	}

	l2 := &lark.BotCard{
		Card:      cardContent,
		MsgType:   "interactive",
		Sign:      sign,
		Timestamp: timestamp,
	}
	botResp, err := l.svcCtx.LarkBotModel.Bot(req.BotKey, l2)
	if err != nil {
		return err
	}
	if botResp.Code != 0 {
		return errors.New(botResp.Msg)
	}
	return nil
}
