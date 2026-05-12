package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {

	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASSWORD")
	dbHostPort := os.Getenv("MYSQL_HOST_PORT")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s", dbUser, dbPass, dbHostPort, dbName))
	if err != nil {
		panic(fmt.Errorf("cannot open mysql db"))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{db: db}
}
