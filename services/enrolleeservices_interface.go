package services

import "github.com/niroopreddym/xm-exercise/models"

//DatabaseServicesIface exposes the public methods to be used by the handler
type DatabaseServicesIface interface {
	GetListOfAllAvailableProducts(productID int) (*models.Company, error)
	UpdateProductsCountToInventory(product models.Company) (int, error)
	BookProducts(productID int, noOfProductsBooked int) error
}
