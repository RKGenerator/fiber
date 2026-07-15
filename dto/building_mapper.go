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
		District:     buildingModel.District.Id,
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
			District:     buildingModel.District.Id,
		}
		buildings = append(buildings, building)
	}
	return buildings
}

func GetBuildingsDetailResponse(buildingModels []*model.BuildingWithStats, roomStatsModels []model.RoomBuldingStats) []BuildingDetailsResponse {
	buildings := make([]BuildingDetailsResponse, 0, len(buildingModels))
	statsMap := make(map[int64][]model.RoomBuldingStats)

	for _, stat := range roomStatsModels {
		statsMap[stat.BuildingId] = append(statsMap[stat.BuildingId], stat)
	}

	for _, buildingModel := range buildingModels {
		building := BuildingDetailsResponse{
			BuildingResponse: BuildingResponse{
				Id:           buildingModel.Id,
				Latitude:     buildingModel.Latitude,
				Longitude:    buildingModel.Longitude,
				Name:         buildingModel.Name,
				ImageUrl:     buildingModel.ImageUrl,
				BuildingType: buildingModel.TypeOfBuilding,
			},
			District:        GetDistrict(buildingModel.District),
			TotalCount:      buildingModel.TotalCount,
			MedianPrice:     buildingModel.MedianPrice,
			MedianPriceSQFT: buildingModel.MedianPriceSQFT,
			MedianAreaSQFT:  buildingModel.MedianAreaSQFT,
			RoomStats:       GetRoomStats(statsMap[buildingModel.Id]),
		}
		buildings = append(buildings, building)
	}
	return buildings
}

func GetRoomStats(roomStatsModels []model.RoomBuldingStats) []RoomStats {
	roomStats := make([]RoomStats, 0, len(roomStatsModels))

	for _, roomStatsModel := range roomStatsModels {
		roomStat := RoomStats{
			Rooms:           roomStatsModel.Rooms,
			Count:           roomStatsModel.Count,
			MedianPrice:     roomStatsModel.MedianPrice,
			MedianPriceSQFT: roomStatsModel.MedianPriceSQFT,
			MedianAreaSQFT:  roomStatsModel.MedianAreaSQFT,
		}
		roomStats = append(roomStats, roomStat)
	}

	return roomStats
}
