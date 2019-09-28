package config

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {
	db, err = sql.Open("mssql", "server=localhost;user id=bujosa;password=123;database=EnZonaRD;")
	return
}
