package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
}

type MySQLDB struct {
	db *sql.DB
}

func New(cfg Config) *MySQLDB {
	
	// user:password@tcp(host:port)/dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Errorf("cannot open mysql db: %w", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Errorf("cannot connect to mysql db: %w", err))
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{db: db}
}