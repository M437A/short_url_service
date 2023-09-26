package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
	"os"
	"time"
)

var DB *pgxpool.Pool

func GetDBPool() *pgxpool.Pool {
	connectionString := viper.GetString("db.url")
	return connectToDB(connectionString)
}

func connectToDB(connectionString string) *pgxpool.Pool {
	for i := 0; i < 3; i++ {
		dbpool, err := pgxpool.New(context.Background(), connectionString)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when creating connection pool: %v\n", err)
			fmt.Println("Retrying connection after 5 seconds...")
			time.Sleep(5 * time.Second)
		}

		fmt.Println("The connection to the PostgreSQL database is established!")

		DB = dbpool
		return dbpool
	}
	panic("Can't connect to db")
}
