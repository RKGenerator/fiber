package model

import (
	"time"
)

type Apartment struct {
	Id             int64 `gorm:"primaryKey"`
	UpdatedAt      time.Time
	SourcePriceAED int64
	Address        string `gorm:"type:varchar(300)"`
	PropertyType   string `gorm:"type:varchar(300)"`
	Purpose        string `gorm:"type:varchar(300)"`
	NBedrooms      int16
	AreaSQFT       int32
	Description    string
	BuildingID     int64
	CreatedAt      time.Time
	PriceSQFT      float32
	Disabled       bool `gorm:"default:false"`
	TouchedAt      time.Time
	DistrictId     int64
	ExternalId     int64
	Floor          int16
	Booking        bool `gorm:"default:false"`

	Bookings       []Booking       `gorm:"foreignKey:ApartmentId"`
	PropertyImages []PropertyImage `gorm:"foreignKey:PropertyId"`
}

func (Apartment) TableName() string {
	return "parser_apartment"
}
