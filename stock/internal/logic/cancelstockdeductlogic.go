package logic

import (
	"context"
	"fbstrans/stock/model"
	"fmt"

	"fbstrans/stock/internal/svc"
	"fbstrans/stock/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CancelStockDeductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCancelStockDeductLogic(ctx context.Context, svcCtx *svc.ServiceContext) CancelStockDeductLogic {
	return CancelStockDeductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CancelStockDeductLogic) CancelStockDeduct(req types.DeductRequest) (*types.DeductResponse, error) {

	stock,err:=l.svcCtx.StockModel.FindOneByGoodsId(req.GoodsId)
	if err != nil{
		return &types.DeductResponse{}, fmt.Errorf("库存不存在 goodsId:%d,err:%v",req.GoodsId,err)
	}

	var data model.Stock
	data.Id = stock.Id
	data.Num =req.Num
	if err:=l.svcCtx.StockModel.UpdateStockNum(data);err!= nil{
		return  &types.DeductResponse{}, fmt.Errorf("库存不足 goodsId:%d ,err:%v",req.GoodsId,err)
	}

	return &types.DeductResponse{}, nil
}
