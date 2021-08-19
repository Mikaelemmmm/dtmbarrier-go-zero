package handler

import (
	"fbstrans/xhttp"
	"net/http"

	"fbstrans/order/internal/logic"
	"fbstrans/order/internal/svc"
	"fbstrans/order/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func QuickOrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewQuickOrderLogic(r.Context(), ctx)
		resp, err := l.QuickOrder(req)
		xhttp.HttpResult(r,w,resp,err)
	}
}
