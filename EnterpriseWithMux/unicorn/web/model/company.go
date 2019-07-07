package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Company struct {
	gorm.Model
	Name          string    `gorm:"type:varchar(12)" json:"name"`
	WebSite       string    `gorm:"type:varchar(64)" json:"web_site"`
	Valuation     uint      `gorm:"type:bigint" json:"valuation"`
	ValuationDate time.Time `json:"valuation_date"`
	CountryID     uint      `gorm:"index"`
	CategoryID    uint      `gorm:"index"`
}

//func (c Company) TableName() string {
//	return "company"
//}
