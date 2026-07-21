package dto

import "time"

type ApartmentResponse struct {
	Id           int64     `json:"id"`
	SourcePrice  int64     `json:"source_price_aed"`
	Address      string    `json:"address"`
	PropertyType string    `json:"property_type"`
	Purpose      string    `json:"purpose"`
	NBedrooms    int16     `json:"n_bedrooms"`
	AreaSQFT     int32     `json:"area_sqft"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	PriceSQFT    float32   `json:"price_sqft"`
	Floor        int16     `json:"floor"`
	Booking      bool      `json:"booking"`
	ExternalId   int64     `json:"external_id"`
	Disabled     bool      `json:"disabled"`
	IsFavorite   bool      `json:"is_favorite"`
}

type ApartmentDetailsResponse struct {
	ApartmentResponse
	PropertyImagesResponse
	Bulding BuildingResponse `json:"building"`
	Tags    []Tag            `json:"tags"`
}

type ApartmentDetailsRequest struct {
	DistrictId   *int64 `query:"district"`
	AreaSQFTFrom *int32 `query:"area_sqft_from"`
	AreaSQFTTo   *int32 `query:"area_sqft_to"`
	FloorTo      *int16 `query:"floor_to"`
	FloorFrom    *int16 `query:"floor_from"`

	PaginatorRequest
}
