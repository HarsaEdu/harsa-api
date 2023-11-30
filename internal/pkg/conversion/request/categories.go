package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CategoryCreateRequestToCategoriesModel(request web.CategoryCreateRequest) *domain.Category {
	return &domain.Category{
		Name:        request.Name,
		Description: request.Description,
		Image_url:   "https://placewaifu.com/image/200/200",
	}
}

func CategoryUpdateRequestToCategoriesModel(request web.CategoryUpdateRequest) *domain.Category {
	return &domain.Category{
		Name:        request.Name,
		Description: request.Description,
	}
}

func CategoryUploadImageRequestToCategoriesModel(request web.CategoryUploadImageRequest) *domain.Category {
	return &domain.Category{
		Image_url: request.CategoryImage,
	}
}
