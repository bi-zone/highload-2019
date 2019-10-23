package models

import (
	"database/sql"
	"fmt"
)

// Config holds the configuration used for database
type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBname   string
}

// NewDB Creates new PostgresDB
func NewDB(cfg Config) (*sql.DB, error) {
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Index indexs song in db
func Index(db *sql.DB, songName string, hashArray []int) error {
	// TODO: Implement
	return nil
}

// Recognize recognizes song in db
func Recognize(db *sql.DB, hashArray []int) (string, error) {
	// TODO: Implement
	return "Not Found", nil
}

// Delete deletes song from bd
func Delete(db *sql.DB, name string) (affected int64, err error) {
	// TODO: Implement
	return 0, nil
}
