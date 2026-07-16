package model

type User struct {
	Id          int    `gorm:"primaryKey"`
	Password    string `gorm:"type:varchar(128)"`
	IsSuperuser bool   `gorm:"default:false"`
	Email       string `gorm:"type:varchar(254)"`
	IsStaff     bool   `gorm:"default:false"`
	IsActive    bool   `gorm:"default:false"`
}

func (*User) TableName() string {
	return "auth_user"
}
