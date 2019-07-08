package company

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/database"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/make_result"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/model"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func topCompaniesHandler(w http.ResponseWriter, r *http.Request) {

	top := r.URL.Query().Get("top")
	countryName := r.URL.Query().Get("country_name")

	orderBy := "valuation desc"
	var limit int
	if top == "" {
		limit = 1
	} else {
		var err error
		limit, err = strconv.Atoi(top)
		if err != nil {
			make_result.Result(w, r, http.StatusBadRequest, err)
			return
		}
	}

	var companies []model.Company
	query := database.Engine
	if countryName != "" {
		var country model.Country
		database.Engine.Where("name like ?", fmt.Sprintf("%%%s%%", countryName)).First(&country)
		if country.ID != 0 {
			query = query.Where("country_id = ?", country.ID)
		}
	}
	if dbError := query.Order(orderBy).Limit(limit).Find(&companies).Error; dbError != nil {
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}

	var companiesSerializer []model.CompanySerializer
	for _, i := range companies {
		companiesSerializer = append(companiesSerializer, i.Serializer(database.Engine))
	}
	make_result.Result(w, r, http.StatusOK, companiesSerializer)
}

type Sum struct {
	CountryName string `json:"country_name"`
	Valuation   string `json:"valuation"`
}

type Scan struct {
	Value uint `json:"value"`
}

func sumCompaniesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryName := vars["country_name"]

	query := database.Engine
	var country model.Country
	if countryName != "all" {
		if dbError := query.Where("name like ?", countryName).First(&country).Error; dbError != nil {
			make_result.Result(w, r, http.StatusBadRequest, dbError)
			return
		}
		query = query.Where("country_id = ?", country.ID)
	}
	var sumScan Scan
	query.Raw("sum(valuation) as value").Scan(&sumScan)

	var result Sum
	if country.ID != 0 {
		result.CountryName = country.Name
	} else {
		result.CountryName = "All"
	}
	v := strconv.FormatFloat(float64(sumScan.Value)/float64(math.Pow(10, 8)), 'f', 0, 32)
	result.Valuation = fmt.Sprintf("%s亿", v)
	make_result.Result(w, r, http.StatusOK, result)

}

func oneCompanyHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var param oneCompany
	param.Name = name
	log.Println("Param: ", param)
	if err := param.Valid(); err != nil {
		log.Println(err)
		make_result.Result(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var companies []model.Company

	if dbError := database.Engine.Where("name like ?", fmt.Sprintf("%%%s%%", param.Name)).Find(&companies).Error; dbError != nil {
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}

	var companiesSerializer []model.CompanySerializer
	for _, i := range companies {
		companiesSerializer = append(companiesSerializer, i.Serializer(database.Engine))
	}
	make_result.Result(w, r, http.StatusOK, companiesSerializer)

}
