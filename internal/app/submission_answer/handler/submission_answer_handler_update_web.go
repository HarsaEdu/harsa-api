package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (submissionAnswerHandler *SubmissionAnswerHandlerImpl) UpdateWeb(ctx echo.Context) error {
	idParam := ctx.Param("subsAnsId")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return fmt.Errorf("error when convert id to int")
	}

	req := web.SubmissionAnswerUpdateWeb{}
	err = ctx.Bind(&req)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind submission answer update web request", err)
	}

	err = submissionAnswerHandler.SubmissionAnswerservice.UpdateWeb(ctx, &req, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "not found") {
			return res.StatusNotFound(ctx, "submission answer not found", err)
		}

		return res.StatusInternalServerError(ctx, "failed update submission answer web, something happen", err)
	}

	return res.StatusOK(ctx, "success to update submission answer", nil, nil)

}
