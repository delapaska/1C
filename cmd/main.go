package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/delapaska/1C/cmd/api"
	"github.com/delapaska/1C/configs"
	"github.com/delapaska/1C/db"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// @title 1C API
// @version 13.37
// @description САМЫЙ АХУЕННЫЙ СЕРВИС ДЛЯ РАБОТЫ С 1С

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.Envs.Host, configs.Envs.DBPort,
		configs.Envs.DBUser, configs.Envs.DBPassword, configs.Envs.DBName)
	fmt.Println(configs.Envs.DBPassword)
	db, err := db.NewPostgreSQLstorage(psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewApiServer(db)

	server.Run()

}

func initStorage(db *sql.DB) {
	err := db.Ping()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected")
}
