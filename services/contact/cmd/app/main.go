package main

import (
	"ass3/pkg/store/postgres"
	"ass3/services/contact/internal"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
	db := postgres.DBconnection("localhost", "postgres", "sultan", "5432", "Company")

	fmt.Println("DB connection: ", db)

	defer db.Close()
	contactRepository := internal.NewContactRepository()
	contactUseCase := internal.NewContactUseCase(contactRepository)
	contactDelivery := internal.NewContactDelivery(contactUseCase)
	router := mux.NewRouter()

	router.HandleFunc("/contacts", contactDelivery.CreateContactHandler).Methods("POST")
	router.HandleFunc("/contacts/{id}", contactDelivery.ReadContactHandler).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)

}
