package logic

import (
	"context"
	"fmt"

	"fbstrans/order/internal/svc"
	"fbstrans/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/yedf/dtm/dtmcli"

)

const DtmServer = "http://localhost:8080/api/dtmsvr"
const OrderServer = "http://localhost:8001"
const StockServer = "http://localhost:8002"

type QuickOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuickOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) QuickOrderLogic {
	return QuickOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuickOrderLogic) QuickOrder(req types.CreateOrderRequest) (*types.CreateOrderResponse, error) {

	saga := dtmcli.NewSaga(DtmServer, dtmcli.MustGenGid(DtmServer)).
		Add(OrderServer+"/order/createOrderRequest", OrderServer+"/order/cancelOrderRequest", req).
		Add(StockServer+"/stock/deduct", StockServer+"/stock/cancelDeduct", req)
	if err := saga.Submit();err!= nil{
		return nil,fmt.Errorf("saga err: %v",err)
	}

	return &types.CreateOrderResponse{}, nil
}
