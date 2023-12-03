package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func CourseDomainToCourseGetAllResponse(courseDomain []domain.Course, countUser int64, countUserActive int64) []web.CourseResponseForIntructur {
	var courseResponseForIntructur []web.CourseResponseForIntructur

	for _, course := range courseDomain {
		feedback := ConvertAllFeedBackResponseMobile(course.Feedback)
		courseResponseForIntructur = append(courseResponseForIntructur, web.CourseResponseForIntructur{
			ID:          course.ID,
			Title:       course.Title,
			ImageUrl:    course.ImageUrl,
			FeedBack:    feedback,
			TotalUser:   countUser,
			TotalActiveUser: countUserActive,
		})
	}

	return courseResponseForIntructur
}

func CourseDomainToCourseGetResponse(course *domain.Course, countUser int64, countUserActive int64) *web.CourseResponseForIntructur {
	var courseResponseForIntructur *web.CourseResponseForIntructur

	feedback := ConvertAllFeedBackResponseMobile(course.Feedback)
	courseResponseForIntructur = &web.CourseResponseForIntructur{
		ID:          course.ID,
		Title:       course.Title,
		ImageUrl:    course.ImageUrl,
		FeedBack:    feedback,
		TotalUser:   countUser,
		TotalActiveUser: countUserActive,
	}

	return courseResponseForIntructur
}

func ConvertAllCourseIntructure(course *domain.Course) *web.AllCourseResponseForIntructur {
	var courseResponseForIntructur *web.AllCourseResponseForIntructur

	courseResponseForIntructur = &web.AllCourseResponseForIntructur{
		ID:          course.ID,
		Title:       course.Title,
		ImageUrl:    course.ImageUrl,
	}

	return courseResponseForIntructur
}

func AllCourseResponseForIntructur(course []web.AllCourseResponseForIntructur, firstName string) *web.DashboardAllCourseIntructur{

	return &web.DashboardAllCourseIntructur{
		FirstName: firstName,
		Course: course,
	}
}

func ConvertCourseResponse(course *domain.Course) *web.CourseResponse {
	return &web.CourseResponse{
		ID:          course.ID,
		Title:       course.Title,
		Description: course.Description,
		UserID:      course.UserID,
	}
}

func CourseResponseForIntructur(course []web.CourseResponseForIntructur, firstName string) *web.DashboardIntructur{

	return &web.DashboardIntructur{
		FirstName: firstName,
		Course: course,
	}
}

func CourseDomainToCourseGetByIdResponse(courseDomain *domain.CourseEntity, module int64) *web.GetCourseResponseById {
	name := courseDomain.FirstName + " " + courseDomain.LastName
	courseGetResponse := &web.GetCourseResponseById{
		ID:          courseDomain.ID,
		Title:       courseDomain.Title,
		Description: courseDomain.Description,
		ImageUrl:    courseDomain.ImageUrl,
		Enrolled:    courseDomain.Enrolled,
		Rating:      courseDomain.Rating,
		TotalModules: module,
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

func CourseDomainToCourseGetAllResponseMobile(courseDomain []domain.Course) []web.GetCourseResponseMobile {
	var courseGetAllResponse []web.GetCourseResponseMobile

	for _, course := range courseDomain {

		if course.CategoryID != 0 {
			name := course.User.UserProfile.FirstName + " " + course.User.UserProfile.LastName
			courseGetAllResponse = append(courseGetAllResponse, web.GetCourseResponseMobile{
				ID:          course.ID,
				Title:       course.Title,
				Description: course.Description,
				ImageUrl:    course.ImageUrl,
				Rating:      course.Rating,
				Category: &web.CategoryForCourseResponse{
					ID:   course.Category.ID,
					Name: course.Category.Name,
				},
				User: &web.UserForCourseResponseMobile{
					ID:    course.UserID,
					Name: name,
				},
			})
		}
	}

	return courseGetAllResponse
}