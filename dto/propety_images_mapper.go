package dto

import (
	"test-fiber/model"
)

func GetImagesResponse(apartsModel *model.Apartment) PropertyImagesResponse {

	urls := make([]*string, 0, len(apartsModel.PropertyImages))
	for _, imageModel := range apartsModel.PropertyImages {
		urls = append(urls, &imageModel.Url)
	}

	return PropertyImagesResponse{Urls: urls}
}
