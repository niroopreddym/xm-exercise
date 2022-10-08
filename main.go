package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/niroopreddym/xm-exercise/handlers"
	"github.com/niroopreddym/xm-exercise/services"
)

func main() {
	dbInstance := services.NewDatabaseServicesInstance()

	router := mux.NewRouter()

	handler := handlers.NewCompaniesHandler(dbInstance)
	fmt.Println("started listening on port : ", 9294)
	router.Handle("/companies", http.HandlerFunc(handler.PostCompany)).Methods("POST")
	router.Handle("/companies", http.HandlerFunc(handler.ListAllCompanies)).Methods("GET")
	router.Handle("/companies/{company_id}", http.HandlerFunc(handler.GetCompanyDetails)).Methods("GET")

	http.ListenAndServe(":9294", router)
}
