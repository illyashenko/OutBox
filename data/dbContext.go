package data

import (
	"OutBox/appconfig"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	once    sync.Once
	context *sql.DB
)

func PostgreSqlContext(config appconfig.Postgres) *sql.DB {
	if context == nil {
		once.Do(
			func() {
				cnx, err := sql.Open("postgres", config.ConnectionString)
				if err != nil {
					log.Fatal("Error in connect DataBase", err)
				}
				context = cnx
			},
		)
	}
	return context
}
