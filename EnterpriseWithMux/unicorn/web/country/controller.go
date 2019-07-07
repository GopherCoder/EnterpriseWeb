package country

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/database"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/make_result"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllCountriesHandler(w http.ResponseWriter, r *http.Request) {
	var countries []model.Country

	if dbError := database.Engine.Find(&countries).Error; dbError != nil {
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}

	var results []model.CountrySerializer
	for _, i := range countries {
		results = append(results, i.Serializer())
	}
	make_result.Result(w, r, http.StatusOK, results)
}

func getOneCountryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryId := vars["country_id"]
	log.Println("countryId :", countryId)
	var country model.Country
	if dbError := database.Engine.Where("id = ?", countryId).First(&country).Error; dbError != nil {
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}
	make_result.Result(w, r, http.StatusOK, country.Serializer())

}

func postOneCountryHandler(w http.ResponseWriter, r *http.Request) {
	var param CreateCountry
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		log.Println(err)
		make_result.Result(w, r, http.StatusBadRequest, err)
		return
	}
	if err := param.Valid(); err != nil {
		log.Println(err)
		make_result.Result(w, r, http.StatusBadRequest, err)
		return
	}
	var country model.Country
	country = model.Country{
		Name: param.Data.Name,
	}
	tx := database.Engine.Begin()
	defer tx.Close()
	if dbError := tx.Save(&country).Error; dbError != nil {
		tx.Rollback()
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}
	tx.Commit()
	make_result.Result(w, r, http.StatusOK, country.Serializer())

}

func deleteOneCountryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryId := vars["country_id"]
	tx := database.Engine.Begin()
	defer tx.Close()

	var country model.Country
	if dbError := tx.Where("id = ?", countryId).First(&country).Error; dbError != nil {
		tx.Rollback()
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}
	tx.Delete(&country)
	tx.Commit()
	make_result.Result(w, r, http.StatusOK, country.Serializer())

}
