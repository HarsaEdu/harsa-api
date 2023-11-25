package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CourseDomainToCourseGetAllResponse(courseDomain []domain.CourseEntity) []web.GetCourseResponse {
	var courseGetAllResponse []web.GetCourseResponse

	for _, course := range courseDomain {

		if course.CategoryID != 0 {
			name := course.FirstName + " " + course.LastName
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
					ID:   course.CategoryID,
					Name: course.CategoryName,
				},
				User: &web.UserForCourseResponse{
					ID:    course.UserID,
					Name: name,
					Role: course.RoleName,
				},
			})
		}
	}

	return courseGetAllResponse
}

func CourseDomainToCourseGetByIdResponse(courseDomain *domain.CourseEntity) *web.GetCourseResponse {
	name := courseDomain.FirstName + " " + courseDomain.LastName
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
			ID:   courseDomain.CategoryID,
			Name: courseDomain.CategoryName,
		},
		User: &web.UserForCourseResponse{
			ID:    courseDomain.UserID,
			Name: name,
			Role: courseDomain.RoleName,
		},
	}

	// for _, module := range courseDomain.Modules {
	// 	courseGetResponse.Modules = append(courseGetResponse.Modules, &web.ModulesForCourseResponse{
	// 		ID:          module.ID,
	// 		Title:       module.Title,
	// 		Order: module.Order,
	// 		Type: module.Type,
	// 		SubModules: []*web.SubModulesForModuleForCourseResponse{}, // Initialize an empty slice,
	// 	})

	// 	for _, subModule := range module.SubModules {
	// 		courseGetResponse.Modules[len(courseGetResponse.Modules)-1].SubModules = append(courseGetResponse.Modules[len(courseGetResponse.Modules)-1].SubModules,
	// 		&web.SubModulesForModuleForCourseResponse{
	// 			ID: subModule.ID,
	// 			Title: subModule.Title,
	// 			Type: subModule.Type,
	// 		})
	// 	}
		
	// }
	
	return courseGetResponse
}