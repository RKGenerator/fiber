package model

type PropertyImage struct {
	Id         int64  `qorm:"primaryKey"`
	Url        string `qorm:"type:varchar(1000)"`
	PropertyId int64  `qorm:"index"`
	Processed  bool   `qorm:"default:false"`
}

func (PropertyImage) TableName() string {
	return "parser_propertiesimages"
}
