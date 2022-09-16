package rest

type CreateCompanyPayload struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}

type UpdateCompanyPayload struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"copuntry"`
	Website string `json:"website"`
	Phone   string `json:"phone"`
}
