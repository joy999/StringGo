package utils

import (
	"database/sql"
	"errors"

	"sync"
)

type SqlTx struct {
	*sql.Tx
	db *SqlDB
}

type SqlOp interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Rollback() error
	Close() error
	Commit() error
	Begin() (SqlOp, error)
}

func (this *SqlTx) Close() error {
	return this.Commit()
}

func (this *SqlTx) Begin() (SqlOp, error) {
	return this, nil
}

func (this *SqlDB) Commit() error {
	return errors.New("不支持的方法！")
}

func (this *SqlTx) Commit() error {
	this.db.txLock.Lock()
	defer this.db.txLock.Unlock()

	this.db.txLevel--

	if this.db.txLevel == 0 {
		this.db.tx = nil
		return this.Tx.Commit()
	} else {
		return nil
	}

}
func (this *SqlDB) Rollback() error {
	return errors.New("不支持的方法！")
}
func (this *SqlTx) Rollback() error {
	this.db.txLock.Lock()
	defer this.db.txLock.Unlock()

	this.db.txLevel = 0
	this.db.tx = nil

	return this.Tx.Rollback()

}

type SqlDB struct {
	*sql.DB
	tx      *SqlTx
	txLevel int
	txLock  sync.Mutex
}

func (this *SqlDB) Begin() (SqlOp, error) {
	this.txLock.Lock()
	defer this.txLock.Unlock()

	if this.txLevel == 0 {
		tx, err := this.DB.Begin()
		if err != nil {
			return nil, err
		} else {
			this.tx = new(SqlTx)
			this.tx.Tx = tx
			this.tx.db = this
		}
	}

	this.txLevel++

	return this.tx, nil
}

func (this *SqlDB) Close() error {
	return this.DB.Close()
}

func OpenDatabase(driverName, dataSourceName string) (SqlOp, error) {

	db, err := sql.Open(driverName, dataSourceName)

	sqldb := new(SqlDB)
	sqldb.DB = db
	sqldb.txLevel = 0
	sqldb.tx = nil

	err = db.Ping()

	return sqldb, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
