package dto

import (
	"log"
	"test-fiber/model"
)

func GetApartmentsResponse(apartModels []*model.Apartment) ([]*ApartmentResponse, error) {
	if apartModels == nil {
		return nil, nil
	}

	var apartments []*ApartmentResponse

	for _, apartmentModel := range apartModels {
		log.Println(apartmentModel)
		apartment := ApartmentResponse{
			Id:           apartmentModel.Id,
			SourcePrice:  apartmentModel.SourcePriceAED,
			Address:      apartmentModel.Address,
			PropertyType: apartmentModel.PropertyType,
			Purpose:      apartmentModel.Purpose,
			NBedrooms:    apartmentModel.NBedrooms,
			AreaSQFT:     apartmentModel.AreaSQFT,
			Description:  apartmentModel.Description,
			CreatedAt:    apartmentModel.CreatedAt,
			UpdatedAt:    apartmentModel.UpdatedAt,
			PriceSQFT:    apartmentModel.PriceSQFT,
			Floor:        apartmentModel.Floor,
			Booking:      apartmentModel.Booking,
			ExternalId:   apartmentModel.ExternalId,
			Disabled:     apartmentModel.Disabled}

		apartments = append(apartments, &apartment)
	}

	return apartments, nil

}

func GetApartmentsDetailsResponse(apartModels []*model.Apartment) ([]*ApartmentDetailsResponse, error) {
	if apartModels == nil {
		return nil, nil
	}

	apartments := make([]*ApartmentDetailsResponse, 0, len(apartModels))

	for _, apartmentModel := range apartModels {
		apartment := ApartmentDetailsResponse{
			ApartmentResponse: ApartmentResponse{
				Id:           apartmentModel.Id,
				SourcePrice:  apartmentModel.SourcePriceAED,
				Address:      apartmentModel.Address,
				PropertyType: apartmentModel.PropertyType,
				Purpose:      apartmentModel.Purpose,
				NBedrooms:    apartmentModel.NBedrooms,
				AreaSQFT:     apartmentModel.AreaSQFT,
				Description:  apartmentModel.Description,
				CreatedAt:    apartmentModel.CreatedAt,
				UpdatedAt:    apartmentModel.UpdatedAt,
				PriceSQFT:    apartmentModel.PriceSQFT,
				Floor:        apartmentModel.Floor,
				Booking:      apartmentModel.Booking,
				ExternalId:   apartmentModel.ExternalId,
				Disabled:     apartmentModel.Disabled,
			},
			PropertyImagesResponse: GetImagesResponse(apartmentModel),
			Bulding:                GetBuildingResponse(apartmentModel.Building),
			Tags:                   GetTags(apartmentModel.Tags),
		}

		apartments = append(apartments, &apartment)

	}

	return apartments, nil
}
