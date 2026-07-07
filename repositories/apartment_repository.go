package repositories

import (
	"log"
	"test-fiber/dto"
	"test-fiber/model"

	"gorm.io/gorm"
)

type ApartmentsRepository struct {
	conn *gorm.DB
}

func NewApartmentsRepository(conn *gorm.DB) *ApartmentsRepository {
	return &ApartmentsRepository{
		conn: conn,
	}
}

func (r *ApartmentsRepository) GetApartments() ([]*model.Apartment, error) {
	var apartments []*model.Apartment

	err := r.conn.Find(&apartments).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return apartments, nil
}

func (r *ApartmentsRepository) GetReq(req dto.ApartmentDetailsRequest) ([]*model.Apartment, error) {

	query := r.conn.Model(&model.Apartment{}).Preload("PropertyImages")

	if req.DistrictId != nil {
		query = query.Where("district_is = ?", req.DistrictId)
	}

	if req.AreaSQFTFrom != nil {
		query = query.Where("area_sqft >= ?", req.AreaSQFTFrom)
	}

	if req.AreaSQFTTo != nil {
		query = query.Where("area_sqft <= ?", req.AreaSQFTTo)
	}

	if req.FloorFrom != nil {
		query = query.Where("floor >= ?", req.FloorFrom)
	}

	if req.FloorTo != nil {
		query = query.Where("floor <= ?", req.FloorTo)
	}

	var apartments []*model.Apartment
	err := query.Find(&apartments).Error
	if err != nil {
		return nil, err
	}

	return apartments, nil
}
