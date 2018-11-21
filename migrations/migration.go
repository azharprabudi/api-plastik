package migrations

import (
	"strconv"

	"github.com/api-plastik/db"
	helpers "github.com/api-plastik/helpers/db"
	dbModel "github.com/api-plastik/helpers/db/model"
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
	err := migratePgSQL(db.PgSQL)
	if err != nil {
		panic(err)
	}
}

// running migration pgsql
func migratePgSQL(sql *sqlx.DB) error {
	c := helpers.CreateTrx(sql)
	err := helpers.RunTrx(c, func(tx *sqlx.Tx) error {
		initVerTblPgSQL(tx)

		// get current version table
		currVer := getCurrVer()

		// get old version
		oldVer, err := getOldVer(tx)
		if err != nil {
			return err
		}

		// check version
		if oldVer != currVer {
			// running the query
			err = doMigrateSQL(tx, oldVer, currVer)
			if err != nil {
				return err
			}

			// update curr version
			err = updateVer(tx, currVer)
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

func initVerTblPgSQL(tx *sqlx.Tx) {
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

// get current version
func getCurrVer() int {
	return len(execPgSQL)
}

// get older version sql
func getOldVer(tx *sqlx.Tx) (int, error) {
	// initialize variable
	meta := new(model.Meta)
	where := &dbModel.Condition{
		Key:      "key",
		Operator: "=",
		Value:    "db-version",
		NextCond: "",
	}

	// execute query builder
	query := helpers.QueryWhere("meta", []*dbModel.Condition{where})
	err := tx.QueryRowx(query).StructScan(meta)
	if err != nil {
		return 0, err
	}

	// parse to int
	ver, err := strconv.Atoi(meta.Value)
	if err != nil {
		return 0, err
	}

	// return ver, nil
	return ver, nil
}

// function to running list of migration
func doMigrateSQL(tx *sqlx.Tx, startVer int, untilVer int) error {
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
func updateVer(tx *sqlx.Tx, currVer int) error {
	// parse vers to str
	verStr := strconv.Itoa(currVer)
	meta := model.Meta{
		Key:   "db-version",
		Value: verStr,
	}

	condition := &dbModel.Condition{
		Key:      "key",
		Operator: "=",
		Value:    "db-version",
		NextCond: "",
	}
	query := helpers.UpdateWhere("meta", meta, []*dbModel.Condition{condition})

	// execute the query
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}
	return nil
}
