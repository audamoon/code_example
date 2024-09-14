package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host        string
	Port        string
	DBName      string
	User        string
	Password    string
	SSLMode     string
	MaxConn     int
	MaxIdleConn int
}

var Instance *sql.DB

func Init(c *Config) error {
	if Instance != nil {
		return nil
	}

	connSettings := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.Host,
		c.Port,
		c.DBName,
		c.User,
		c.Password,
		c.SSLMode,
	)

	db, err := sql.Open("postgres", connSettings)
	if err != nil {
		return fmt.Errorf("db init: %w", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("db ping: %w", err)
	}

	db.SetMaxOpenConns(c.MaxConn)
	db.SetMaxIdleConns(c.MaxIdleConn)

	Instance = db

	return nil
}
