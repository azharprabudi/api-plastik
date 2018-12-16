package config

type key string

// DBKey ...
const DBKey key = "DBKey"

var (
	/*
		list of server configuration
	*/

	// Mode ...
	Mode = "development"
	// Port ...
	Port = ":3000"

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
				"value": "host=%s port=%s user=%s password=%s dbname=%s",
			},
		},
	}

	/*

		client secret

	*/

	// CLIENTSECRET ...
	CLIENTSECRET = "plastik"
)
