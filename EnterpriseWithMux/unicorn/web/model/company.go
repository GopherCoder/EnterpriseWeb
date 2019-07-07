package model

import "time"

type Company struct {
	base
	Name          string    `gorm:"type:varchar(12)" json:"name"`
	WebSite       string    `gorm:"type:varchar(64)" json:"web_site"`
	Valuation     uint      `gorm:"type:bigint" json:"valuation"`
	ValuationDate time.Time `gorm:"type:timestamp with time zone" json:"valuation_date"`
	CountryID     uint
	CategoryID    uint
}

func (c Company) TableName() string {
	return "company"
}
