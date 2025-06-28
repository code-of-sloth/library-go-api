package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var DbConn *pgx.Conn

func initDb() (err error) {
	dbUser := os.Getenv(dbUserEnv)
	dbPsw := os.Getenv(dbPswEnv)
	dbName := os.Getenv(dbNameEnv)
	dbPort := os.Getenv(dbPortEnv)
	url := fmt.Sprintf("postgres://%s:%s@db:%s/%s", dbUser, dbPsw, dbPort, dbName)
	log.Printf("\n initdb:db url:%s", url)
	DbConn, err = pgx.Connect(context.Background(), url)
	if err != nil {
		err = fmt.Errorf("initdb:error connecting to db: %w", err)
		return
	}
	return
}
