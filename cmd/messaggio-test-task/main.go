package main

import (
	"fmt"
	"log"

	"github.com/zhetkerbaevan/messaggio-test-task/cmd/api"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/config"
	"github.com/zhetkerbaevan/messaggio-test-task/internal/db"
)

func main() {
	db, err := db.NewPostgreSQLStorage(config.Config{
		DBHost: config.Envs.DBHost,
		DBPort: config.Envs.DBPort,
		DBUser: config.Envs.DBUser,
		DBPassword: config.Envs.DBPassword,
		DBName: config.Envs.DBName,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to DB")

	//Start server
	server := api.NewAPIServer(db, config.Envs.Port)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}