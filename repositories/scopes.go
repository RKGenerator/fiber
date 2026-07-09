package repositories

import "gorm.io/gorm"

func Paginator(page, page_size int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * page_size
		return db.Offset(offset).Limit(page_size)
	}
}
