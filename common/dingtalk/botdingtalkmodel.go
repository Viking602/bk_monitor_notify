package dingtalk

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"net/http"
)

type (
	BotDingTalk interface {
		Bot(path string, data *MarkDownDingTalk) (*BotResponse, error)
	}
	dingTalkDefaultModel struct {
		*defaultModel
	}

	MarkDownDingTalk struct {
		Msgtype  string       `json:"msgtype"`
		Markdown MarkDownDing `json:"markdown"`
		At       AtDing       `json:"at,omitempty"`
	}
	MarkDownDing struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}
	AtDing struct {
		AtMobiles []string `json:"atMobiles"`
		AtUserIds []string `json:"atUserIds"`
		IsAtAll   bool     `json:"isAtAll"`
	}

	BotResponse struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
	}
)

func NewDingTalkBot() BotDingTalk {
	return &dingTalkDefaultModel{
		defaultModel: newDefaultModel(),
	}
}

func (l dingTalkDefaultModel) Bot(path string, data *MarkDownDingTalk) (*BotResponse, error) {
	var result BotResponse
	do, err := httpc.Do(context.Background(), http.MethodPost, l.url+l.botPath+path, data)
	if err != nil {
		return nil, err
	}
	resp, err := io.ReadAll(do.Body)
	if err != nil {
		return nil, err
	}
	if err := jsonx.Unmarshal(resp, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
