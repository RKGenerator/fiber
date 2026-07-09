package model

/*
	USE MANY 2 MANY RELATION

	type ApartmentTag struct {
		Id          int64 `gorm:"primaryKey"`
		ApartmentId int64 `gorm:"index"`
		TagsId      []Tag `gorm:"foreignkey:Id"`
	}

	func (ApartmentTag) TableName() string {
		return "parser_apartment_tags"
	}
*/
type Tag struct {
	Id   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (Tag) TableName() string {
	return "parser_tags"
}
