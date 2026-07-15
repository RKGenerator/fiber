package model

type Building struct {
	Id             int64 `gorm:"primaryKey"`
	Latitude       float32
	Longitude      float32
	Name           string `gorm:"type:varchar(255)"`
	TypeOfBuilding string `gorm:"type:varchar(255)"`
	ImageUrl       string `gorm:"type:varchar(1000)"`
	DistrictId     int64  `gorm:"column:district;index"`

	District District `gorm:"foreignKey:DistrictId"`
}

func (Building) TableName() string {
	return "parser_building"
}

type BuildingWithStats struct {
	Building
	TotalCount      int64    `gorm:"column:total_count"`
	MedianPrice     *float64 `gorm:"column:median_price"`
	MedianPriceSQFT *float64 `gorm:"column:median_price_sqft"`
	MedianAreaSQFT  *float64 `gorm:"column:median_area_sqft"`
}

type RoomBuldingStats struct {
	BuildingId      int64    `gorm:"column:building_id"`
	Rooms           int      `gorm:"column:rooms"`
	Count           int64    `gorm:"column:count"`
	MedianPrice     *float64 `gorm:"column:median_price"`
	MedianPriceSQFT *float64 `gorm:"column:median_price_sqft"`
	MedianAreaSQFT  *float64 `gorm:"column:median_area_sqft"`
}
