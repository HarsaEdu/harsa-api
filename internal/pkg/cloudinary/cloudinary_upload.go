package cloudinary

import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

func (cloudinaryUpdloader *CloudinaryUpdloaderImpl) Uploader(c echo.Context, fileheader, folderName string) (string, error) {

	fileHeader, _ := c.FormFile(fileheader)
	folderPath := fmt.Sprintf("harsa/%s", folderName)

	if fileHeader != nil {
		file, _ := fileHeader.Open()
		ctx := context.Background()
		urlCloudinary := cloudinaryUpdloader.Config.Url
		cldService, _ := cloudinary.NewFromURL(urlCloudinary)
		response, err := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: folderPath})
		if err != nil {
			return "", err
		}
		return response.SecureURL, nil
	}
	return "", fmt.Errorf("file not found")
}