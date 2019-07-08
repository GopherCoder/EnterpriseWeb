package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(32)" json:"name"`
	WebSite       string    `gorm:"type:varchar(64)" json:"web_site"`
	Valuation     uint      `gorm:"type:bigint" json:"valuation"`
	ValuationDate time.Time `json:"valuation_date"`
	CountryID     uint      `gorm:"index"`
	CategoryID    uint      `gorm:"index"`
}

//func (c Company) TableName() string {
//	return "company"
//}

type CompanySerializer struct {
	Id            uint      `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	WebSite       string    `json:"web_site"`
	Valuation     uint      `json:"valuation"`
	ValuationDate time.Time `json:"valuation_date"`
	CountryName   string    `json:"country_name"`
	CategoryName  string    `json:"category_name"`
}

func (c Company) Serializer(db *gorm.DB) CompanySerializer {

	countryName := func(id uint) string {
		var country Country
		db.Where("id = ?", id).First(&country)
		return country.Name
	}

	categoryName := func(id uint) string {
		var category Category
		db.Where("id = ?", id).First(&category)
		return category.Name
	}

	return CompanySerializer{
		Id:            c.ID,
		CreatedAt:     c.CreatedAt,
		UpdatedAt:     c.UpdatedAt,
		Name:          c.Name,
		WebSite:       c.WebSite,
		Valuation:     c.Valuation,
		ValuationDate: c.ValuationDate,
		CategoryName:  categoryName(c.CategoryID),
		CountryName:   countryName(c.CountryID),
	}
}
