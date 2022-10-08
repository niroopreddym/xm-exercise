package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/niroopreddym/xm-exercise/models"
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
	companyDetails := models.Company{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err := json.Unmarshal([]byte(strBufferValue), &companyDetails)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	errorMessages := []string{}
	postRequestBodyInitialValidation(companyDetails, &errorMessages)
	if len(errorMessages) > 0 {
		responseController(w, http.StatusBadRequest, errorMessages)
		return
	}

	uniqueID, err := handler.CompanyService.CreateCompany(&companyDetails)
	if err != nil {
		fmt.Println(err)
		responseController(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseController(w, http.StatusOK, map[string]string{
		"companyID": fmt.Sprintf("%v", uniqueID),
	})
}

//ListAllCompanies list details about all companies
func (handler *CompaniesHandler) ListAllCompanies(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	xforward := r.Header.Get("X-Forwarded-For")
	fmt.Println("IP : ", ip)
	fmt.Println("X-Forwarded-For : ", xforward)

	lstCompanies, err := handler.CompanyService.GetListOfAllCompanies()
	if err != nil {
		fmt.Println(err)
		responseController(w, http.StatusInternalServerError, "Error occured while fetching the company details")
		return
	}

	responseController(w, http.StatusOK, lstCompanies)
}

//GetCompanyDetails gets the company details based on filter
func (handler *CompaniesHandler) GetCompanyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["company_id"]

	companyID, err := strconv.Atoi(id)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	companyDetails, err := handler.CompanyService.GetCompanyDetails(companyID)
	if err != nil {
		fmt.Println(err)
		responseController(w, http.StatusInternalServerError, "Error occured while fetching the company details")
		return
	}

	responseController(w, http.StatusOK, companyDetails)
}

//PutCompanyDetails ...
func (handler *CompaniesHandler) PutCompanyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["company_id"]

	companyID, err := strconv.Atoi(id)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}
	companyDetails := models.Company{}
	bodyBytes, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		responseController(w, http.StatusInternalServerError, readErr)
		return
	}

	strBufferValue := string(bodyBytes)
	err = json.Unmarshal([]byte(strBufferValue), &companyDetails)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	err = handler.CompanyService.PutCompanyDetails(companyID, &companyDetails)
	if err != nil {
		fmt.Println(err)
		responseController(w, http.StatusInternalServerError, "Error occured while updating the company details")
		return
	}

	responseController(w, http.StatusNoContent, "")
}

//DeleteCompanyDetails ....
func (handler *CompaniesHandler) DeleteCompanyDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["company_id"]

	companyID, err := strconv.Atoi(id)
	if err != nil {
		responseController(w, http.StatusInternalServerError, err)
		return
	}

	err = handler.CompanyService.DeleteCompanyDetails(companyID)
	if err != nil {
		fmt.Println(err)
		responseController(w, http.StatusInternalServerError, "Error occured while deleting the company details")
		return
	}

	responseController(w, http.StatusNoContent, "")
}

func responseController(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func postRequestBodyInitialValidation(companyDetails models.Company, errorMessages *[]string) {
	if companyDetails.Name == "" {
		errorMessage := "Attribute Missing: Name in the request body"
		*errorMessages = append(*errorMessages, errorMessage)
	}
}
