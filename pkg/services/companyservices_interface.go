package services

import "github.com/niroopreddym/xm-exercise/pkg/models"

//DatabaseServicesIface exposes the public methods to be used by the handler
type DatabaseServicesIface interface {
	CreateCompany(company *models.Company) (int, error)
	GetListOfAllCompanies() ([]*models.Company, error)
	GetCompanyDetails(companyID int) (*models.Company, error)
	PutCompanyDetails(companyID int, company *models.Company) error
	DeleteCompanyDetails(companyID int) error
}
