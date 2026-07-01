package services

import (
	"test-fiber/repositories"
)

type ApartmentsServices struct {
	apartmentsRepository repositories.ApartmentsInterface
}

func NewUserService(apartmentRep repositories.ApartmentsInterface) *ApartmentsServices {
	return &ApartmentsServices{
		apartmentsRepository: apartmentRep,
	}
}

func (s *ApartmentsServices) GetUser() (*int, error) {
	return s.apartmentsRepository.GetApartments()
}

func (s *ApartmentsServices) GetBedrooms() (int, error) {
	return s.apartmentsRepository.GetNBedrooms()
}
