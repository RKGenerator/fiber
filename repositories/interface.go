package repositories

import "test-fiber/model"

type ApartmentsInterface interface {
	GetApartments() ([]*model.Apartment, error)
	GetNBedrooms() (int, error)
}
