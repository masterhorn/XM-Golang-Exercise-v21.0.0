package integration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	database "github.com/masterhorn/XM-Golang-Exercise-v21.0.0/db"
	"github.com/masterhorn/XM-Golang-Exercise-v21.0.0/internal"
	"github.com/masterhorn/XM-Golang-Exercise-v21.0.0/tests/helpers"
)

var route = "/company"

func TestCompanies(t *testing.T) {
	handlers := internal.Handlers()
	server := httptest.NewServer(handlers)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	// create company
	companyPayload := helpers.GenerateCompanyCreatePayload()

	createdCompany := e.POST(route).WithJSON(companyPayload).Expect().Status(http.StatusOK).
		JSON().Object().ValueEqual("name", companyPayload.Name)

	// create terminal model struct to further comparison
	company := database.CompanyModel{
		Name:    companyPayload.Name,
		Code:    companyPayload.Code,
		Country: companyPayload.Country,
		Website: companyPayload.Website,
		Phone:   companyPayload.Phone,
	}

	company.ID = int(createdCompany.Value("id").Number().Raw())
	companiesPathAndID := fmt.Sprintf(`%s/%d`, route, company.ID)

	// get list of companies
	e.GET(route).Expect().Status(http.StatusOK).JSON().Array().NotEmpty()

	// get specified company created earlier
	e.GET(companiesPathAndID).Expect().Status(http.StatusOK).JSON().Object().Equal(company)

	updatedPayload := companyPayload
	updatedPayload.Name = fmt.Sprintf(`%s_%s`, "autotest", helpers.GenerateString(6))

	updatedCompany := company
	updatedCompany.Name = updatedPayload.Name

	// update name of specified company created earlier
	e.PUT(companiesPathAndID).WithJSON(updatedPayload).Expect().
		Status(http.StatusOK).JSON().Object().Equal(updatedCompany)

	// delete specified company
	e.DELETE(companiesPathAndID).Expect().Status(http.StatusNoContent)
}
