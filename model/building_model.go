package model

type Building struct {
	Id             int64 `gorm:"primaryKey"`
	Latitude       float32
	Longitude      float32
	Name           string `gorm:"type:varchar(255)"`
	TypeOfBuilding string `gorm:"type:varchar(255)"`
	ImageUrl       string `gorm:"type:varchar(1000)"`

	Apartments []Apartment `gorm:"foreignKey:BuildingId"`
	District   District    `gorm:"foreignKey:Id"`
}

func (Building) TableName() string {
	return "parser_building"
}
