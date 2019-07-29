package admin

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/assistance"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/log"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"fmt"
	"time"
)

func DefaultAdmin() (*model.AdminSerializer, error) {
	return &model.AdminSerializer{
		Id:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "GraphQL",
		Phone:     "18717711819",
		Token:     "18717711819token",
	}, nil
}

func GetAdmin(id int64) (*model.AdminSerializer, error) {
	var admin model.Admin
	var result model.AdminSerializer
	if has, err := database.Engine.ID(id).Get(&admin); !has || err != nil {
		if !has {
			log_for_lottery.Println("record not found")
			return &result, fmt.Errorf("record not found")
		}
		if err != nil {
			log_for_lottery.Println(err)
			return &result, err
		}
	}
	result = admin.Serializer()
	return &result, nil
}

func Login(params LoginParam) (*model.AdminSerializer, error) {
	var admin model.Admin
	if has, err := database.Engine.Where("phone = ?", params.Phone).Get(&admin); !has || err != nil {
		return nil, err
	}
	if ok := assistance.CompareHashAndPassword([]byte(admin.Password), []byte(params.Password)); !ok {
		return nil, fmt.Errorf("password fail")
	}
	var result model.AdminSerializer
	result = admin.Serializer()
	return &result, nil
}

func CreateAdmin(params LoginParam) (*model.AdminSerializer, error) {
	var admin model.Admin
	var result model.AdminSerializer
	secret, _ := assistance.GenerateFromPassword(params.Password)
	token := assistance.GenerateToken(10)
	admin = model.Admin{
		Password: string(secret),
		Token:    token,
		Phone:    params.Phone,
		Name:     params.Phone,
	}
	tx := database.Engine.NewSession()
	tx.Begin()
	_, err := tx.InsertOne(&admin)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	result = admin.Serializer()
	return &result, nil
}

func UpdateAdmin(params UpdateAdminParams) (*model.AdminSerializer, error) {

	if err := params.Valid(); err != nil {
		return nil, err
	}

	tx := database.Engine.NewSession()
	tx.Begin()
	var admin model.Admin
	if has, err := tx.ID(params.AdminId).Get(&admin); !has || err != nil {
		return nil, fmt.Errorf("record not found")
	}
	admin.Name = params.Name
	if _, dbError := tx.ID(admin.Id).Cols("name").Update(&admin); dbError != nil {
		tx.Rollback()
		return nil, dbError
	}
	tx.Commit()
	var result model.AdminSerializer
	result = admin.Serializer()
	return &result, nil
}
