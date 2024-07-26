package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/config"
)

func NewPostgreSQLStorage(cfg config.Config) (*sql.DB, error) {
	//Open connection to DB
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		return nil, err
	}

	//Verify connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}