package repositories

import (
	"fmt"
	"test-fiber/model"

	"gorm.io/gorm"
)

type BuildingRepository struct {
	conn *gorm.DB
}

func NewBuildingRepositories(conn *gorm.DB) *BuildingRepository {
	return &BuildingRepository{
		conn: conn,
	}
}

func (r *BuildingRepository) GetBuildingsDetail() ([]*model.Building, error) {
	query := r.conn.Model(&model.Building{}).Preload("District")

	var buildings []*model.Building
	err := query.Find(&buildings).Error
	if err != nil {
		return nil, err
	}
	fmt.Println(buildings[2])
	return buildings, nil
}
