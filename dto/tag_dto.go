package dto

import "test-fiber/model"

type Tag struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func GetTags(models []model.Tag) []Tag {

	tags := make([]Tag, 0, len(models))

	for _, model := range models {
		tag := Tag{
			Id:   model.Id,
			Name: model.Name,
		}
		tags = append(tags, tag)
	}
	return tags
}
