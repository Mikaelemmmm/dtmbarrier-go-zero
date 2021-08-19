package model

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
	"strings"
)

var (
	orderFieldNames          = builderx.RawFieldNames(&Order{})
	orderRows                = strings.Join(orderFieldNames, ",")
	orderRowsExpectAutoSet   = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	orderRowsWithPlaceHolder = strings.Join(stringx.Remove(orderFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	OrderModel interface {
		Insert(data Order) (sql.Result, error)
		FindOne(id int64) (*Order, error)
		Update(data Order) error
		Delete(id int64) error
		UpdateStatus(data Order) error
		FindOneByUserIdGoodsId(userId,goodsId int64) (*Order, error)
		GetConn() sqlx.SqlConn
	}

	defaultOrderModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Order struct {
		Id         int64  `db:"id"`
		GoodsId    int64  `db:"goods_id"`
		Num        int64  `db:"num"`
		UserId     int64  `db:"user_id"`
		CreateTime string `db:"create_time"`
		UpdateTime string `db:"update_time"`
		Status     int64  `db:"status"`
	}
)

func NewOrderModel(conn sqlx.SqlConn) OrderModel {
	return &defaultOrderModel{
		conn:  conn,
		table: "`order`",
	}
}

func (m *defaultOrderModel) Insert(data Order) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, orderRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.GoodsId, data.Num, data.UserId,data.Status)
	return ret, err
}

func (m *defaultOrderModel) FindOne(id int64) (*Order, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderRows, m.table)
	var resp Order
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderModel) Update(data Order) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.GoodsId, data.Num, data.UserId, data.Id)
	return err
}

func (m *defaultOrderModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}

func (m *defaultOrderModel) UpdateStatus(data Order) error {
	query := fmt.Sprintf("update %s set status = ?  where `id` = ?", m.table)
	_, err := m.conn.Exec(query, data.Status, data.Id)
	return err
}

func (m *defaultOrderModel) FindOneByUserIdGoodsId(userId,goodsId int64) (*Order, error) {
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and goods_id = ?  limit 1", orderRows, m.table)
	var resp Order
	err := m.conn.QueryRow(&resp, query, userId,goodsId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *defaultOrderModel) GetConn() sqlx.SqlConn{
	return m.conn
}
