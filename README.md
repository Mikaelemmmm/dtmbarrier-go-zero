### go-zero使用分布式事务dtm的barrier

在go-zero中使用分布式事务[dtm](https://github.com/yedf/dtm)时候，使用它的子事务屏蔽功能，发现必须要暴露出sql.DB，但是go-zero封装在底层无法拿到sql.DB，这里基于go-zero的sqlx，修改了下，可以使用go-zero的sqlx来替代dtmcli的barrier中的sql.DB

#### http使用示例

调用 http://127.0.0.1:8001/order/quickOrder 

下订单、扣库存



使用方法 handler中

【注】：这里只是演示用法，更优雅方式在logic中处理

```go
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
```

BarrierModel单独一个库保存，只需要暴露出来db连接即可

```go
package dtmzero

import (
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlx"
)


type (
	BarrierModel interface {
		GetDBConn() sqlx.SqlConn
	}

	defaultBarrierModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Barrier struct {
		Id         int64     `db:"id"`
		TransType  string    `db:"trans_type"`
		Gid        string    `db:"gid"`
		BranchId   string    `db:"branch_id"`
		BranchType string    `db:"branch_type"`
		BarrierId  string    `db:"barrier_id"`
		Reason     string    `db:"reason"` // the branch type who insert this record
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewBarrierModel(conn sqlx.SqlConn) BarrierModel {
	return &defaultBarrierModel{
		conn:  conn,
		table: "`barrier`",
	}
}

func (m defaultBarrierModel) GetDBConn() sqlx.SqlConn {
	return m.conn
}
```



#### grpc使用示例







