package handler

import (
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerHandler *SubmissionAnswerHandlerImpl) Get(ctx echo.Context) error {

	idSubmissionParam := ctx.Param("idSubmission")
	idSubmission, err := strconv.Atoi(idSubmissionParam)
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}

	params := ctx.QueryParams()
	offset, err := strconv.Atoi(params.Get("offset"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid offset", err)
	}

	limit, err := strconv.Atoi(params.Get("limit"))
	if err != nil {
		return res.StatusBadRequest(ctx, "invalid limit", err)
	}

	answer, pagination, err := submissionAnswerHandler.SubmissionAnswerservice.Get(offset, limit, idSubmission, params.Get("search"))
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNoContent(ctx, "submission answer not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed to get all submission answer , something happen", err)
	}
	return res.StatusOK(ctx, "success get submissions", answer, pagination)
}
