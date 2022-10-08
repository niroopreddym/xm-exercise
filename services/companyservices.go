package services

import (
	"errors"
	"fmt"
	"log"

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

//CreateCompany ...
func (service *DatabaseService) CreateCompany(company *models.Company) (int, error) {
	defer service.DatabaseService.DbClose()

	insertStat := `insert into Company(name, code, country, website, phone) select $1, $2, $3, $4, $5 returning id`

	companyID, err := service.DatabaseService.DbExecuteScalarReturningID(insertStat, company.Name, company.Code, company.Country, company.Website, company.Phone)

	if err != nil {
		log.Println(err)
		return 0, err
	}
	return companyID, nil
}

//GetListOfAllCompanies get list of all available products
func (service *DatabaseService) GetListOfAllCompanies() ([]*models.Company, error) {
	defer service.DatabaseService.DbClose()
	query := "select * from Company"
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	var txResult []*models.Company
	for rowsAffected.Next() {
		var id int
		var name string
		var code string
		var country string
		var website string
		var phone string
		if err := rowsAffected.Scan(&id, &name, &code, &country, &website, &phone); err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		company := &models.Company{
			ID:      id,
			Name:    name,
			Code:    code,
			Country: country,
			Website: website,
			Phone:   phone,
		}

		txResult = append(txResult, company)
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, errors.New("internal server error")
	}

	return txResult, nil
}

//GetCompanyDetails adds products to the inventory
func (service *DatabaseService) GetCompanyDetails(companyID int) (*models.Company, error) {
	defer service.DatabaseService.DbClose()
	query := "select * from Company where id=" + fmt.Sprint(companyID)
	tx, err := service.DatabaseService.TxBegin()
	rowsAffected, err := service.DatabaseService.TxQuery(tx, query)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	var txResult *models.Company

	for rowsAffected.Next() {
		var id int
		var name string
		var code string
		var country string
		var website string
		var phone string
		if err := rowsAffected.Scan(&id, &name, &code, &country, &website, &phone); err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		txResult = &models.Company{
			ID:      id,
			Name:    name,
			Code:    code,
			Country: country,
			Website: website,
			Phone:   phone,
		}
	}

	if err = service.DatabaseService.TxComplete(tx); err != nil {
		return nil, errors.New("internal server error")
	}

	return txResult, nil
}

//PutCompanyDetails ...
func (service *DatabaseService) PutCompanyDetails(companyID int, company *models.Company) error {
	putStat := `update Company set name=$1, code=$2, country=$3, website=$4, phone=$5 where id=` + fmt.Sprint(companyID)
	_, err := service.DatabaseService.DbExecuteScalar(putStat, company.Name, company.Code, company.Country, company.Website, company.Phone)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//DeleteCompanyDetails ...
func (service *DatabaseService) DeleteCompanyDetails(companyID int) error {
	deleteStat := `Delete from Company where id=` + fmt.Sprint(companyID)
	_, err := service.DatabaseService.DbExecuteScalar(deleteStat)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
