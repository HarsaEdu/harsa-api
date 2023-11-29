package conversion

import (
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func ConvertUserCourse(userCourse *domain.UserProfile) *web.UserForCourse {
	name := userCourse.FirstName + " " + userCourse.LastName
	return &web.UserForCourse {
		ID: userCourse.ID,
		Name: name,
		Job: userCourse.Job,
	}
}

func ConvertCourse(course *domain.Course, enrolled int64, totalModul int64) *web.CourseForTraking{
	
	user := ConvertUserCourse(&course.User.UserProfile)

	feedback := ConvertAllFeedBackResponseMobile(course.Feedback)

	return &web.CourseForTraking {
		ID: course.ID,
		Title: course.Title,
		Description: course.Description,
		ImageUrl: course.ImageUrl,
		Rating: course.Rating,
		Enrolled: enrolled,
		TotalModules: totalModul,
		User: *user,
		Feedback: feedback,
		
	}
}

func ConvertUserCourseTraking(userCourse *domain.UserProfile) *web.UserForTracking {
	name := userCourse.FirstName + " " + userCourse.LastName
	return &web.UserForTracking {
		ID: userCourse.ID,
		Name: name,
	}
}

func ConvertCourseTrackingRespose(courseTracking *domain.CourseTracking, enrolled int64, totalModul int64) *web.CourseTrackingResponseMobile {
	
	course := ConvertCourse(&courseTracking.Course, enrolled, totalModul)

	user := ConvertUserCourseTraking(&courseTracking.User.UserProfile)

	module := ConvertAllHistoryModuleResponseMobile(courseTracking.HistoryModule)
	
	return &web.CourseTrackingResponseMobile{
		ID: courseTracking.ID,
		User: *user,
		Progress: courseTracking.Progress,
		Status: courseTracking.Status,
		Course: *course,
		HistoryModule: module,
	}
}

