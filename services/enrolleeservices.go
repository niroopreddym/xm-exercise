package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/niroopreddym/xm-exercise/database"
	"github.com/niroopreddym/xm-exercise/models"
)

//DatabaseService is the class implementation for ProductServicesIface interface
type DatabaseService struct {
	DatabaseService database.DbIface
}

//NewDatabaseServicesInstance instantiates the struct
func NewDatabaseServicesInstance() *DatabaseService {
	return &DatabaseService{
		DatabaseService: database.DBNewHandler(),
	}
}

//GetListOfAllAvailableProducts get list of all available products
func (service *DatabaseService) GetListOfAllAvailableProducts(productID int) (*models.Company, error) {
	defer service.DatabaseService.DbClose()
	query := "select * from public.products where id = " + strconv.Itoa(productID)
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	var txResult *models.Company
	for rowsAffected.Next() {
		var id int
		var productName string
		var availableQuantity int
		if err := rowsAffected.Scan(&id, &productName, &availableQuantity); err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		txResult = &models.Company{
			ID: id,
		}
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, errors.New("internal server error")
	}

	return txResult, nil
}

//UpdateProductsCountToInventory adds products to the inventory
func (service *DatabaseService) UpdateProductsCountToInventory(product models.Company) (int, error) {
	query := "select AddProductsQuantity(" + strconv.Itoa(product.ID) + "," + strconv.Itoa(1) + ")"

	tx, err := service.DatabaseService.TxBegin()
	_, err = service.DatabaseService.TxExecuteStmt(tx, query)

	defer service.DatabaseService.DbClose()
	if err != nil {
		return int(http.StatusInternalServerError), errors.New("internal server error")
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return int(http.StatusInternalServerError), errors.New("internal server error")
	}

	return int(http.StatusOK), nil
}

//BookProducts adds products to the inventory
func (service *DatabaseService) BookProducts(productID int, noOfProductsBooked int) error {
	defer service.DatabaseService.DbClose()
	query := "select BookProducts(" + strconv.Itoa(productID) + "," + strconv.Itoa(noOfProductsBooked) + ")"

	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		return errors.New("internal server error")
	}
	var txResult *bool
	for rowsAffected.Next() {

		if err := rowsAffected.Scan(&txResult); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Value: %t\n", *txResult)
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return errors.New("internal server error")
	}

	if *txResult == false {
		return errors.New("available quantity for the product is less than the requested quantity")
	}

	return nil
}