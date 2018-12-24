package config

type key string

// DBKey ...
const DBKey key = "DBKey"

var (
	/*
		list of server configuration
	*/

	// Port ...
	Port = "8080"

	/*
		list of db configuration
	*/

	// DriverDB ...
	DriverDB = []string{"postgres"}
	// DBSource ...
	DBSource = map[string]map[string]map[string]string{
		// TODO : can add multiple DB source in here, for configuration. And just put the driver at the top variable
		"postgres": {
			"type": {
				"value": "sql",
			},
			"env": {
				"HOST": "DB_PG_HOST",
				"PORT": "DB_PG_PORT",
				"NAME": "DB_PG_NAME",
				"USER": "DB_PG_USER",
				"PASS": "DB_PG_PASS",
			},
			"dbSource": {
				"development": "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
				"production":  "host=%s port=%s user=%s password=%s dbname=%s TimeZone=Asia/Jakart",
			},
		},
	}

	/*

		client secret

	*/

	// CLIENTSECRET ...
	CLIENTSECRET = "plastik"
)
