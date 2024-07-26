package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/config"
)

func NewPostgreSQLStorage(cfg config.Config) (*sql.DB, error) {
	//Open connection to DB
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
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