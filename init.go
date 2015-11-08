package myopenid

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func DbConnection(connString string) (*MysqlDiscoveryCache, *MysqlNonceStore) {
	var err error
	db, err = sql.Open("mysql", connString)
	if err != nil {
		log.Printf("\nMYSQLOPENID ERROR: %s", err.Error())
	}

	discoveryCache := new(MysqlDiscoveryCache)
	nonceStore := new(MysqlNonceStore)

	return discoveryCache, nonceStore
}
