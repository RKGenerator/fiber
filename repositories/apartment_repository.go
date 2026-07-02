package repositories

import (
	"log"
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

	err := r.conn.Where("n_bedrooms > 1").Find(&apartments).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return apartments, nil
}

func (r *ApartmentsRepository) GetNBedrooms() (int, error) {
	return 0, nil
}

/*
func (r *ApartmentsRepository) GetNBedrooms() (int, error) {
	var n_bedrooms int
	row := r.conn.QueryRow(context.Background(), "SELECT id FROM parser_apartment WHERE n_bedrooms = $1", 2)

	err := row.Scan(&n_bedrooms)
	if err != nil {
		log.Fatal(err)
	}

	return n_bedrooms, nil
}
*/
