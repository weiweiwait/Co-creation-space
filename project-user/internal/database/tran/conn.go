package tran

type DbConn interface {
	Begin()
	Rollback()
	Commit()
}
