package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerHandler *SubmissionAnswerHandlerImpl) Update(ctx echo.Context) error {
	idParam := ctx.Param("subsAnsId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	idUser := ctx.Get("user_id").(uint)

	req := web.SubmissionAnswerUpdateRequest{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind submission answer update request", err)
	}

	err = submissionAnswerHandler.SubmissionAnswerservice.Update(ctx, &req, id, int(idUser))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "submission answer not found", err)
		}
		if strings.Contains(err.Error(), "file") {
			return res.StatusBadRequest(ctx, "Only accept a PDF file", err)
		}
		return res.StatusInternalServerError(ctx, "failed update submission answer, something happen", err)

	}

	return res.StatusOK(ctx, "success to update submission answer", nil, nil)
}
