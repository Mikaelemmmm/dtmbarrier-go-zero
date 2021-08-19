package svc

import (
	"fbstrans/dtmzero"
	"fbstrans/order/internal/config"
	"fbstrans/order/model"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	BarrierModel dtmzero.BarrierModel
	OrderModel model.OrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		BarrierModel: dtmzero.NewBarrierModel(sqlx.NewMysql(c.BarrierDataSource)),
		OrderModel: model.NewOrderModel(sqlx.NewMysql(c.DataSource)),
	}
}
