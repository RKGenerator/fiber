package dto

type BuildingResponse struct {
	Id           int64    `json:"id"`
	Name         string   `json:"name"`
	ImageUrl     string   `json:"image"`
	Latitude     float32  `json:"lat"`
	Longitude    float32  `json:"lon"`
	BuildingType string   `json:"building_type"`
	District     District `json:"district"`
}

type BuildingDetailsResponse struct {
	BuildingResponse
}

type BuildingRequest struct {
	District  *int64   `query:"district"`
	NBedrooms *[]int64 `query:"n_bedrooms"`

	PaginatorRequest
}
