package services

import (
	"log"
	"test-fiber/dto"
	"test-fiber/repositories"
)

type ApartmentServices struct {
	apartmentsRepository repositories.ApartmentsInterface
}

func NewApartmentsService(apartmentRep repositories.ApartmentsInterface) *ApartmentServices {
	return &ApartmentServices{
		apartmentsRepository: apartmentRep,
	}
}

func (s *ApartmentServices) GetApartments() ([]*dto.ApartmentResponse, error) {
	query, err := s.apartmentsRepository.GetApartments()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	dto := dto.GetApartmentsResponse(query)

	return dto, nil
}

func (s *ApartmentServices) GetApartmentsDetail(req dto.ApartmentDetailsRequest) ([]*dto.ApartmentDetailsResponse, error) {

	querry, err := s.apartmentsRepository.GetApartmentsDetail(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dto := dto.GetApartmentsDetailsResponse(querry)

	return dto, nil

}
