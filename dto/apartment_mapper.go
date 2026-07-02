package dto

import "test-fiber/model"

func GetApartmentsDTO(apartsModel []*model.Apartment) ([]*ApartmentDTO, error) {
	if apartsModel == nil {
		return nil, nil
	}

	var apartments []*ApartmentDTO

	for _, apartmentModel := range apartsModel {
		apartment := ApartmentDTO{
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
