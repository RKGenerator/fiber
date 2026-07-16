package repositories

import (
	"test-fiber/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User

	err := r.conn.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
