package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/middleware"
	"github.com/labstack/echo/v4"
)

func (certificateRoutes *CertificateRoutesImpl) CertificateMobile(apiGroup *echo.Group){
	certificatesGroup := apiGroup.Group("/certificates")

	certificatesGroup.GET("/:id/download", certificateRoutes.CertificateHandler.GenerateCertificate, middleware.StudentMiddleare)
}