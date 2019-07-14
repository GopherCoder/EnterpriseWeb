package admin

import (
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/database"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/pkg/error"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/make_response"
	"EnterpriseWeb/EnterpriseWithHttpRouter/tencent_vote/web/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Admin struct {
}

var Default = Admin{}

func (a Admin) Register(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		make_response.Result(writer, request, http.StatusBadRequest, error_tencent_votes.ErrorMethod)
		return
	}
	var param RegisterParams
	if err := json.NewDecoder(request.Body).Decode(&param); err != nil {
		make_response.Result(writer, request, http.StatusBadRequest, fmt.Errorf("bind error fail").Error())
		return
	}
	if err := param.Valid(); err != nil {
		make_response.Result(writer, request, http.StatusBadRequest, fmt.Errorf("invaild param fail").Error())
		return
	}
	log.Println("Param: ", param)
	var admin model.Admin
	if dbError := database.Engine.Where("phone = ?", param.Data.Phone).First(&admin).Error; dbError == nil {
		make_response.Result(writer, request, http.StatusBadRequest, error_tencent_votes.ErrorExistsRecord)
		return
	}

	password, _ := generateFromPassword(param.Data.Password)
	token := generateToken(10)
	admin = model.Admin{
		Phone:    param.Data.Phone,
		Password: string(password),
		Token:    token,
	}
	if dbError := database.Engine.Save(&admin).Error; dbError != nil {
		make_response.Result(writer, request, http.StatusBadRequest, fmt.Errorf("save record fail"))
		return
	}
	make_response.Result(writer, request, http.StatusOK, admin.Serializer())
}

func (a Admin) Login(writer http.ResponseWriter, request *http.Request) {
	var param RegisterParams
	if err := json.NewDecoder(request.Body).Decode(&param); err != nil {
		make_response.Result(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("Param: ", param)
	if err := param.Valid(); err != nil {
		make_response.Result(writer, request, http.StatusBadRequest, err.Error())
		return
	}
	var admin model.Admin
	if dbError := database.Engine.Where("phone = ?", param.Data.Phone).First(&admin).Error; dbError != nil {
		make_response.Result(writer, request, http.StatusBadRequest, dbError.Error())
		return
	}
	if ok := compareHashAndPassword([]byte(admin.Password), []byte(param.Data.Password)); !ok {
		make_response.Result(writer, request, http.StatusBadRequest, fmt.Errorf("password fail").Error())
		return
	}
	make_response.Result(writer, request, http.StatusOK, admin.Serializer())

}

func (a Admin) Logout(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/ping", http.StatusOK)
}
