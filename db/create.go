package db

import (
	"log"
)

const (
	sqlCreateProbeResults = `CREATE TABLE probe_results (
    src_ip   INTEGER,
    dest_ip   INTEGER,
    delay INTEGER,
    stamp INTEGER,
    probe_type    INTEGER)`

	sqlProbeResultsDestIndex = `CREATE INDEX i_probe_results_dest ON probe_results (dest_ip ,stamp DESC)`
	sqlPorbeResultStampIndex = `CREATE INDEX i_probe_results_stamp ON probe_results (stamp DESC)`
)

func initTables() error {
	err := checkTable("probe_results", sqlCreateProbeResults)
	if err != nil {
		return err
	}

	err = checkIndex("i_probe_results_dest", sqlProbeResultsDestIndex)
	if err != nil {
		return err
	}

	err = checkIndex("i_probe_results_stamp", sqlPorbeResultStampIndex)
	if err != nil {
		return err
	}
	return nil
}

func checkTable(name, sql string) (err error) {
	stmt, err := db.Prepare("select name from sqlite_master where name =? and type = 'table'")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		log.Println("Database table", name, "ok")
		return nil
	}

	_, err = db.Exec(sql)

	if err == nil {
		log.Println("Database table", name, "created")
	}
	return err
}

func checkIndex(name, sql string) error {
	stmt, err := db.Prepare("select name from sqlite_master where name =? and type = 'index'")
	if err != nil {
		return err
	}
	defer stmt.Close()
	rows, err := stmt.Query(name)
	if err != nil {
		return err
	}
	defer rows.Close()
	if rows.Next() {
		log.Println("Database index", name, "ok")
		return nil
	}

	_, err = db.Exec(sql)
	if err == nil {
		log.Println("Database index", name, "created")
	}
	return err
}
