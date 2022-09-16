package helpers

import (
	"fmt"

	"github.com/masterhorn/XM-Golang-Exercise-v21.0.0/api/rest"
)

func GenerateCompanyCreatePayload() rest.CreateCompanyPayload {
	companyName := fmt.Sprintf(`%s_%s`, "autotest", GenerateString(6))
	companyCode := fmt.Sprintf(`%s_%s`, "autotest", GenerateString(6))
	website := fmt.Sprintf(`%s_%s`, "autotest", GenerateString(6))
	country := fmt.Sprintf(`%s_%s`, "autotest", GenerateString(6))
	phone := GenerateString(6)

	companyPayload := rest.CreateCompanyPayload{
		Name:    companyName,
		Code:    companyCode,
		Country: country,
		Website: website,
		Phone:   phone,
	}
	return companyPayload
}
