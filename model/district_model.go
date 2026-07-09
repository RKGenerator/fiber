package model

type District struct {
	Id   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}

func (District) TableName() string {
	return "parser_district"
}
