package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var (
	DB   *pgxpool.Pool
	once sync.Once
)

func ConnectDb() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalln("No .env file loaded", err.Error())
			return
		}

		dbURL := os.Getenv("POSTGRES_URL")
		if dbURL == "" {
			log.Fatalln("No URL found in .env")
			return
		}

		config, err := pgxpool.ParseConfig(dbURL)
		if err != nil {
			log.Fatalln("Unable to parse pgxpool", err.Error())
			return
		}

		config.MaxConns = 10
		config.MinConns = 2
		config.HealthCheckPeriod = 1 * time.Minute

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalln("Error while parsing with new configuration", err.Error())
			return
		}

		DB = pool
		log.Println("Connected to db: ecommerce_order_management_system_payments")
	})
}

func GetDB() *pgxpool.Pool {
	if DB == nil {
		ConnectDb()
	}
	return DB
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		log.Fatalln("Database connection closed: ecommerce_order_management_system_payments")
	}
}
