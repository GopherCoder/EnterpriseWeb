package model

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"fmt"
)

func GetAddresses(adminID int64, orderBy string) ([]*AddressSerialize, error) {
	var addresses []Address
	var results []*AddressSerialize
	if orderBy == "" {
		orderBy = "created_at"
	}
	if dbError := database.Engine.ID(adminID).Desc(orderBy).Find(&addresses); dbError != nil {
		return results, dbError
	}
	for _, i := range addresses {
		s := i.Serializer()
		results = append(results, &s)
	}
	return results, nil
}

func CreateAddress(adminId int64, detail string) (*AddressSerialize, error) {
	var admin Admin
	if has, dbError := database.Engine.ID(adminId).Get(&admin); !has || dbError != nil {
		return nil, fmt.Errorf("record not found")
	}
	var address Address
	address = Address{
		AdminId: adminId,
		Detail:  detail,
	}
	tx := database.Engine.NewSession()
	tx.Begin()
	if _, dbError := tx.InsertOne(&address); dbError != nil {
		tx.Rollback()
		return nil, dbError
	}
	tx.Commit()
	var result AddressSerialize
	result = address.Serializer()
	return &result, nil

}

func UpdateAddress(id int64, detail string) (*AddressSerialize, error) {
	var address Address
	tx := database.Engine.NewSession()
	tx.Begin()
	if ok, dbError := tx.ID(id).Get(&address); !ok || dbError != nil {
		return nil, fmt.Errorf("record not found")
	}
	address.Detail = detail
	if _, dbError := tx.ID(id).Update(&address); dbError != nil {
		tx.Rollback()
		return nil, dbError
	}
	tx.Commit()
	var result AddressSerialize
	result = address.Serializer()
	return &result, nil

}
