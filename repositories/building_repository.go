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

func (r *BuildingRepository) GetBuildings() ([]*model.Building, error) {
	query := r.conn.Model(&model.Building{}).Preload("District")

	var buildings []*model.Building
	err := query.Find(&buildings).Error
	if err != nil {
		return nil, err
	}

	return buildings, nil
}

func (r *BuildingRepository) GetBuildingsWithStats() ([]*model.BuildingWithStats, error) {
	query := r.conn.Model(&model.BuildingWithStats{}).Preload("District").
		Select(`parser_building.*, COUNT(parser_apartment.id) AS total_count,
	percentile_cont(0.5) WITHIN GROUP (ORDER BY parser_apartment.source_price_aed) AS median_price,
	percentile_cont(0.5) WITHIN GROUP (ORDER BY parser_apartment.price_sqft) AS median_price_sqft,
	percentile_cont(0.5) WITHIN GROUP (ORDER BY parser_apartment.area_sqft) AS median_area_sqft`).
		Joins("LEFT JOIN parser_apartment ON parser_apartment.building_id = parser_building.id").
		Where("parser_apartment.disabled = ?", false).Group("parser_building.id")

	var buildings []*model.BuildingWithStats
	err := query.Scan(&buildings).Error
	if err != nil {
		return nil, err
	}

	return buildings, nil
}

func (r *BuildingRepository) GetBuildingStatsByBedrooms() ([]model.RoomBuldingStats, error) {
	var res []model.RoomBuldingStats
	err := r.conn.Table("parser_building").Select(`parser_building.id as building_id,
	parser_apartment.n_bedrooms as rooms,
	COUNT(parser_apartment.id) as count,
	percentile_cont(0.5) WITHIN GROUP (ORDER BY parser_apartment.source_price_aed) as median_price,
	percentile_cont(0.5) WITHIN GROUP (ORDER BY parser_apartment.price_sqft) as median_price_sqft,
	percentile_cont(0.5) WITHIN GROUP (ORDER BY parser_apartment.area_sqft) as median_area_sqft`).
		Joins("LEFT JOIN parser_apartment ON parser_apartment.building_id = parser_building.id").
		Where("parser_apartment.disabled = ?", false).
		Group("parser_building.id, parser_apartment.n_bedrooms").Scan(&res).Error

	if err != nil {
		return nil, err
	}
	fmt.Println(*res[0].MedianAreaSQFT)
	return res, nil
}
