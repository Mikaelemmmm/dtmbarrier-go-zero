package dtmzero

import (
	"database/sql"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/test/model"
	"github.com/yedf/dtm/dtmcli"
	"net/url"
)

// BusiFunc type for busi func
type BusiFunc func(db sqlx.Session) error

// BranchBarrier every branch info
type BranchBarrier struct {
	TransType  string
	Gid        string
	BranchID   string
	BranchType string
	BarrierID  int
}

func (bb *BranchBarrier) String() string {
	return fmt.Sprintf("transInfo: %s %s %s %s", bb.TransType, bb.Gid, bb.BranchID, bb.BranchType)
}

// BarrierFromQuery construct transaction info from request
func BarrierFromQuery(qs url.Values) (*BranchBarrier, error) {
	return BarrierFrom(qs.Get("trans_type"), qs.Get("gid"), qs.Get("branch_id"), qs.Get("branch_type"))
}

// BarrierFrom construct transaction info from request
func BarrierFrom(transType, gid, branchID, branchType string) (*BranchBarrier, error) {
	ti := &BranchBarrier{
		TransType:  transType,
		Gid:        gid,
		BranchID:   branchID,
		BranchType: branchType,
	}
	if ti.TransType == "" || ti.Gid == "" || ti.BranchID == "" || ti.BranchType == "" {
		return nil, fmt.Errorf("invlid trans info: %v", ti)
	}
	return ti, nil
}

func insertBarrier(session sqlx.Session, transType string, gid string, branchID string, branchType string, barrierID string, reason string) (int64, error) {
	if branchType == "" {
		return 0, nil
	}
	result,err:= session.Exec("insert ignore into barrier(trans_type, gid, branch_id, branch_type, barrier_id, reason) values(?,?,?,?,?,?)", transType, gid, branchID, branchType, barrierID, reason)
	if err != nil{
		return 0,err
	}
	return result.RowsAffected()
}

// Call 子事务屏障，详细介绍见 https://zhuanlan.zhihu.com/p/388444465
// db: 本地数据库
// transInfo: 事务信息
// bisiCall: 业务函数，仅在必要时被调用
// 返回值:
// 如果发生悬挂，则busiCall不会被调用，直接返回错误 ErrFailure，全局事务尽早进行回滚
// 如果正常调用，重复调用，空补偿，返回的错误值为nil，正常往下进行
func (bb *BranchBarrier) Call(db sqlx.SqlConn, busiCall BusiFunc) (rerr error) {
	bb.BarrierID = bb.BarrierID + 1
	bid := fmt.Sprintf("%02d", bb.BarrierID)

	ti := bb
	originType := map[string]string{
		"cancel":     "try",
		"compensate": "action",
	}[ti.BranchType]


	return db.Transact(func(session sqlx.Session) error {

		originAffected, err := insertBarrier( session,ti.TransType, ti.Gid, ti.BranchID, originType, bid, ti.BranchType)
		if err != nil{
			return err
		}
		currentAffected, err := insertBarrier(session, ti.TransType, ti.Gid, ti.BranchID, ti.BranchType, bid, ti.BranchType)
		if err != nil{
			return err
		}

		if (ti.BranchType == "cancel" || ti.BranchType == "compensate") && originAffected > 0 { // 这个是空补偿，返回成功
			return nil
		} else if currentAffected == 0 { // 插入不成功
			var result sql.NullString
			query:= "select 1 from barrier where trans_type=? and gid=? and branch_id=? and branch_type=? and barrier_id=? and reason=?"
			err = session.QueryRow(&result,query,ti.TransType, ti.Gid, ti.BranchID, ti.BranchType, bid, ti.BranchType)

			if err == model.ErrNotFound { // 不是当前分支插入的，那么是cancel插入的，因此是悬挂操作，返回失败，AP收到这个返回，会尽快回滚
				rerr = dtmcli.ErrFailure
				return rerr
			}

			return nil
		}
		return busiCall(session)
	})

}



