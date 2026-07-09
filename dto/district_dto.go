package dto

import "test-fiber/model"

type District struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetDistrict(model model.District) District {
	return District{
		Id:   model.Id,
		Name: model.Name,
	}
}
