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
