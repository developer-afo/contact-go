package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseFacade *gorm.DB

type DatabaseInterface interface {
	Connection() *gorm.DB
}

type connection struct {
	database *gorm.DB
}

func StartDatabaseClient() DatabaseInterface {

	dsn := "host=localhost port=5432 user=postgres dbname=contact_db password=somepassword"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true})

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Database connection was successfull")

	database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	database.Exec("CREATE EXTENSION IF NOT EXISTS pgcrypto")

	//defer db.Close()

	DatabaseFacade = database

	return &connection{
		database: database,
	}
}

func (conn connection) Connection() *gorm.DB {
	return conn.database
}
