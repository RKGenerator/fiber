package dto

import "time"

type ApartmentResponse struct {
	Id             int64     `json:"id"`
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

type ApartmentDetailsResponse struct {
	ApartmentResponse
	PropertyImagesResponse
}

type ApartmentDetailsRequest struct {
	DistrictId   *int64 `query:"district"`
	AreaSQFTFrom *int32 `query:"area_sqft_from"`
	AreaSQFTTo   *int32 `query:"area_sqft_to"`
	FloorTo      *int16 `query:"floor_to"`
	FloorFrom    *int16 `query:"floor_from"`
}
