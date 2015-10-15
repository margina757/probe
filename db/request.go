package db

import (
	"database/sql"
)

func QuerySQL(sql string) (rows *sql.Rows, e error) {
	ctx := newContextSQL()
	ctx.sql = sql
	ctx.ctxType = queryType
	chanContexSQL <- ctx

	rows = <-ctx.rows
	e = ctx.err
	return
}

func ExecSQL(sql string) (e error) {
	ctx := newContextSQL()
	ctx.sql = sql
	ctx.ctxType = execType
	chanContexSQL <- ctx

	_ = <-ctx.rows

	return ctx.err
}

func doQuerySQL(ctx *contexSQL) {
	rows, e := conn.Query(ctx.sql)
	ctx.err = e
	ctx.rows <- rows
	close(ctx.rows)
}

func doExecSQL(ctx *contexSQL) {
	_, ctx.err = conn.Exec(ctx.sql)
	ctx.rows <- nil
	close(ctx.rows)
}

// func QueryStmt(stmt *sql.Stmt) (rows *sql.Rows, e error) {
// 	c := newContext()
// 	c.stmt = c.stmt
// 	chanContex <- c

// 	rows = <-c.rows
// 	e = c.err
// 	return
// }

// func ExecStmt(stmt *sql.Stmt) (e error) {
//  c := newContext()
//  c.stmt = stmt
//  chanContex <- c
//  _ = <-c.rows
//  return c.err
// }

// func doQueryStmt(c *contex) {
// 	rows, e := c.stmt.Query(...)
// 	c.err = e
// 	c.rows <- rows
// 	close(c.rows)
// 	return
// }
