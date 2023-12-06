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
		ImageUrl: userCourse.ImageUrl,
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
		ImageUrl: userCourse.ImageUrl,
	}
}

func ConvertCourseTrakingMobile(courseTracking *domain.CourseTracking, progress float32) *web.GetAllCourseForTraking {
	
	userIntructur := ConvertUserCourse(&courseTracking.Course.User.UserProfile)
	userStudent   := ConvertUserCourseTraking(&courseTracking.User.UserProfile)

	return &web.GetAllCourseForTraking {
			TrakingID: courseTracking.ID,
			CourseID: courseTracking.CourseID,
			UserIntructur: *userIntructur,
			UserStudent: *userStudent,
			ImageUrl: courseTracking.Course.ImageUrl,
			Title: courseTracking.Course.Title,
			Description: courseTracking.Course.Description,
			Status: courseTracking.Status,
			Progress: progress,
			CreatedAt: courseTracking.CreatedAt,
	}
}

func ConvertCourseTrackingRespose(courseTracking *domain.CourseTracking, course *domain.Course,sections []web.SectionResponseMobile, enrolled int64, totalModul int64, progress float32) *web.CourseTrackingResponseMobile {
	
	user := ConvertUserCourseTraking(&courseTracking.User.UserProfile)

	courseRes := ConvertCourse(course, enrolled, totalModul)
	
	courseTrackingRes := &web.CourseTrackingResponse{
		ID: courseTracking.ID,
		User: *user,
		Progress: progress,
		Status: courseTracking.Status,
	}

	return &web.CourseTrackingResponseMobile{
			CourseTracking: *courseTrackingRes,
			Course: *courseRes,
			Sections : sections,
	}
}


func ConvertCourseTrackingResponeWeb(courseTracking *domain.CourseTracking) *web.CourseTrackingResponseWeb{

	return &web.CourseTrackingResponseWeb{
		ID : courseTracking.Course.ID,
		Title: courseTracking.Course.Title,
	}
}

func ConvertAllCourseTrackingResponeWeb(courseTrackings []domain.CourseTracking) []web.CourseTrackingResponseWeb{

	var courseTrackingResponseWebs []web.CourseTrackingResponseWeb

	for _, courseTracking := range courseTrackings {
		courseTrackingResponseWebs = append(courseTrackingResponseWebs, *ConvertCourseTrackingResponeWeb(&courseTracking))
	}

	return courseTrackingResponseWebs
}


func ConvertCourseTrackingUserWeb(courseTracking *domain.CourseTracking) *web.CourseTrackingUserWeb{
	name := courseTracking.User.UserProfile.FirstName + " " + courseTracking.User.UserProfile.LastName
	return &web.CourseTrackingUserWeb{
		CourseTrakingID: courseTracking.ID,
		UserID : courseTracking.User.ID,
		Name: name,
		UserName: courseTracking.User.Username,
		Email: courseTracking.User.Email,
		PhoneNumber: courseTracking.User.UserProfile.PhoneNumber,
		Address: courseTracking.User.UserProfile.Address,
	}
}


func ConvertAllCourseTackingUserWeb(courseTrackings []domain.CourseTracking) []web.CourseTrackingUserWeb{

	var courseTrackingUserWebs []web.CourseTrackingUserWeb

	for _, courseTracking := range courseTrackings {
		courseTrackingUserWebs = append(courseTrackingUserWebs, *ConvertCourseTrackingUserWeb(&courseTracking))
	}

	return courseTrackingUserWebs
}

