package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/assistance"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	log_for_lottery "EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/log"
	"fmt"
	"time"
)

func DefaultAdmin() (*Admin, error) {
	return &Admin{
		Base: Base{
			Id:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		Name:     "Graph",
		Phone:    "18717711819",
		Password: "1871771819password",
		Token:    "18717711819token",
	}, nil
}

func GetAdmin(id int64) (*Admin, error) {
	var result Admin
	if has, err := database.Engine.ID(id).Get(&result); !has || err != nil {
		if !has {
			log_for_lottery.Println("record not found")
			return &result, fmt.Errorf("record not found")
		}
		if err != nil {
			log_for_lottery.Println(err)
			return &result, err
		}
	}
	return &result, nil
}

func CreateAdmin(phone string, password string) (*Admin, error) {
	var result Admin
	secret, _ := assistance.GenerateFromPassword(password)
	token := assistance.GenerateToken(10)
	result = Admin{
		Password: string(secret),
		Token:    token,
		Phone:    phone,
	}
	tx := database.Engine.NewSession()
	tx.Begin()
	_, err := tx.InsertOne(&result)
	if err != nil {
		tx.Rollback()
		return &result, err
	}
	tx.Commit()
	return &result, nil
}
