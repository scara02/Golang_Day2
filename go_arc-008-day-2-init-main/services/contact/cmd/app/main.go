package main

import (
	"fmt"
	"log"
    "net/http"

	"github.com/gorilla/mux"

	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/repository"
	"architecture_go/services/contact/internal/usecase"
	"architecture_go/services/contact/internal/delivery"
)

func main() {
	fmt.Println("Hello World!")

	db, err := postgres.DBconnection("postgres", "1234", "localhost", "5432", "go")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer db.Close()

	contactRepository := repository.NewContactRepositoryImpl()
    contactUseCase := usecase.NewUseCase(contactRepository)
    contactDelivery := delivery.NewDelivery(contactUseCase)

    router := mux.NewRouter()

    router.HandleFunc("/contacts", contactDelivery.ContactHandler.CreateContactHandler).Methods("POST")
    router.HandleFunc("/contacts/{id}", contactDelivery.ContactHandler.GetContactHandler).Methods("GET")

    http.Handle("/", router)
    http.ListenAndServe(":3000", nil)

}
