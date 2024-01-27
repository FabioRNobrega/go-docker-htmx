package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func initDB() (*pgx.Conn, error) {
	connConfig, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, fmt.Errorf("error parsing connection string: %w", err)
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return conn, nil
}

func main() {
	conn, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if conn != nil {
			conn.Close(context.Background())
		}
	}()

	fmt.Println("Connected to the PostgreSQL database successfully!")
}
