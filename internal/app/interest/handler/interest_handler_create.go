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

func (handler *InterestHandlerImpl) CreateInterest(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("profile_id"))
	if err != nil {
		return res.StatusInternalServerError(ctx, "failed to convert param id to int: ", err)
	}
	fmt.Println("id", id)

	request := web.InterestRequest{}
	err = ctx.Bind(&request)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to bind request", err)
	}

	err = handler.Service.CreateInterest(uint(id), &request)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "interest exists") {
			return res.StatusAlreadyExist(ctx, "interest already exist", err)
		}
		return res.StatusInternalServerError(ctx, "failed to create interest, something happen", err)

	}
	return res.StatusCreated(ctx, "success to create interest", nil, nil)
}
