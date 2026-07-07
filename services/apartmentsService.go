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

func (s *ApartmentsServices) GetApartments() ([]*dto.ApartmentResponse, error) {
	query, err := s.apartmentsRepository.GetApartments()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	dto, err := dto.ApartmentsResponse(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto, nil
}

func (s *ApartmentsServices) GetReq(req dto.ApartmentDetailsRequest) ([]*dto.ApartmentDetailsResponse, error) {

	querry, err := s.apartmentsRepository.GetReq(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dto, err := dto.ApartmentsDetailsResponse(querry)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return dto, nil

}
