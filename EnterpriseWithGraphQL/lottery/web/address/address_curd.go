package address

import (
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/pkg/database"
	"EnterpriseWeb/EnterpriseWithGraphQL/lottery/web/model"
	"fmt"
)

func GetAddresses(params GetAddressParams) ([]*model.AddressSerialize, error) {
	var addresses []model.Address
	var results []*model.AddressSerialize
	if params.OrderBy == "" {
		params.OrderBy = "created_at"
	}
	if dbError := database.Engine.ID(params.AdminId).Desc(params.OrderBy).Limit(params.Limit, params.Offset).Find(&addresses); dbError != nil {
		return results, dbError
	}
	for _, i := range addresses {
		s := i.Serializer()
		results = append(results, &s)
	}
	return results, nil
}

func CreateAddress(params CreateAddressParams) (*model.AddressSerialize, error) {
	var admin model.Admin
	if has, dbError := database.Engine.ID(params.AdminId).Get(&admin); !has || dbError != nil {
		return nil, fmt.Errorf("record not found")
	}
	var address model.Address
	address = model.Address{
		AdminId: params.AdminId,
		Detail:  params.Detail,
	}
	tx := database.Engine.NewSession()
	tx.Begin()
	if _, dbError := tx.InsertOne(&address); dbError != nil {
		tx.Rollback()
		return nil, dbError
	}
	tx.Commit()
	var result model.AddressSerialize
	result = address.Serializer()
	return &result, nil

}

func UpdateAddress(params UpdateAddressParams) (*model.AddressSerialize, error) {
	var address model.Address
	tx := database.Engine.NewSession()
	tx.Begin()
	if ok, dbError := tx.ID(params.Id).Get(&address); !ok || dbError != nil {
		return nil, fmt.Errorf("record not found")
	}
	address.Detail = params.Detail
	if _, dbError := tx.ID(address.Id).Cols("detail").Update(&address); dbError != nil {
		tx.Rollback()
		return nil, dbError
	}
	tx.Commit()
	var result model.AddressSerialize
	result = address.Serializer()
	return &result, nil

}
