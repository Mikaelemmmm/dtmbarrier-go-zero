package handler

import (
	"fbstrans/dtmzero"
	"fbstrans/xhttp"
	"fmt"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"net/http"

	"fbstrans/stock/internal/logic"
	"fbstrans/stock/internal/svc"
	"fbstrans/stock/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func StockDeductHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeductRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		fmt.Println(222222)

		l := logic.NewStockDeductLogic(r.Context(), ctx)

		//子事务屏蔽
		barrier,err:=dtmzero.BarrierFromQuery(r.Form)
		if err != nil{
			logx.Error("barrier err:%v",err)
		}
		var resp *types.DeductResponse
		err = barrier.Call(ctx.BarrierModel.GetDBConn(), func(db sqlx.Session) error {
			resp, err = l.StockDeduct(req)
			if err != nil{
				return err
			}
			return nil
		})
		xhttp.HttpDTMResult(r,w,resp,err)
	}
}
