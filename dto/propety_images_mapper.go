package dto

import (
	"log"
	"test-fiber/model"
)

func ImagesDTO(apartsModel *model.Apartment) *PropertyImagesDTO {
	if apartsModel == nil {
		log.Println("huy")
		return nil
	}
	urls := make([]*string, 0, len(apartsModel.PropertyImages))
	for _, imageModel := range apartsModel.PropertyImages {
		log.Println(imageModel.Url)
		urls = append(urls, &imageModel.Url)
	}

	return &PropertyImagesDTO{Urls: urls}
}
