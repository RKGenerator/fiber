package dto

import (
	"test-fiber/model"
)

func GetBuildingResponse(buildingModel model.Building) BuildingResponse {

	return BuildingResponse{
		Id:           buildingModel.Id,
		Latitude:     buildingModel.Latitude,
		Longitude:    buildingModel.Longitude,
		Name:         buildingModel.Name,
		ImageUrl:     buildingModel.ImageUrl,
		BuildingType: buildingModel.TypeOfBuilding,
		District:     GetDistrict(buildingModel.District),
	}
}

func GetBuildingsResponse(buildingModels []*model.Building) []BuildingResponse {
	buildings := make([]BuildingResponse, 0, len(buildingModels))
	for _, buildingModel := range buildingModels {
		building := BuildingResponse{
			Id:           buildingModel.Id,
			Latitude:     buildingModel.Latitude,
			Longitude:    buildingModel.Longitude,
			Name:         buildingModel.Name,
			ImageUrl:     buildingModel.ImageUrl,
			BuildingType: buildingModel.TypeOfBuilding,
			District:     GetDistrict(buildingModel.District),
		}
		buildings = append(buildings, building)
	}
	return buildings
}
