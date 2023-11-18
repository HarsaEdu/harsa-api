package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func (cloudinaryUploader *CloudinaryUploaderImpl) Uploader(c echo.Context, fileheader, folderName string) (string, error) {

	fileHeader, _ := c.FormFile(fileheader)
	folderPath := fmt.Sprintf("harsa/%s", folderName)

	if fileHeader != nil {
		file, _ := fileHeader.Open()
		ctx := context.Background()
		urlCloudinary := fmt.Sprintf("cloudinary://%s:%s@%s", cloudinaryUploader.Config.ApiKey, cloudinaryUploader.Config.ApiSecret, cloudinaryUploader.Config.CloudName)
		cldService, _ := cloudinary.NewFromURL(urlCloudinary)
		response, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folderPath})
		if err != nil {
			return "", err
		}
		return response.SecureURL, nil
	}
	return "", fmt.Errorf("file not found")
}