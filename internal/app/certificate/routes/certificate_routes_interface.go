package routes

import (
	"github.com/HarsaEdu/harsa-api/internal/app/certificate/handler"
	"github.com/labstack/echo/v4"
)

type CertificateRoutes interface {
	CertificateMobile(apiGroup *echo.Group)
}

type CertificateRoutesImpl struct {
	CertificateHandler handler.CertificateHandler
}

func NewCertificateRoutes(CertificateHandler handler.CertificateHandler) CertificateRoutes {
	return &CertificateRoutesImpl{
		CertificateHandler: CertificateHandler,
	}
}