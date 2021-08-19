package handler

import (
	"fbstrans/dtmzero"
	"fbstrans/xhttp"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"net/http"

	"fbstrans/order/internal/logic"
	"fbstrans/order/internal/svc"
	"fbstrans/order/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func CreateOrderHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrderRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		fmt.Printf("rFrom:%+v",r.Form)

		l := logic.NewCreateOrderLogic(r.Context(), ctx)

		//子事务屏蔽
		barrier,err:=dtmzero.BarrierFromQuery(r.Form)
		if err != nil{
			logx.Error("barrier err:%v",err)
		}
		var resp *types.CreateOrderResponse
		err = barrier.Call(ctx.BarrierModel.GetDBConn(), func(db sqlx.Session) error {
			resp, err = l.CreateOrder(req)
			if err != nil{
				return err
			}
			return nil
		})

		xhttp.HttpDTMResult(r,w,resp,err)
	}
}
