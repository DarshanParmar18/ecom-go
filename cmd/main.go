package main

import (
	"log"

	"github.com/darshanparmar18/ecom/cmd/api"
	"github.com/darshanparmar18/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "asd",
		Addr:                 "127.0.1:3306",
		DBName:               "ecom",
		Net:                  "",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}
