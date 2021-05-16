package database

import (
	"database/sql"
	"fmt"
)


//to get the value os.Getenv("host"), run 'export host=localhost'
const (
	host     = "hpe.c5p6cfuydxbj.us-east-1.rds.amazonaws.com"
	port     = 5432
	user     = "postgres"
	password = "5Peht4VxUI2V1VkBUo01"
	dbname   = "postgres"
)

func StartDatabase() *sql.DB {
	
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}