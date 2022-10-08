package handlers

import "net/http"

//HandlerIface provides the interface for handling the product related
type HandlerIface interface {
	PostCompany(w http.ResponseWriter, r *http.Request)
	ListAllCompanies(w http.ResponseWriter, r *http.Request)
	GetCompanyDetails(w http.ResponseWriter, r *http.Request)
	PutCompanyDetails(w http.ResponseWriter, r *http.Request)
	DeleteCompanyDetails(w http.ResponseWriter, r *http.Request)
}
