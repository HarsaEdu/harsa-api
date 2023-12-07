package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerHandler *SubmissionAnswerHandlerImpl) Create(ctx echo.Context) error {
	paramSubmission := ctx.Param("idSubmission")

	idSubmission, err := strconv.Atoi(paramSubmission)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	idUser := ctx.Get("user_id").(uint)

	req := web.SubmissionAnswerRequest{}
	err = ctx.Bind(&req)

	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind submission answer request", err)
	}

	err = submissionAnswerHandler.SubmissionAnswerservice.Create(ctx, &req, idSubmission, int(idUser))
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "file") {
			return res.StatusBadRequest(ctx, "Only accept a PDF file", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create submission answer, something happen", err)
	}
	return res.StatusCreated(ctx, "success to create submission answer", nil, nil)
}
