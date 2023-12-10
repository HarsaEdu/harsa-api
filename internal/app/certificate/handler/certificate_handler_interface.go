package handler

import (
	"github.com/HarsaEdu/harsa-api/internal/app/certificate/service"
	"github.com/labstack/echo/v4"
)

type CertificateHandler interface {
	GenerateCertificate(ctx echo.Context) error
}

type CertificateHandlerImpl struct {
	CertificateService service.CertificateService
}

func NewCertificateHandler(certificateService service.CertificateService) CertificateHandler {
	return &CertificateHandlerImpl{
		CertificateService: certificateService,
	}
}