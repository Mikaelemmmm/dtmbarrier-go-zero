package logic

import (
	"context"

	"fbstrans/stock/internal/svc"
	"fbstrans/stock/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StockLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) StockLogic {
	return StockLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StockLogic) Stock(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
