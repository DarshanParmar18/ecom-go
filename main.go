package main

import (
	"log"

	"github.com/darshanparmar18/ecom-go/cmd/api"
)

func main() {

	server := api.NewAPIServer(":8080",nil)

	if err := server.Run(); err != nil{
		log.Fatal(err)
	}

}