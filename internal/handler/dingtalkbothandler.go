package handler

import (
	"bk_monitor_notify/common/response"
	"net/http"

	"bk_monitor_notify/internal/logic"
	"bk_monitor_notify/internal/svc"
	"bk_monitor_notify/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func dingTalkBotHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BKCallBack
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewDingTalkBotLogic(r.Context(), ctx)
		err := l.DingTalkBot(&req)
		response.Response(w, nil, err)

	}
}
