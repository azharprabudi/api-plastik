package migrations

import (
	"strconv"

	"github.com/azharprabudi/api-plastik/db"
	qb "github.com/azharprabudi/api-plastik/helper/querybuilder"
	qbmodel "github.com/azharprabudi/api-plastik/helper/querybuilder/model"
	trx "github.com/azharprabudi/api-plastik/helper/transaction"
	execsql "github.com/azharprabudi/api-plastik/migrations/exec-sql"
	"github.com/azharprabudi/api-plastik/migrations/model"
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
	t := trx.NewTransaction()
	c := t.CreateTrx(sql)
	err := t.RunTrx(c, func(tx *sqlx.Tx) error {
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
			doMigrateSQL(tx, oldVer, currVer)

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

func initVerTblPgSQL(tx *sqlx.Tx) bool {
	/* create table version pg */
	sql := `
	CREATE TABLE IF NOT EXISTS  "meta"(
		"key" varchar(100) NOT NULL,
		"value" varchar(100) NOT NULL,
		CONSTRAINT info_pk PRIMARY KEY ("key")
	);
	INSERT INTO "meta"(key, value)
		VALUES ('db-version', '0')
		ON CONFLICT DO NOTHING;
	`
	_, err := tx.Exec(sql)
	if err != nil {
		return true
	}
	return false
}

// get current version
func getCurrVer() int {
	return len(execPgSQL)
}

// get older version sql
func getOldVer(tx *sqlx.Tx) (int, error) {
	// initialize variable
	meta := new(model.Meta)
	where := &qbmodel.Condition{
		Key:      "key",
		Operator: "=",
		Value:    "db-version",
		NextCond: "",
	}

	// execute query builder
	q := qb.NewQueryBuilder()
	query := q.QueryWhere("meta", []*qbmodel.Condition{where}, nil)

	// check the previous table already exists or new to created
	err := tx.QueryRowx(query).StructScan(meta)

	// check error
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
func doMigrateSQL(tx *sqlx.Tx, startVer int, untilVer int) {
	for i := startVer; i < untilVer; i++ {
		_, err := tx.Exec(execPgSQL[i])
		if err != nil {
			panic(err)
		}
	}
}

// update new version db migrations
func updateVer(tx *sqlx.Tx, currVer int) error {
	// parse vers to str
	verStr := strconv.Itoa(currVer)
	meta := model.Meta{
		Key:   "db-version",
		Value: verStr,
	}

	condition := &qbmodel.Condition{
		Key:      "key",
		Operator: "=",
		Value:    "db-version",
		NextCond: "",
	}

	// execute query
	q := qb.NewQueryBuilder()
	query := q.UpdateWhere("meta", meta, []*qbmodel.Condition{condition})

	// execute the query
	_, err := tx.Exec(query, meta.Key, meta.Value)

	if err != nil {
		return err
	}
	return nil
}
