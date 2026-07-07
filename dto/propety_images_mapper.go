package dto

import (
	"log"
	"test-fiber/model"
)

func ImagesDTO(apartsModel *model.Apartment) *PropertyImagesResponse {
	if apartsModel == nil {
		log.Println("huy")
		return nil
	}
	urls := make([]*string, 0, len(apartsModel.PropertyImages))
	for _, imageModel := range apartsModel.PropertyImages {
		urls = append(urls, &imageModel.Url)
	}

	return &PropertyImagesResponse{Urls: urls}
}
