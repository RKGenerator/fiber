package repositories

import (
	"test-fiber/dto"
	"test-fiber/model"
)

type ApartmentsInterface interface {
	GetApartments() ([]*model.Apartment, error)
	GetReq(dto.ApartmentDetailsRequest) ([]*model.Apartment, error)
}
