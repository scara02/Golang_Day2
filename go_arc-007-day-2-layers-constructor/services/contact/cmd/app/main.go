package main

import (
	"fmt"
	"log"

	"architecture_go/pkg/store/postgres"
)

func main() {
	fmt.Println("Hello World!")

	db, err := postgres.DBconnection("postgres", "1234", "localhost", "5432", "go")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer db.Close()

}
