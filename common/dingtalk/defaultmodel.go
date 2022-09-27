package dingtalk

type (
	defaultModel struct {
		url     string
		botPath string
	}
)

func newDefaultModel() *defaultModel {
	return &defaultModel{
		url:     "https://oapi.dingtalk.com",
		botPath: "/robot/send?access_token=",
	}
}
