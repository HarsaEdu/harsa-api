package cloudinary

import (
	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/labstack/echo/v4"
)

type CloudinaryUploader interface {
	Uploader(c echo.Context, fileheader, folderName string, required bool) (string, error)
}

type CloudinaryUploaderImpl struct {
	Config *configs.CloudinaryConfig
}

func NewClodinaryUploader(config *configs.CloudinaryConfig) CloudinaryUploader {
	return &CloudinaryUploaderImpl{
		Config: config,
	}
}

