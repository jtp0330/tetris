package db

import (
	"fmt"

	"github.com/lib/pq"
)

// connect to user_data postgres to record user data, and score in postgres docker container
func Postgres_Connect() {
	// connect to postgres db
	pgUrl, err := pq.ParseURL("postgres://postgres:postgres@localhost:5432/user_data?sslmode=disable")

	if err != nil {
		panic(err)
	}

	fmt.Print("Successfully connected to Postgres DB: {}", pgUrl)

}
