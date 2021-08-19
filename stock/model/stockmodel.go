package model

import (
	"database/sql"
	"fbstrans/order/model"
	"fmt"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	stockFieldNames          = builderx.RawFieldNames(&Stock{})
	stockRows                = strings.Join(stockFieldNames, ",")
	stockRowsExpectAutoSet   = strings.Join(stringx.Remove(stockFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	stockRowsWithPlaceHolder = strings.Join(stringx.Remove(stockFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
)

type (
	StockModel interface {
		Insert(data Stock) (sql.Result, error)
		FindOne(id int64) (*Stock, error)
		Update(data Stock) error
		Delete(id int64) error
		FindOneByGoodsId(id int64) (*Stock, error)
		UpdateStockNum(data Stock) error
		GetConn() sqlx.SqlConn

	}

	defaultStockModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Stock struct {
		Id         int64     `db:"id"`
		GoodsId    int64     `db:"goods_id"`
		Num        int64     `db:"num"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func NewStockModel(conn sqlx.SqlConn) StockModel {
	return &defaultStockModel{
		conn:  conn,
		table: "`stock`",
	}
}

func (m *defaultStockModel) Insert(data Stock) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, stockRowsExpectAutoSet)
	ret, err := m.conn.Exec(query, data.GoodsId, data.Num)
	return ret, err
}

func (m *defaultStockModel) FindOne(id int64) (*Stock, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", stockRows, m.table)
	var resp Stock
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStockModel) Update(data Stock) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, stockRowsWithPlaceHolder)
	_, err := m.conn.Exec(query, data.GoodsId, data.Num, data.Id)
	return err
}

func (m *defaultStockModel) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.Exec(query, id)
	return err
}


func (m *defaultStockModel) FindOneByGoodsId(id int64) (*Stock, error) {
	query := fmt.Sprintf("select %s from %s where `goods_id` = ? limit 1", stockRows, m.table)
	var resp Stock
	err := m.conn.QueryRow(&resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, model.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultStockModel) UpdateStockNum(data Stock) error {
	query := fmt.Sprintf("update %s set num = num + ?  where `id` = ?", m.table)
	_, err := m.conn.Exec(query,  data.Num, data.Id)
	return err
}

func (m *defaultStockModel) GetConn() sqlx.SqlConn{
	return m.conn
}
