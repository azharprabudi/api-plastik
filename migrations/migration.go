package migrations

import (
	"strconv"

	"github.com/api-plastik/db"
	qb "github.com/api-plastik/helper/querybuilder"
	qbModel "github.com/api-plastik/helper/querybuilder/model"
	trx "github.com/api-plastik/helper/transaction"
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
	t := trx.NewTransaction()
	c := t.CreateTrx(sql)
	err := t.RunTrx(c, func(tx *sqlx.Tx) error {
		isAlreadyCreated := initVerTblPgSQL(tx)

		// get current version table
		currVer := getCurrVer()

		// get old version
		oldVer, err := getOldVer(sql, tx, isAlreadyCreated)
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
	CREATE TABLE "meta"(
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
func getOldVer(sql *sqlx.DB, tx *sqlx.Tx, isAlreadyCreated bool) (int, error) {
	// initialize variable
	var err error
	meta := new(model.Meta)
	where := &qbModel.Condition{
		Key:      "key",
		Operator: "=",
		Value:    "db-version",
		NextCond: "",
	}

	// execute query builder
	q := qb.NewQueryBuilder()
	query := q.QueryWhere("meta", []*qbModel.Condition{where})

	// check the previous table already exists or new to created
	if isAlreadyCreated == true {
		err = sql.QueryRowx(query).StructScan(meta)
	} else {
		err = tx.QueryRowx(query).StructScan(meta)
	}

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
		tx.Exec(execPgSQL[i])
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

	condition := &qbModel.Condition{
		Key:      "key",
		Operator: "=",
		Value:    "db-version",
		NextCond: "",
	}

	// execute query
	q := qb.NewQueryBuilder()
	query := q.Update("meta", meta, []*qbModel.Condition{condition})

	// execute the query
	_, err := tx.Exec(query)

	if err != nil {
		return err
	}
	return nil
}
