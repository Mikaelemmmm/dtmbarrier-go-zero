package logic

import (
	"context"
	"fbstrans/order/model"
	"fmt"

	"fbstrans/order/internal/svc"
	"fbstrans/order/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CreateOrderCancelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderCancelLogic(ctx context.Context, svcCtx *svc.ServiceContext) CreateOrderCancelLogic {
	return CreateOrderCancelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderCancelLogic) CreateOrderCancel(req types.CreateOrderRequest) (*types.CreateOrderResponse, error) {

	fmt.Printf("req:%+v\n",req)
	order,err:=l.svcCtx.OrderModel.FindOneByUserIdGoodsId(req.UserId,req.GoodsId)
	if err != nil && err != model.ErrNotFound{
		return nil,fmt.Errorf("订单不存在")
	}
	fmt.Printf("2223343434343 \n")

	if order != nil{
		updateOrder:= *order
		updateOrder.Status = -1
		_ = l.svcCtx.OrderModel.UpdateStatus(updateOrder)
	}else{
		logx.Info("订单不存在，不需要处理")
	}

	return &types.CreateOrderResponse{}, nil
}
