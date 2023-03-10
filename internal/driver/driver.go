package driver

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

// DB holds the database pool
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxDBOpenConn = 10
	maxDBIdleConn = 5
	maxDbLifeTime = time.Minute * 5
)

// ConnectSql create database pool for postgres
func ConnectSql(dsn string) (*DB, error) {
	conn, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}
	conn.SetConnMaxLifetime(maxDbLifeTime)
	conn.SetMaxIdleConns(maxDBIdleConn)
	conn.SetMaxOpenConns(maxDBOpenConn)
	dbConn.SQL = conn
	err = testDB(dbConn.SQL)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}

func testDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func NewDatabase(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}
