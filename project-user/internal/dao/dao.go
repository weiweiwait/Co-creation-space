package dao

import (
	"my_project/project-user/internal/database"
	"my_project/project-user/internal/database/gorms"
)

type TransactionImpl struct {
	conn database.DbConn
}

func (t *TransactionImpl) Action(f func(conn database.DbConn) error) error {
	//同一个事务结束后连接就没有了，所以每次得重新连接
	t.conn.Begin()
	err := f(t.conn)
	if err != nil {
		t.conn.Rollback()
		return err
	}
	t.conn.Commit()
	return nil
}

func NewTransaction() *TransactionImpl {
	return &TransactionImpl{
		conn: gorms.NewTran(),
	}
}
