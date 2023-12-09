package handler

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/internal/pkg/res"
	"github.com/HarsaEdu/harsa-api/internal/pkg/validation"
	"github.com/labstack/echo/v4"
)

func (moduleHandler ModuleHandlerImpl) CreateSection(ctx echo.Context) error {
	paramCourseId := ctx.Param("courseId")
	courseId, err := strconv.Atoi(paramCourseId)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to convert course id : ", err)
	}

	id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	sectionCreateRequest := web.SectionRequest{}
	err = ctx.Bind(&sectionCreateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request invalid", err)
	}

	var countVideo int = 1
	var countPPT int = 1
	youtubePattern := regexp.MustCompile(`(https?://)?(www\.)?(youtube|youtu|youtube-nocookie)\.(com|be)/.+`)
	googleSlidePattern := regexp.MustCompile(`(https?://)?(docs\.google\.com/presentation/d/[\w-]+)/?.*`)
	
	subModule := []domain.SubModule{}

	for _, submodule := range sectionCreateRequest.Modules.SubModules {
		matchYoutube := youtubePattern.MatchString(submodule.ContentUrl)
		matchGoogleSlide := googleSlidePattern.MatchString(submodule.ContentUrl)
	
		if matchYoutube && !matchGoogleSlide {
			submodule.Type = "video"
			submodule.Title = sectionCreateRequest.Modules.Title + " video - " + strconv.Itoa(countVideo)
			fmt.Println(submodule.Title)
			countVideo++
		} else if matchGoogleSlide && !matchYoutube {
			submodule.Type = "ppt"
			submodule.Title = sectionCreateRequest.Modules.Title + " ppt - " + strconv.Itoa(countPPT)
			fmt.Println(submodule.Title)
			countPPT++
		} else {
			return res.StatusBadRequest(ctx, "link not valid", fmt.Errorf("link not youtube or google slide"))
		}
		subModule = append(subModule, submodule)
	}

	sectionCreateRequest.Modules.SubModules = subModule
	
	err = moduleHandler.ModuleService.CreateSection(&sectionCreateRequest, uint(courseId), uint(id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot create this section" ,err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, err.Error(), err)
		}
		return res.StatusInternalServerError(ctx, "failed to create module, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create module", nil, nil)
}

func (moduleHandler ModuleHandlerImpl) CreateModule(ctx echo.Context) error {
	paramSectionId := ctx.Param("sectionId")
	sectionId, err := strconv.Atoi(paramSectionId)
	if err != nil {
		return res.StatusBadRequest(ctx, "failed to convert section id : ", err)
	}

	id := ctx.Get("user_id").(uint)

	roleInterface := ctx.Get("role_name")

	roleString := fmt.Sprintf("%s", roleInterface)

	moduleCreateRequest := web.ModuleRequest{}
	err = ctx.Bind(&moduleCreateRequest)
	if err != nil {
		return res.StatusBadRequest(ctx, "data request invalid", err)
	}

	var countVideo int = 1
	var countPPT int = 1
	youtubePattern := regexp.MustCompile(`(https?://)?(www\.)?(youtube|youtu|youtube-nocookie)\.(com|be)/.+`)
	googleSlidePattern := regexp.MustCompile(`(https?://)?(docs\.google\.com/presentation/d/[\w-]+)/?.*`)
	
	subModule := []domain.SubModule{}

	for _, submodule := range moduleCreateRequest.SubModules {
		matchYoutube := youtubePattern.MatchString(submodule.ContentUrl)
		matchGoogleSlide := googleSlidePattern.MatchString(submodule.ContentUrl)
	
		if matchYoutube && !matchGoogleSlide {
			submodule.Type = "video"
			submodule.Title = moduleCreateRequest.Title + " video - " + strconv.Itoa(countVideo)
			fmt.Println(submodule.Title)
			countVideo++
		} else if matchGoogleSlide && !matchYoutube {
			submodule.Type = "ppt"
			submodule.Title = moduleCreateRequest.Title + " ppt - " + strconv.Itoa(countPPT)
			fmt.Println(submodule.Title)
			countPPT++
		} else {
			return res.StatusBadRequest(ctx, "link not valid", fmt.Errorf("link not youtube or google slide"))
		}
		subModule = append(subModule, submodule)
	}
	moduleCreateRequest.SubModules = subModule


	err = moduleHandler.ModuleService.CreateModule(&moduleCreateRequest, uint(sectionId), uint(id), roleString)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return validation.ValidationError(ctx, err)
		}
		if strings.Contains(err.Error(), "unauthorized") {
			return res.StatusUnauthorized(ctx,"you cannot create this module" ,err)
		}
		if strings.Contains(err.Error(), "already exist") {
			return res.StatusAlreadyExist(ctx, err.Error(), err)
		}
		return res.StatusInternalServerError(ctx, "failed to create module, something happen", err)

	}

	return res.StatusCreated(ctx, "success to create module", nil, nil)
}