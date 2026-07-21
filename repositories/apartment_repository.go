package repositories

import (
	"log"
	"test-fiber/dto"
	"test-fiber/model"
	"time"

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

func (r *ApartmentsRepository) GetApartmentsDetail(req dto.ApartmentDetailsRequest) ([]*model.Apartment, error) {

	query := r.conn.Model(&model.Apartment{}).Preload("PropertyImages").Preload("Tags").Preload("Building.District")
	if req.DistrictId != nil {
		query = query.Joins("INNER JOIN parser_building AS buildingN on buildingN.id = parser_apartment.building_id").
			Where("buildingN.district = ?", req.DistrictId)
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

	query = query.Scopes(Paginator(req.Page, req.PageSize))

	var apartments []*model.Apartment
	err := query.Find(&apartments).Error

	if err != nil {
		return nil, err
	}

	return apartments, nil
}

func (r *ApartmentsRepository) GetFavoriteToUser(userId int, apartmentIds []int64) ([]int64, error) {

	var favoriteIds []int64

	if len(apartmentIds) > 0 {
		err := r.conn.Table("parser_favorite_object").
			Where("user_id = ? AND apartment_id IN (?)", userId, apartmentIds).
			Pluck("apartment_id", &favoriteIds).Error
		if err != nil {
			return nil, err
		}

		return favoriteIds, nil
	}

	return nil, nil
}

func (r *ApartmentsRepository) AddFavoriteApartmentToUser(apartId int64, userId int) error {
	data := &model.FavoriteObject{
		CreatedAt:   time.Now(),
		UserId:      userId,
		ApartmentId: apartId,
	}

	err := r.conn.Create(data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ApartmentsRepository) DeleteFavoritesToUser(apartIds []int64, userId int) (*int64, error) {
	res := r.conn.Model(&model.FavoriteObject{}).Where("user_id = ?", userId).
		Where("apartment_id IN (?)", apartIds).Delete(&model.FavoriteObject{})

	if res.Error != nil {
		return nil, res.Error
	}

	return &res.RowsAffected, nil
}
