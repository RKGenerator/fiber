package dto

type BuildingResponse struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	ImageUrl     string  `json:"image"`
	Latitude     float32 `json:"lat"`
	Longitude    float32 `json:"lon"`
	BuildingType string  `json:"building_type"`
}
