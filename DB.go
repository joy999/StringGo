package utils

import (
	"database/sql"
	"errors"

	"sync"
)

var ODB *SqlDB = nil
var ODBLevel int = 0
var ODBLock sync.Mutex

//var ODBStatus int

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
	ODBLock.Lock()
	defer ODBLock.Unlock()

	ODBLevel--

	if ODBLevel == 0 {
		//ODBStatus = 0 //设置数据库为关闭状态
		ODB = nil
		return this.DB.Close()
	}
	return nil
}

func OpenDatabase(driverName, dataSourceName string) (SqlOp, error) {
	ODBLock.Lock()
	defer ODBLock.Unlock()

	if ODB != nil {
		if ODB.tx != nil {
			//现在正处理事务中，则添加一层事务处理
			return ODB.Begin()
		} else {

			ODBLevel++
			return ODB, nil
		}
	}

	db, err := sql.Open(driverName, dataSourceName)

	ODBLevel++
	sqldb := new(SqlDB)
	sqldb.DB = db
	sqldb.txLevel = 0
	sqldb.tx = nil

	ODB = sqldb

	return ODB, err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
