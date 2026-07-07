package model

type Building struct {
	Id             int64   `gorm:"id"`
	Latitude       float32 `gorm:"latitude"`
	Longitude      float32 `gorm:"longitude"`
	Name           string  `gorm:"type:varchar(255)"`
	TypeOfBuidings string  `gorm:"type:varchar(255)"`
	ImageUrl       string  `gorm:"type:varchar(1000)"`

	Apartments []Apartment `gorm:"foreignKey:BuildingId"`
}

func (Building) TableName() string {
	return "parser_building"
}
