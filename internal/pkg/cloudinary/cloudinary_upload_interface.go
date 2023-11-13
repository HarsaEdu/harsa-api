package cloudinary

import (
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/labstack/echo/v4"
)

type CloudinaryUpdloader interface {
	Uploader(c echo.Context, fileheader, folderName string) string
}

type CloudinaryUpdloaderImpl struct {
	Config *configs.CloudinaryConfig
}

func NewClodinaryUploader(config *configs.CloudinaryConfig) CloudinaryUpdloader {
	return &CloudinaryUpdloaderImpl{
		Config: config,
	}
}
