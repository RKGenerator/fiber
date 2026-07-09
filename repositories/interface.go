package repositories

import (
	"test-fiber/dto"
	"test-fiber/model"
)

type ApartmentsInterface interface {
	GetApartments() ([]*model.Apartment, error)
	GetApartmentsDetail(dto.ApartmentDetailsRequest) ([]*model.Apartment, error)
}

type BuildingInterface interface {
	GetBuildingsDetail() ([]*model.Building, error)
}
