package services

import (
	"test-fiber/dto"
	"test-fiber/repositories"
)

type BuildingService struct {
	buildingRepository repositories.BuildingInterface
}

func NewBuildingService(buildingRep repositories.BuildingInterface) *BuildingService {
	return &BuildingService{
		buildingRepository: buildingRep,
	}
}

func (s *BuildingService) GetBuildings() ([]dto.BuildingResponse, error) {
	models, err := s.buildingRepository.GetBuildings()
	if err != nil {
		return nil, err
	}
	dto := dto.GetBuildingsResponse(models)

	return dto, nil
}

func (s *BuildingService) GetBuildingsDetail() ([]dto.BuildingDetailsResponse, error) {
	models, err := s.buildingRepository.GetBuildingsWithStats()
	if err != nil {
		return nil, err
	}

	stats, err := s.buildingRepository.GetBuildingStatsByBedrooms()
	if err != nil {
		return nil, err
	}
	dto := dto.GetBuildingsDetailResponse(models, stats)

	return dto, nil
}
