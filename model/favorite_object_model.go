package model

import "time"

type FavoriteObject struct {
	Id          int64     `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"type:timestamptz"`
	ApartmentId int64
	UserId      int
}

func (FavoriteObject) TableName() string {
	return "parser_favorite_object"
}
