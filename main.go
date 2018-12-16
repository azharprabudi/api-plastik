package main

import (
	"github.com/azharprabudi/api-plastik/cmd"
	"github.com/azharprabudi/api-plastik/db"
	"github.com/azharprabudi/api-plastik/migrations"
	"github.com/azharprabudi/api-plastik/router"
)

func main() {
	/* open db connection */
	db := db.OpenConnectionDB()

	/* running migration */
	migrations.RunMigration(db)

	/* create routing, and running all routing */
	r := router.InitRoute(db)

	/* running server */
	cmd.StartServer(r)
}
