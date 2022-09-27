package logic

import (
	"bk_monitor_notify/common/dingtalk"
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

type DingTalkBotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDingTalkBotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DingTalkBotLogic {
	return &DingTalkBotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DingTalkBotLogic) DingTalkBot(req *types.BKCallBack) error {
	toString, err := jsonx.MarshalToString(req)
	if err != nil {
		return err
	}
	l.Infof(toString)

	var anomalyMsg string

	var tempColor string

	// 判断异常消息
	if req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field1.AnomalyMessage != "" {
		// 告警等级为1=致命则为红色
		tempColor = "FF0000"
		anomalyMsg = req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field1.AnomalyMessage
	} else if req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field2.AnomalyMessage != "" {
		// 告警等级为2=预警则为橙色
		tempColor = "FF8000"
		anomalyMsg = req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field2.AnomalyMessage
	} else if req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field3.AnomalyMessage != "" {
		// 告警等级为3=提醒则为黄色
		tempColor = "FCDB03"
		anomalyMsg = req.LatestAnomalyRecord.OriginAlarm.Anomaly.Field3.AnomalyMessage
	}

	var title string
	if req.Type == "RECOVERY_NOTICE" {
		// 告警恢复则标题为绿色
		tempColor = "00FF2B"
		title = fmt.Sprintf("<font color=\"#%s\">告警恢复 - %s业务</font>", tempColor, req.BkBizName)
	} else {
		// 否则按异常消息等级分配颜色并且命名标题
		title = fmt.Sprintf("<font color=\"#%s\">%d级告警 - %s业务</font>", tempColor, req.Event.Level, req.BkBizName)
	}

	var bkTopNodes []string

	for _, displayValue := range req.Event.DimensionTranslation.BktopoNode.DisplayValue {
		instName := displayValue.InstName
		objName := displayValue.ObjName
		bkTopNode := fmt.Sprintf(" >**%s:** %s\n", objName, instName)
		bkTopNodes = append(bkTopNodes, bkTopNode)
	}

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

	BkUrl := l.svcCtx.Config.BKUrl

	eventCenterUrl := fmt.Sprintf("%s/o/bk_monitorv3/?bizId=%d#/event-center/detail/%d", BkUrl, req.BkBizId, req.Event.Id)

	msg := fmt.Sprintf("## %s \n\n**告警级别:** <font color=\"#%s\"> %s </font> \n\n**告警策略:** %s \n\n>**%s:** %s \n\n >**%s:** %s \n\n>**异常事件:** %s \n\n>**开始时间:** %s \n\n >**告警时间:** %s \n\n >**结束时间:** %s \n\n >**事件ID:** %d \n\n %s \n\n[查看详情](%s)", title, tempColor, req.Event.LevelName, req.Strategy.Name, req.Event.DimensionTranslation.BkTargetCloudId.DisplayName, req.Event.DimensionTranslation.BkTargetCloudId.Value, req.Event.DimensionTranslation.BkTargetIp.DisplayName, req.Event.DimensionTranslation.BkTargetIp.Value, anomalyMsg, createTime, beginTime, endTime, req.Event.Id, strings.Join(bkTopNodes, "\n"), eventCenterUrl)
	timestamp := time.Now().UnixMilli()
	sign := utils.DingTalkSign(req.Secret, timestamp)
	dingTalkData := fmt.Sprintf("%s&timestamp=%d&sign=%s", req.BotKey, timestamp, sign)
	data := &dingtalk.MarkDownDingTalk{
		Msgtype: "markdown",
		Markdown: dingtalk.MarkDownDing{
			Title: title,
			Text:  msg,
		},
	}

	botResp, err := l.svcCtx.DingTalkBotModel.Bot(dingTalkData, data)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if botResp.Errcode != 0 {
		return errors.New(botResp.Errmsg)
	}
	return nil
}
