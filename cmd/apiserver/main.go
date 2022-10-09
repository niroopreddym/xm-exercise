package main

import (
	"fmt"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"

	"github.com/niroopreddym/xm-exercise/pkg/handlers"
	kafka "github.com/niroopreddym/xm-exercise/pkg/kafkaproducer"
	"github.com/niroopreddym/xm-exercise/pkg/services"
)

func main() {
	//key could be saved somewhere safe
	var mySigningKey = []byte("ultimateStarAjith")

	jwtMiddlewareInstance := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	dbInstance := services.NewDatabaseServicesInstance()
	kafkaProducerInstance := kafka.NewXMKafka()
	router := mux.NewRouter()
	handler := handlers.NewCompaniesHandler(dbInstance, kafkaProducerInstance)

	middlewareHandler := handlers.NewMiddlewareHandler()

	fmt.Println("started listening on port : ", 9294)
	router.Handle("/companies", jwtMiddlewareInstance.Handler(middlewareHandler.IPFIlterMiddleware(http.HandlerFunc(handler.PostCompany)))).Methods("POST")
	router.Handle("/companies", http.HandlerFunc(handler.ListAllCompanies)).Methods("GET")
	router.Handle("/companies/{company_id}", http.HandlerFunc(handler.GetCompanyDetails)).Methods("GET")
	router.Handle("/companies/{company_id}", http.HandlerFunc(handler.PutCompanyDetails)).Methods("PUT")
	router.Handle("/companies/{company_id}", jwtMiddlewareInstance.Handler(middlewareHandler.IPFIlterMiddleware(http.HandlerFunc(handler.DeleteCompanyDetails)))).Methods("DELETE")

	router.HandleFunc("/getjwttoken", middlewareHandler.GenerateToken).Methods("GET")

	// muxServerWrapped := getServerConfig(router)
	http.ListenAndServe(":9294", router)
}

// func getServerConfig(router *mux.Router) *http.Server {
// 	apiServer := &http.Server{
// 		Addr:    "127.0.0.1:9294",
// 		Handler: http.HandlerFunc(),
// 	}

// 	return apiServer
// }
