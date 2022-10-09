package services

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/niroopreddym/xm-exercise/pkg/mocks"
	"github.com/niroopreddym/xm-exercise/pkg/models"
	"github.com/stretchr/testify/assert"
)

func Test_PostCompanyDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dbMock := mocks.NewMockDbIface(controller)
	dbMock.EXPECT().DbExecuteScalarReturningID(gomock.Any(), gomock.Any()).AnyTimes().Return(1, nil)
	dbMock.EXPECT().DbClose().AnyTimes()
	service := DatabaseService{
		DatabaseService: dbMock,
	}

	companyDetails := &models.Company{
		Name:    "TestCompany",
		Code:    "cp-01",
		Country: "Cyprus",
		Website: "www.google.com",
		Phone:   "9884377883",
	}

	id, err := service.CreateCompany(companyDetails)
	assert.Nil(t, err)
	assert.NotNil(t, id, 1)

}
