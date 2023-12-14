package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (CertificateHandler *CertificateHandlerImpl) GenerateCertificate(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	userId := ctx.Get("user_id").(uint)

	certificatePdf, certificate, err := CertificateHandler.CertificateService.GenerateCertificate(uint(id), userId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusBadRequest(ctx, "course tracking not found", err)
		}

		if strings.Contains(err.Error(), "not completed") {
			return res.StatusBadRequest(ctx, "course not completed", err)
		}

		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx, "user unauthorized", err)
		}

		return res.StatusInternalServerError(ctx, "failed to generate certificate", err)
	}

	certificateAttachment := fmt.Sprintf("attachment; filename=%s.pdf", certificate.ID)

	ctx.Response().Header().Set("Content-Type", "application/pdf")
	ctx.Response().Header().Set("Content-Disposition", certificateAttachment)
	ctx.Response().Write(certificatePdf)

	return res.StatusOK(ctx, "certificate generated", certificate, nil)
}