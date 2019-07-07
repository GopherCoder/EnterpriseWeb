package concept

import (
	"EnterpriseWeb/EnterpriseWithMux/unicorn/pkg/database"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/make_result"
	"EnterpriseWeb/EnterpriseWithMux/unicorn/web/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func createOneConceptHandler(w http.ResponseWriter, r *http.Request) {
	var param CreateConcept
	if err := json.NewDecoder(r.Body).Decode(&param); err != nil {
		make_result.Result(w, r, http.StatusBadRequest, err)
		return
	}
	var concept model.Concept
	concept = model.Concept{
		Key:    param.Data.Key,
		Detail: param.Data.Detail,
	}
	tx := database.Engine.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if dbError := tx.Save(&concept).Error; dbError != nil {
		tx.Rollback()
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}
	tx.Commit()
	make_result.Result(w, r, http.StatusOK, concept.Serializer())
}

func getOneConceptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conceptId := vars["concept_id"]

	var concept model.Concept
	if dbError := database.Engine.Where("id = ?", conceptId).First(&concept).Error; dbError != nil {
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}

	make_result.Result(w, r, http.StatusOK, concept.Serializer())
}

func deleteOneConceptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	conceptId := vars["concept_id"]

	tx := database.Engine.Begin()
	defer tx.Close()

	var concept model.Concept
	if dbError := tx.Where("id = ?", conceptId).First(&concept).Error; dbError != nil {
		make_result.Result(w, r, http.StatusBadRequest, dbError)
		return
	}
	if concept.ID != 0 {
		if dbError := tx.Delete(&concept).Error; dbError != nil {
			tx.Rollback()
			make_result.Result(w, r, http.StatusBadRequest, dbError)
			return
		}
	}
	tx.Commit()
	make_result.Result(w, r, http.StatusOK, concept.Serializer())

}
