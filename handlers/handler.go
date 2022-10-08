package handlers

import (
	"net/http"

	"github.com/niroopreddym/xm-exercise/services"
)

//CompaniesHandler is the class implementation for HandlerIface Interface
type CompaniesHandler struct {
	CompanyService services.DatabaseServicesIface
}

//NewCompaniesHandler instantiates the struct
func NewCompaniesHandler(dbService *services.DatabaseService) HandlerIface {
	return &CompaniesHandler{
		CompanyService: dbService,
	}
}

//PostCompany creates the company data
func (handler *CompaniesHandler) PostCompany(w http.ResponseWriter, r *http.Request) {

}

//ListAllCompanies list details about all companies
func (handler *CompaniesHandler) ListAllCompanies(w http.ResponseWriter, r *http.Request) {

}

//GetCompanyDetails gets the company details based on filter
func (handler *CompaniesHandler) GetCompanyDetails(w http.ResponseWriter, r *http.Request) {

}
