package db

import (
	"os"

	"github.com/azharprabudi/api-plastik/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// OpenConnectionDB ...
func OpenConnectionDB() *DB {
	mode, ok := os.LookupEnv("model")
	if !ok {
		mode = "development"
	}

	db := new(DB)
	for _, driver := range config.DriverDB {
		if conf, isExists := config.DBSource[driver]; isExists {
			if conf["type"]["value"] == "sql" {
				sql, err := connectDBTypeSQL(driver, conf["dbSource"][mode], conf["env"])
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
	// dbInfo := fmt.Sprintf(dbSource, getENV(env["HOST"]), getENV(env["PORT"]), getENV(env["USER"]), getENV(env["PASS"]), getENV(env["NAME"]))

	// connect db
	db, err := sqlx.Connect(driver, "postgres://ihpxomsdizdbqz:9b01d647d7756833bd68eb632403187e4746dec8fd2b4eebb00a47e1d1eba570@ec2-54-243-150-10.compute-1.amazonaws.com:5432/d2b2s4nvijmjd0")
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
