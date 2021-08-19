package logic

import (
	"context"
	"fbstrans/stock/model"
	"fmt"

	"fbstrans/stock/internal/svc"
	"fbstrans/stock/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StockDeductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStockDeductLogic(ctx context.Context, svcCtx *svc.ServiceContext) StockDeductLogic {
	return StockDeductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StockDeductLogic) StockDeduct(req types.DeductRequest) (*types.DeductResponse, error) {


	stock,err:=l.svcCtx.StockModel.FindOneByGoodsId(req.GoodsId)
	if err != nil{
		return &types.DeductResponse{}, fmt.Errorf("库存不存在 goodsId:%d,err:%v",req.GoodsId,err)
	}

	var data model.Stock
	data.Id = stock.Id
	data.Num =0 - req.Num
	if err:=l.svcCtx.StockModel.UpdateStockNum(data);err!= nil{
		return  &types.DeductResponse{}, fmt.Errorf("库存不足 goodsId:%d ,err:%v",req.GoodsId,err)
	}

	//return nil,fmt.Errorf("故意的")

	return &types.DeductResponse{}, nil
}
