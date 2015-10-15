package db

import (
	"database/sql"
	"sync"
)

const (
	queryType = 1
	execType  = 2
)

var (
	doneChan      chan int
	chanContex    chan *contex
	chanContexSQL chan *contexSQL
	conn          *sql.DB
	wg            sync.WaitGroup
)

type contex struct {
	stmt    *sql.Stmt
	rows    chan *sql.Rows
	ctxType int
	err     error
}

type contexSQL struct {
	sql     string
	rows    chan *sql.Rows
	ctxType int
	err     error
}

func init() {
	doneChan = make(chan int)
	chanContex = make(chan *contex, 32)
	chanContexSQL = make(chan *contexSQL, 32)
}

/*
OPenSQL打开SQLite数据库并创建一个线程等待数据库操作。
*/
func OpenDB() error {
	wg.Add(1)
	go doDB()
	return nil
}

/*
CloseSQL退出数据库线程并断开数据库连接。
*/
func CloseDB() {
	doneChan <- 1
	wg.Wait()
}

func doDB() {
	defer wg.Done()
	for {
		select {
		case _ = <-chanContex:
		case ctx := <-chanContexSQL:
			if ctx.ctxType == queryType {
				doQuerySQL(ctx)
				break
			}
			if ctx.ctxType == execType {
				doExecSQL(ctx)
				break
			}
		case _ = <-doneChan:
			return
		}
	}
}

func newContext() *contex {
	c := &contex{}
	c.rows = make(chan *sql.Rows)
	return c
}

func newContextSQL() *contexSQL {
	c := &contexSQL{}
	c.rows = make(chan *sql.Rows)
	return c
}
