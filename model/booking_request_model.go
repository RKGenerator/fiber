package model

import "time"

type Booking struct {
	Id          int64  `gorm:"primaryKey"`
	FullName    string `gorm:"type:varchar(100)"`
	PhoneNumber string `gorm:"type:varchar(20)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ApartmentId int64 `gorm:"index"`
	Reviewed    bool  `gorm:"default:false"`

	Apartment Apartment
}

func (Booking) TableName() string {
	return "parser_booking_requests"
}
