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

	dto := dto.GetApartmentsDetailsResponseToAnonym(querry)

	return dto, nil

}

func (s *ApartmentServices) GetApartmentsDetailToAnonym(req dto.ApartmentDetailsRequest) ([]*dto.ApartmentDetailsResponse, error) {

	apartments, err := s.apartmentsRepository.GetApartmentsDetail(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dto := dto.GetApartmentsDetailsResponseToAnonym(apartments)

	return dto, nil

}

func (s *ApartmentServices) GetApartmentsDetailToUser(userId int, req dto.ApartmentDetailsRequest) ([]*dto.ApartmentDetailsResponse, error) {

	apartments, err := s.apartmentsRepository.GetApartmentsDetail(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	apartmentsIds := make([]int64, len(apartments))
	for i, a := range apartments {
		apartmentsIds[i] = a.Id
	}

	favoriteIds, err := s.apartmentsRepository.GetFavoriteToUser(userId, apartmentsIds)
	if err != nil {
		return nil, err
	}

	dto := dto.GetApartmentsDetailsResponseToUser(apartments, favoriteIds)

	return dto, nil

}

func (s *ApartmentServices) AddFavoriteApartment(apartId int64, userId int) error {
	err := s.apartmentsRepository.AddFavoriteApartmentToUser(apartId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *ApartmentServices) DeleteFavoriteApartments(apartIds []int64, userId int) (*dto.FavoritesResponse, error) {
	rowCount, err := s.apartmentsRepository.DeleteFavoritesToUser(apartIds, userId)
	if err != nil {
		return nil, err
	}

	return &dto.FavoritesResponse{
		DeletedCount: *rowCount,
	}, nil
}
