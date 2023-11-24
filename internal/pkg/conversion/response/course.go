package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CourseDomainToCourseGetAllResponse(courseDomain []domain.Course) []web.GetCourseResponse {
	var courseGetAllResponse []web.GetCourseResponse

	for _, course := range courseDomain {

		if course.Category.ID != 0 {
			name := course.User.FirstName + " " + course.User.LastName
			courseGetAllResponse = append(courseGetAllResponse, web.GetCourseResponse{
				ID:          course.ID,
				Title:       course.Title,
				Description: course.Description,
				ImageUrl:    course.ImageUrl,
				Enrolled:    course.Enrolled,
				Rating:      course.Rating,
				CreatedAt:   course.CreatedAt,
				UpdatedAt:   course.UpdatedAt,
				Category: &web.CategoryForCourseResponse{
					ID:   course.Category.ID,
					Name: course.Category.Name,
				},
				User: &web.UserForCourseResponse{
					ID:    course.User.ID,
					Name: name,
					Role: course.User.User.Role.Name,
				},
			})
		}
	}

	return courseGetAllResponse
}

func CourseDomainToCourseGetByIdResponse(courseDomain *domain.Course) *web.GetCourseResponse {
	name := courseDomain.User.FirstName + " " + courseDomain.User.LastName
	courseGetResponse := &web.GetCourseResponse{
		ID:          courseDomain.ID,
		Title:       courseDomain.Title,
		Description: courseDomain.Description,
		ImageUrl:    courseDomain.ImageUrl,
		Enrolled:    courseDomain.Enrolled,
		Rating:      courseDomain.Rating,
		CreatedAt:   courseDomain.CreatedAt,
		UpdatedAt:   courseDomain.UpdatedAt,
		Category: &web.CategoryForCourseResponse{
			ID:   courseDomain.Category.ID,
			Name: courseDomain.Category.Name,
		},
		User: &web.UserForCourseResponse{
			ID:    courseDomain.User.ID,
			Name: name,
			Role: courseDomain.User.User.Role.Name,
		},
	}

	for _, module := range courseDomain.Modules {
		courseGetResponse.Modules = append(courseGetResponse.Modules, &web.ModulesForCourseResponse{
			ID:          module.ID,
			Title:       module.Title,
			Order: module.Order,
			Type: module.Type,
			SubModules: []*web.SubModulesForModuleForCourseResponse{}, // Initialize an empty slice,
		})

		for _, subModule := range module.SubModules {
			courseGetResponse.Modules[len(courseGetResponse.Modules)-1].SubModules = append(courseGetResponse.Modules[len(courseGetResponse.Modules)-1].SubModules,
			&web.SubModulesForModuleForCourseResponse{
				ID: subModule.ID,
				Title: subModule.Title,
				Type: subModule.Type,
			})
		}
		
	}
	return courseGetResponse
}