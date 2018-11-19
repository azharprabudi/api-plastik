package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// OpenConnectionDB ...
func OpenConnectionDB() *DB {
	db := new(DB)
	pgSQL, err := connectToPgSQL()
	if err != nil {
		panic(err)
	}
	// add pgsql to db connection list
	db.PgSQL = pgSQL
	return db
}

// open connection pgsql
func connectToPgSQL() (*sqlx.DB, error) {
	// get information connection info
	dbInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", getENV("DB_PG_HOST"), getENV("DB_PG_PORT"), getENV("DB_PG_USER"), getENV("DB_PG_NAME"))

	// connect db
	db, err := sqlx.Connect("postgres", dbInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// get env variables
func getENV(key string) string {
	val := os.Getenv(key)
	return val
}
