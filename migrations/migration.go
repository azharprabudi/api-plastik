package migrations

import (
	"strconv"

	"github.com/api-plastik/db"
	helpers "github.com/api-plastik/helpers/db"
	execsql "github.com/api-plastik/migrations/exec-sql"
	"github.com/api-plastik/migrations/model"
	"github.com/jmoiron/sqlx"
)

var execPgSQL = []string{
	execsql.First,
}

// RunMigration ...
func RunMigration(db *db.DB) {
	/* migration postgresql */
	err := migrationPgSQL(db.PgSQL)
	if err != nil {
		panic(err)
	}
}

// running migration pgsql
func migrationPgSQL(sql *sqlx.DB) error {
	c := helpers.StartTransaction(sql)
	err := helpers.RunTransaction(c, func(tx *sqlx.Tx) error {
		// initialize version table
		initVersionTablePgSQL(tx)

		// get older version table
		oldVer, err := getOldVersionPgSQL(tx)
		if err != nil {
			return err
		}
		// get current version table
		currVer := currVersionPgSQL()

		if oldVer != currVer {
			// running the query
			err = doMigrationSQL(tx, oldVer, currVer)
			if err != nil {
				return err
			}

			// update curr version
			err = updateVersion(tx, currVer)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func initVersionTablePgSQL(tx *sqlx.Tx) {
	/* create table version pg */
	sql := `
	CREATE TABLE "meta"(
		"key" varchar(100) NOT NULL,
		"value" varchar(100) NOT NULL,
		CONSTRAINT info_pk PRIMARY KEY ("key")
	);
	INSERT INTO "meta"(key, value)
		VALUES ('db-version', '0')
		ON CONFLICT DO NOTHING;
	`
	tx.Exec(sql)
}

// get older version sql
func getOldVersionPgSQL(tx *sqlx.Tx) (int, error) {

	meta := new(model.Meta)
	rowQuery := tx.QueryRowx("select * from meta where key=$1", "db-version")

	// get results
	err := rowQuery.StructScan(meta)
	if err != nil {
		return 0, err
	}

	versionToInt, err := strconv.Atoi(meta.Value)
	if err != nil {
		return 0, err
	}
	return versionToInt, nil
}

// function to running list of migration
func doMigrationSQL(tx *sqlx.Tx, startVer int, untilVer int) error {
	var err error
	for i := startVer; i < untilVer; i++ {
		_, err = tx.Exec(execPgSQL[i])
		if err != nil {
			break
		}
	}
	return err
}

// update new version db migrations
func updateVersion(tx *sqlx.Tx, currVer int) error {
	_, err := tx.Exec("update meta set value=$1 where key='db-version'", currVer)
	if err != nil {
		return err
	}
	return nil
}

// get current version from array at the top of file
func currVersionPgSQL() int {
	return len(execPgSQL)
}
