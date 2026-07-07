package dto

import (
	"log"
	"test-fiber/model"
)

func ApartmentsResponse(apartsModel []*model.Apartment) ([]*ApartmentResponse, error) {
	if apartsModel == nil {
		return nil, nil
	}

	var apartments []*ApartmentResponse

	for _, apartmentModel := range apartsModel {
		log.Println(apartmentModel)
		apartment := ApartmentResponse{
			Id:             apartmentModel.Id,
			SourcePriceAED: apartmentModel.SourcePriceAED,
			Address:        apartmentModel.Address,
			PropertyType:   apartmentModel.PropertyType,
			Purpose:        apartmentModel.Purpose,
			NBedrooms:      apartmentModel.NBedrooms,
			AreaSQFT:       apartmentModel.AreaSQFT,
			Description:    apartmentModel.Description,
			BuildingID:     apartmentModel.BuildingId,
			CreatedAt:      apartmentModel.CreatedAt,
			PriceSQFT:      apartmentModel.PriceSQFT,
			Floor:          apartmentModel.Floor,
			Booking:        apartmentModel.Booking}

		apartments = append(apartments, &apartment)
	}

	return apartments, nil

}

func ApartmentsDetailsResponse(apartsModel []*model.Apartment) ([]*ApartmentDetailsResponse, error) {
	if apartsModel == nil {
		return nil, nil
	}

	apartments := make([]*ApartmentDetailsResponse, 0, len(apartsModel))

	for _, apartmentModel := range apartsModel {
		apartment := ApartmentDetailsResponse{
			ApartmentResponse: ApartmentResponse{
				Id:             apartmentModel.Id,
				SourcePriceAED: apartmentModel.SourcePriceAED,
				Address:        apartmentModel.Address,
				PropertyType:   apartmentModel.PropertyType,
				Purpose:        apartmentModel.Purpose,
				NBedrooms:      apartmentModel.NBedrooms,
				AreaSQFT:       apartmentModel.AreaSQFT,
				Description:    apartmentModel.Description,
				BuildingID:     apartmentModel.BuildingId,
				CreatedAt:      apartmentModel.CreatedAt,
				PriceSQFT:      apartmentModel.PriceSQFT,
				Floor:          apartmentModel.Floor,
				Booking:        apartmentModel.Booking,
			},
			PropertyImagesResponse: *ImagesDTO(apartmentModel),
		}

		apartments = append(apartments, &apartment)

	}

	return apartments, nil
}
