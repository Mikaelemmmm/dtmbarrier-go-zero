package handler

import (
	"net/http"

	"fbstrans/stock/internal/logic"
	"fbstrans/stock/internal/svc"
	"fbstrans/stock/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func StockHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewStockLogic(r.Context(), ctx)
		resp, err := l.Stock(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
