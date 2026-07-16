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
	GetBuildings() ([]*model.Building, error)
	GetBuildingStatsByBedrooms() ([]model.RoomBuldingStats, error)
	GetBuildingsWithStats() ([]*model.BuildingWithStats, error)
}

type UserInterface interface {
	GetUserByEmail(string) (*model.User, error)
}
