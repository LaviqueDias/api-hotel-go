package postgresql

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/LaviqueDias/api-hotel-go/internal/configuration/logger"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var db *sql.DB

func Connection() (*sql.DB, error) {
	logger.Info("Init Connection with Postgres",
		zap.String("journey", "connectionToPostgres"),
	)
	if db == nil {
		var err error
		db, err = sql.Open("postgres", getPostgresDSN())
		if err != nil {
			return nil, err
		}

		// Ajustar o timezone da sessão (não trava a aplicação se falhar)
		if _, err := db.Exec("SET TIME ZONE 'America/Sao_Paulo'"); err != nil {
			logger.Error("warning: não foi possível ajustar time zone da sessão", err,
				zap.String("journey", "connectionToPostgres"))
		}

		if err := db.Ping(); err != nil {
			return nil, err
		}
	}

	logger.Info("Connection with Postgres executed successfully",
		zap.String("journey", "connectionToPostgres"),
	)
	return db, nil
}

func CloseConnection() {
	logger.Info("Init CloseConnection with Postgres",
		zap.String("journey", "closeConnectionToPostgres"),
	)
	if db != nil {
		_ = db.Close()
	}
	logger.Info("CloseConnection with Postgres executed successfully",
		zap.String("journey", "closeConnectionToPostgres"),
	)
}

func getPostgresDSN() string {
	logger.Info("Init GetPostgresDSN",
		zap.String("journey", "connectionToPostgres"),
	)
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	logger.Info("GetPostgresDSN executed successfully",
		zap.String("journey", "connectionToPostgres"),
	)
	// Exemplo: "host=localhost port=5432 user=root password=root dbname=api-hotel-go sslmode=disable"
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)
}


