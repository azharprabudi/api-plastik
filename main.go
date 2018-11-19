package main

import (
	"github.com/api-plastik/config"
	"github.com/api-plastik/db"
	"github.com/api-plastik/route"
)

func main() {
	/* open db connection */
	db := db.OpenConnectionDB()

	/* running migration */

	/* create routing, and running all routing */
	r := route.InitRoute(db)

	/* running server */
	port := ":3000"
	config.StartServer(port, r)
}
