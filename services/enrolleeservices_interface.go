package services

import "github.com/niroopreddym/xm-exercise/models"

//DatabaseServicesIface exposes the public methods to be used by the handler
type DatabaseServicesIface interface {
	CreateCompany(company *models.Company) (int, error)
	GetListOfAllCompanies(productID int) (*models.Company, error)
	GetCompanyDetails(productID int, noOfProductsBooked int) error
}
