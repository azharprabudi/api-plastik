package db

import (
	"fmt"
	"os"

	"github.com/api-plastik/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// OpenConnectionDB ...
func OpenConnectionDB() *DB {
	db := new(DB)
	for _, driver := range config.DriverDB {
		if conf, isExists := config.DBSource[driver]; isExists {
			if conf["type"]["value"] == "sql" {
				sql, err := connectDBTypeSQL(driver, conf["dbSource"]["value"], conf["env"])
				// check error message
				if err != nil {
					panic(err)
				}

				// assign value db
				if driver == "postgres" {
					db.PgSQL = sql
				}

			}
		}
	}
	return db
}

// open connection sql
func connectDBTypeSQL(driver string, dbSource string, env map[string]string) (*sqlx.DB, error) {
	// get information connection info
	dbInfo := fmt.Sprintf(dbSource, getENV(env["HOST"]), getENV(env["PORT"]), getENV(env["USER"]), getENV(env["PASS"]), getENV(env["NAME"]))

	// connect db
	db, err := sqlx.Connect(driver, dbInfo)
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
