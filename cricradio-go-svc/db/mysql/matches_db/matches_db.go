package matches_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB

	username = "root"
	password = "password"
	host     = "mysql-db"
	schema   = "matches_db"
)

func init() {
	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	log.Println(fmt.Sprintf("about to connect to %s\n", dataSource))

	Client, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

	/*
		To Do:

		load Match URLs from DB on start
		create DB & Schema if not exists

	*/

}
