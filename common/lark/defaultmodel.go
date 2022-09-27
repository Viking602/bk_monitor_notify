package lark

type (
	defaultModel struct {
		url     string
		botPath string
	}
)

func newDefaultModel() *defaultModel {
	return &defaultModel{
		url:     "https://open.feishu.cn",
		botPath: "/open-apis/bot/v2/hook/",
	}
}
