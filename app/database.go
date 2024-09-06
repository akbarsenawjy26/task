package app

import (
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open(`postgres`, `postgresql://postgres:403201aa@localhost:5432/postgres?sslmode=disable`)
	if err != nil {
		return nil
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
