package logic

import (
	"context"
	"fbstrans/order/model"
	"fmt"

	"fbstrans/order/internal/svc"
	"fbstrans/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateOrderLogic {
	return CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req types.CreateOrderRequest) (*types.CreateOrderResponse, error) {


	var orderData model.Order
	orderData.UserId = req.UserId
	orderData.GoodsId = req.GoodsId
	orderData.Num = req.Num
	if _,err:=l.svcCtx.OrderModel.Insert(orderData);err!= nil{
		return &types.CreateOrderResponse{}, fmt.Errorf("下单失败,err:%v",err)
	}

	return &types.CreateOrderResponse{}, nil
}
