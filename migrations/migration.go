package migrations

import (
	"github.com/api-plastik/db"
	helpers "github.com/api-plastik/helpers/db"
	execsql "github.com/api-plastik/migrations/exec-sql"
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
		checkVersionPgSQL(tx)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// check version db
func checkVersionPgSQL(tx *sqlx.Tx) {

}

// get current version from array at the top of file
func newVersionPgSQL() int {
	return len(execPgSQL)
}
