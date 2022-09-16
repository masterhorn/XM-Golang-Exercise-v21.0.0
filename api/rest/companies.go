package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/gorilla/mux"
	"github.com/masterhorn/XM-Golang-Exercise-v21.0.0/cfg"
	database "github.com/masterhorn/XM-Golang-Exercise-v21.0.0/db"
)

func ListCompanies(w http.ResponseWriter, r *http.Request) {
	db := database.GetDb(cfg.Config.DbConnectionString, cfg.Config.DbUser, cfg.Config.DbPassword, cfg.Config.DbName)
	defer db.Close()

	var company []database.CompanyModel

	err := db.Model(&company).Returning("*").Select()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to get list of companies")
		return
	}

	err = cfg.Ren.JSON(w, http.StatusOK, company)
	if err != nil {
		log.Printf("JSON rendering failed: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := database.GetDb(cfg.Config.DbConnectionString, cfg.Config.DbUser, cfg.Config.DbPassword, cfg.Config.DbName)
	defer db.Close()

	companyID := vars["id"]

	var company database.CompanyModel

	err := db.Model(&company).Returning("*").Where("id = ?", companyID).Select()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get company %s: %s\n", companyID, err)
		code := http.StatusInternalServerError
		if err == pg.ErrNoRows {
			errorMessage = fmt.Sprintln("No company with id:", companyID)
			code = http.StatusNotFound
		}
		http.Error(w, errorMessage, code)
		log.Println(errorMessage)
		return
	}

	err = cfg.Ren.JSON(w, http.StatusOK, company)
	if err != nil {
		log.Printf("JSON rendering failed: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var payload CreateCompanyPayload
	db := database.GetDb(cfg.Config.DbConnectionString, cfg.Config.DbUser, cfg.Config.DbPassword, cfg.Config.DbName)
	defer db.Close()

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read JSON: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(requestBody, &payload)
	if err != nil {
		log.Printf("Error unmarshaling: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	company := database.CompanyModel{
		Name:    payload.Name,
		Code:    payload.Code,
		Country: payload.Country,
		Website: payload.Website,
		Phone:   payload.Phone,
	}

	_, err = db.Model(&company).Returning("*").Insert()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to create a company. %s", err)
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), `violates foreign key constraint "country_id"`) {
			errorMessage = fmt.Sprintf("Invalid station_id. %s", err)
			code = http.StatusBadRequest
		}
		log.Print(errorMessage)
		http.Error(w, errorMessage, code)
		return
	}

	err = cfg.Ren.JSON(w, http.StatusOK, company)
	if err != nil {
		log.Printf("JSON rendering failed: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var payload UpdateCompanyPayload
	db := database.GetDb(cfg.Config.DbConnectionString, cfg.Config.DbUser, cfg.Config.DbPassword, cfg.Config.DbName)
	defer db.Close()

	companyID := vars["id"]

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read JSON: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(requestBody, &payload)
	if err != nil {
		log.Printf("Error unmarshaling: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var company database.CompanyModel

	err = db.Model(&company).Returning("*").Where("id = ?", companyID).Select()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get company %s: %s\n", companyID, err)
		code := http.StatusInternalServerError
		if err == pg.ErrNoRows {
			errorMessage = fmt.Sprintln("No company with id:", companyID)
			code = http.StatusNotFound
		}
		http.Error(w, errorMessage, code)
		log.Println(errorMessage)
		return
	}

	if payload.Name != "" {
		company.Name = payload.Name
	}
	if payload.Code != "" {
		company.Code = payload.Code
	}
	if payload.Website != "" {
		company.Website = payload.Website
	}
	if payload.Phone != "" {
		company.Phone = payload.Phone
	}

	_, err = db.Model(&company).WherePK().Update()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to update a company. %s", err)
		code := http.StatusInternalServerError
		if strings.Contains(err.Error(), `violates foreign key constraint "country_id"`) {
			errorMessage = fmt.Sprintf("Invalid station_id. %s", err)
			code = http.StatusBadRequest
		}
		log.Print(errorMessage)
		http.Error(w, errorMessage, code)
		return
	}

	err = cfg.Ren.JSON(w, http.StatusOK, company)
	if err != nil {
		log.Printf("JSON rendering failed: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := database.GetDb(cfg.Config.DbConnectionString, cfg.Config.DbUser, cfg.Config.DbPassword, cfg.Config.DbName)
	defer db.Close()

	companyID := vars["id"]

	var company database.CompanyModel

	result, err := db.Model(&company).Where("id = ?", companyID).Delete()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to delete company %s: %s\n", companyID, err)
		log.Println(errorMessage)
		http.Error(w, errorMessage, http.StatusInternalServerError)
		return
	}
	if result.RowsAffected() == 0 {
		http.Error(w, fmt.Sprintln("No company with ID:", companyID), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
