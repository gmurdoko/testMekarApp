package config

import (
	"database/sql"
	"fmt"
	"log"
	"testMekarApp/utils"

	//for connection mysql db
	_ "github.com/go-sql-driver/mysql"
)

func ReadEnvironmentConnection() *sql.DB {
	dbEngine := utils.ViperGetEnvironment("DB_ENGINE", "mysql")
	dbUser := utils.ViperGetEnvironment("DB_USER", "root") //root
	dbPassword := utils.ViperGetEnvironment("DB_PASSWORD", "")
	dbHost := utils.ViperGetEnvironment("DB_HOST", "localhost") //localhost
	dbPort := utils.ViperGetEnvironment("DB_PORT", "3306")      //3306
	dbSchema := utils.ViperGetEnvironment("DB_SCHEMA", "db_test_mekar")

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	db, err := sql.Open(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
