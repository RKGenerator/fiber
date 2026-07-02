package dto

import "time"

type ApartmentDTO struct {
	SourcePriceAED int64     `json:"sourcePrice"`
	Address        string    `json:"address"`
	PropertyType   string    `json:"propertyType"`
	Purpose        string    `json:"purpose"`
	NBedrooms      int16     `json:"nBedrooms"`
	AreaSQFT       int32     `json:"areaSQFT"`
	Description    string    `json:"description"`
	BuildingID     int64     `json:"buildingId"`
	CreatedAt      time.Time `json:"createdAt"`
	PriceSQFT      float32   `json:"priceSQFT"`
	Floor          int16     `json:"floor"`
	Booking        bool      `json:"booking"`
}
