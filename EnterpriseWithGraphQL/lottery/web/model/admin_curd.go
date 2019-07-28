package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/assistance"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	log_for_lottery "EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/log"
	"fmt"
	"time"
)

func DefaultAdmin() (*AdminSerializer, error) {
	return &AdminSerializer{
		Id:        1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "GraphQL",
		Phone:     "18717711819",
		Token:     "18717711819token",
	}, nil
}

func GetAdmin(id int64) (*AdminSerializer, error) {
	var admin Admin
	var result AdminSerializer
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

func Login(params LoginParam) (*AdminSerializer, error) {
	var admin Admin
	if has, err := database.Engine.Where("phone = ?", params.Phone).Get(&admin); !has || err != nil {
		return nil, err
	}
	if ok := assistance.CompareHashAndPassword([]byte(admin.Password), []byte(params.Password)); !ok {
		return nil, fmt.Errorf("password fail")
	}
	var result AdminSerializer
	result = admin.Serializer()
	return &result, nil
}

func CreateAdmin(params LoginParam) (*AdminSerializer, error) {
	var admin Admin
	var result AdminSerializer
	secret, _ := assistance.GenerateFromPassword(params.Password)
	token := assistance.GenerateToken(10)
	admin = Admin{
		Password: string(secret),
		Token:    token,
		Phone:    params.Phone,
		Name:     params.Password,
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
