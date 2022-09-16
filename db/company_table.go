package db

type CompanyModel struct {
	tableName struct{} `pg:"company2"`
	ID        int      `json:"id" pg:"id"`
	Name      string   `json:"name"`
	Code      string   `json:"code"`
	Country   string   `json:"country"`
	Website   string   `json:"website"`
	Phone     string   `json:"phone"`
}
