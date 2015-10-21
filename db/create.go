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

	sqlCreateProbeResultsIndex = `CREATE INDEX i_probe_results ON probe_results (dest_ip ,stamp DESC)`
)

func initTables() error {
	err := checkProbeResultsTable()
	if err != nil {
		return err
	}
	return nil
}

func checkTable(table string) (exist bool, err error) {
	stmt, err := db.Prepare("select name from sqlite_master where name =? and type = 'table'")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(table)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}

func createTable(sql string) (err error) {
	_, err = db.Exec(sql)
	return err
}
func checkProbeResultsTable() error {
	exist, err := checkTable(tProbeResults)
	if err != nil {
		return err
	}
	if !exist {
		err = createTable(sqlCreateProbeResults)
		if err != nil {
			return err
		}
		log.Println("Database table probe_results created")
		err = createTable(sqlCreateProbeResultsIndex)
		if err == nil {
			log.Println("Database index i_probe_results created")
		}
		return err
	}
	return nil
}

func createProbeResultsTable() error {
	return nil
}
