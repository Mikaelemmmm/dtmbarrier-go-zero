// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"fbstrans/stock/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/from/:name",
				Handler: StockHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/stock/deduct",
				Handler: StockDeductHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/stock/cancelDeduct",
				Handler: CancelStockDeductHandler(serverCtx),
			},
		},
	)
}
