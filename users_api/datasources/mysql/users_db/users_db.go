package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	envErr := godotenv.Load()
	var (
		username = os.Getenv("MYSQL_USERNAME")
		password = os.Getenv("MYSQL_PASSWORD")
		host     = os.Getenv("MYSQL_HOST")
		schema   = os.Getenv("MYSQL_SCHEMA")
	)
	if envErr != nil {
		log.Fatal(envErr)
	}
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
}
