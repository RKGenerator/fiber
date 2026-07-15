package dto

type BuildingResponse struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	ImageUrl     string  `json:"image"`
	Latitude     float32 `json:"lat"`
	Longitude    float32 `json:"lon"`
	BuildingType string  `json:"building_type"`
	District     int64   `json:"district"`
}

type BuildingDetailsResponse struct {
	BuildingResponse
	District        District    `json:"district"`
	TotalCount      int64       `json:"listing_count"`
	MedianPrice     *float64    `json:"median_price"`
	MedianPriceSQFT *float64    `json:"median_price_sqft"`
	MedianAreaSQFT  *float64    `json:"median_area_sqft"`
	RoomStats       []RoomStats `json:"room_state"`
}

type BuildingRequest struct {
	District  *int64   `query:"district"`
	NBedrooms *[]int64 `query:"n_bedrooms"`

	PaginatorRequest
}

type RoomStats struct {
	Rooms           int      `json:"rooms"`
	Count           int64    `json:"count"`
	MedianPrice     *float64 `json:"median_price"`
	MedianPriceSQFT *float64 `json:"median_price_sqft"`
	MedianAreaSQFT  *float64 `json:"median_area_sqft"`
}
