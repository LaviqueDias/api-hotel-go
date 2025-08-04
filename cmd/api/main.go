package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	"github.com/LaviqueDias/api-hotel-go/internal/database/postgresql"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Olá, mundo")
	db, err := postgresql.Connection()
	if err != nil {
		log.Fatal("Err to connect to bd", err)
	}
	defer postgresql.CloseConnection()

	p := db.Ping()
	fmt.Println(p)
}