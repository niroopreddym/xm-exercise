package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/niroopreddym/xm-exercise/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_PostBankDetails(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	serviceMock := mocks.NewMockDatabaseServicesIface(controller)
	serviceMock.EXPECT().CreateCompany(gomock.Any()).AnyTimes().Return(1, nil)

	kafkaMock := mocks.NewMockIKafka(controller)
	kafkaMock.EXPECT().PushToKafkaStream(gomock.Any()).AnyTimes()

	companyHandler := CompaniesHandler{
		CompanyService: serviceMock,
		KafkaService:   kafkaMock,
	}

	jsonData := strings.NewReader(`{
    "name":"TestCompany",
	"code":"cp-01",
	"country":"Cyprus",
	"website":"www.google.com",
	"phone":"9884377883"
}`)

	req, err := http.NewRequest("GET", "localhost:9294", jsonData)
	assert.Nil(t, err)

	res := httptest.NewRecorder()
	posthandler := http.HandlerFunc(companyHandler.PostCompany)
	posthandler.ServeHTTP(res, req)
	assert.Equal(t, res.Code, 200)
}
