package main

import (
	"database/sql"
	"log"

	"github.com/darshanparmar18/ecom/cmd/api"
	"github.com/darshanparmar18/ecom/config"
	"github.com/darshanparmar18/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {

	db, err := db.MyNewSQLStorage(mysql.Config{
		User: config.Envs.DBUser,
		Passwd: config.Envs.DBPassword,
		Addr: config.Envs.DBAddress,
		DBName: config.Envs.DBName,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080",db)

	server.Run()
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil{
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}