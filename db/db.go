package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"probe/internal/types"
	"probe/internal/utils"
	"sync"
)

const (
	queryType = 1
	execType  = 2
)

const (
	tProbeResults = "probe_results"
)

var (
	doneChan        chan int
	chanProbeResult chan *types.ProbeResult

	db *sql.DB
	wg sync.WaitGroup
)

func init() {
	doneChan = make(chan int)
	chanProbeResult = make(chan *types.ProbeResult, 1024)
}

/*
OPenSQL打开SQLite数据库并创建一个线程等待数据库操作。
*/
func OpenDB() error {

	openDB()
	e := initTables()
	if e != nil {
		log.Println(e)
		return e
	}
	log.Println("Database table check ok")

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

	stmtInsertProbeResult, err := db.Prepare("insert into probe_results values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return
	}
	defer stmtInsertProbeResult.Close()
	for {
		select {
		case probe := <-chanProbeResult:
			insertProbeResult(stmtInsertProbeResult, probe)
		case _ = <-doneChan:
			return
		}
	}
}

func openDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./lonlife.db")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Open database lonlife.db success")
	}
	return nil
}

func insertProbeResult(stmt *sql.Stmt, result *types.ProbeResult) {
	src := utils.IPToIntString(result.Src)
	dest := utils.IPToIntString(result.Dest)
	_, e := stmt.Exec(src, dest, result.Delay, result.Stamp, result.Type)
	if e != nil {
		log.Println(e)
	}
}
