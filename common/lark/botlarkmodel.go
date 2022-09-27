package lark

import (
	"context"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/rest/httpc"
	"io"
	"net/http"
)

var _ BotLark = (*larkDefaultModel)(nil)

type (
	BotLark interface {
		Bot(path string, data *BotCard) (*BotResponse, error)
	}
	larkDefaultModel struct {
		*defaultModel
	}
	BotCard struct {
		Card      Card   `json:"card"`
		MsgType   string `json:"msg_type"`
		Sign      string `json:"sign"`
		Timestamp int64  `json:"timestamp"`
	}

	Card struct {
		Config   CardConfig  `json:"config"`
		Elements interface{} `json:"elements"`
		Header   CardHeader  `json:"header"`
	}

	CardHeader struct {
		Template string    `json:"template"`
		Title    *CardText `json:"title"`
	}

	CardConfig struct {
		WideScreenMode bool `json:"wide_screen_mode"`
	}

	CardElements struct {
		Fields  []*CardFields `json:"fields,omitempty"`
		Tag     string        `json:"tag"`
		Content string        `json:"content,omitempty"`
	}

	CardFields struct {
		IsShort bool      `json:"is_short"`
		Text    *CardText `json:"text"`
	}
	CardText struct {
		Content string `json:"content"`
		Tag     string `json:"tag"`
	}

	BotResponse struct {
		Code          int         `json:"code,omitempty"`
		Data          interface{} `json:"data,omitempty"`
		Msg           string      `json:"msg,omitempty"`
		StatusCode    int64       `json:"StatusCode,omitempty"`
		StatusMessage string      `json:"StatusMessage,omitempty"`
	}
)

func NewLarkBot() BotLark {
	return &larkDefaultModel{
		defaultModel: newDefaultModel(),
	}
}

func (l larkDefaultModel) Bot(path string, data *BotCard) (*BotResponse, error) {
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
