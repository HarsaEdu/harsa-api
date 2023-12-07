package service

import "fmt"

func (courseService *CourseServiceImpl) Delete(id uint, userId uint, role string) error {

	courseExist, err := courseService.CourseRepository.CekIdFromCourse(userId, id, role)
	if err != nil { 
		return fmt.Errorf("error when cek id user in course delete : %s", err.Error())
	}

	err = courseService.CourseRepository.Delete(courseExist)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err.Error())
	}

	return nil
}