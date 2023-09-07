package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/FianGumilar/restful-api-echo/config"
	_ "github.com/lib/pq"
)

var DbPg *sql.DB

func GetDbPostgres(conf *config.AppConfig) *sql.DB {
	var err error

	dsn := fmt.Sprintf(
		"host=%s "+
			"port=%s "+
			"user=%s "+
			"password=%s "+
			"dbname=%s "+
			conf.Database.Host,
		conf.Database.Port,
		conf.Database.User,
		conf.Database.Pass,
		conf.Database.Name,
	)

	DbPg, err = sql.Open("postgres", dsn)

	err = DbPg.Ping()

	if err != nil {
		log.Printf("Failed connect to database: %s", err)
	}
	return DbPg
}
