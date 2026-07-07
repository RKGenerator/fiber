package dto

import (
	"log"
	"test-fiber/model"
)

func ApartmentsDTO(apartsModel []*model.Apartment) ([]*ApartmentDTO, error) {
	if apartsModel == nil {
		return nil, nil
	}

	var apartments []*ApartmentDTO

	for _, apartmentModel := range apartsModel {
		log.Println(apartmentModel)
		apartment := ApartmentDTO{
			Id:             apartmentModel.Id,
			SourcePriceAED: apartmentModel.SourcePriceAED,
			Address:        apartmentModel.Address,
			PropertyType:   apartmentModel.PropertyType,
			Purpose:        apartmentModel.Purpose,
			NBedrooms:      apartmentModel.NBedrooms,
			AreaSQFT:       apartmentModel.AreaSQFT,
			Description:    apartmentModel.Description,
			BuildingID:     apartmentModel.BuildingID,
			CreatedAt:      apartmentModel.CreatedAt,
			PriceSQFT:      apartmentModel.PriceSQFT,
			Floor:          apartmentModel.Floor,
			Booking:        apartmentModel.Booking}

		apartments = append(apartments, &apartment)
	}

	return apartments, nil

}

func ApartmentsDetailsDTO(apartsModel []*model.Apartment) ([]*ApartmentDetailsDTO, error) {
	if apartsModel == nil {
		return nil, nil
	}

	apartments := make([]*ApartmentDetailsDTO, 0, len(apartsModel))

	for _, apartmentModel := range apartsModel {
		apartment := ApartmentDetailsDTO{
			ApartmentDTO: ApartmentDTO{
				Id:             apartmentModel.Id,
				SourcePriceAED: apartmentModel.SourcePriceAED,
				Address:        apartmentModel.Address,
				PropertyType:   apartmentModel.PropertyType,
				Purpose:        apartmentModel.Purpose,
				NBedrooms:      apartmentModel.NBedrooms,
				AreaSQFT:       apartmentModel.AreaSQFT,
				Description:    apartmentModel.Description,
				BuildingID:     apartmentModel.BuildingID,
				CreatedAt:      apartmentModel.CreatedAt,
				PriceSQFT:      apartmentModel.PriceSQFT,
				Floor:          apartmentModel.Floor,
				Booking:        apartmentModel.Booking,
			},
			PropertyImagesDTO: *ImagesDTO(apartmentModel),
		}

		apartments = append(apartments, &apartment)

	}

	return apartments, nil
}
