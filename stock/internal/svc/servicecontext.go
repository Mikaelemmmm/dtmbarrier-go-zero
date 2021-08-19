package svc

import (
	"fbstrans/dtmzero"
	"fbstrans/stock/internal/config"
	"fbstrans/stock/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	BarrierModel dtmzero.BarrierModel
	StockModel model.StockModel

}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BarrierModel: dtmzero.NewBarrierModel(sqlx.NewMysql(c.BarrierDataSource)),
		StockModel: model.NewStockModel(sqlx.NewMysql(c.DataSource)),
	}
}
