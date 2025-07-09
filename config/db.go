package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Serve para inicializar a conexão com banco de dados
func SetupDB() *sql.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	dbConnection, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database")

	return dbConnection
}
