package services

import (
	"log"
	"test-fiber/dto"
	"test-fiber/repositories"
)

type ApartmentsServices struct {
	apartmentsRepository repositories.ApartmentsInterface
}

func NewApartmentsService(apartmentRep repositories.ApartmentsInterface) *ApartmentsServices {
	return &ApartmentsServices{
		apartmentsRepository: apartmentRep,
	}
}

func (s *ApartmentsServices) GetApartments() ([]*dto.ApartmentDTO, error) {
	query, err := s.apartmentsRepository.GetApartments()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	dto, err := dto.ApartmentsDTO(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto, nil
}

func (s *ApartmentsServices) GetApartmentsExpirence() ([]*dto.ApartmentDetailsDTO, error) {
	querry, err := s.apartmentsRepository.GetApartmentsExpirence()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	dto, err := dto.ApartmentsDetailsDTO(querry)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto, nil
}

func (s *ApartmentsServices) GetBedrooms() (int, error) {
	return s.apartmentsRepository.GetNBedrooms()
}
