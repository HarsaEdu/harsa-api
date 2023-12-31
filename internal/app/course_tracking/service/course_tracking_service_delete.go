package service

import "fmt"

func (courseTrackingService *CourseTrackingServiceImpl) Delete(courseTrackingId uint, courseId uint,userId uint, role string) error {

	trackingExist, err := courseTrackingService.CourseTrackingRepository.CekIdFromCourse(userId, courseTrackingId, role)
	if err != nil {
		return fmt.Errorf("error when cek id user in enrolled delete : %s", err.Error())
	}

	err = courseTrackingService.CourseTrackingRepository.Delete(trackingExist)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err.Error())
	}

	return nil
}