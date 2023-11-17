package service

import "fmt"

func (courseService *CourseServiceImpl) Delete(id uint) error {

	IfExist, _ := courseService.CourseRepository.GetById(id)
	if IfExist == nil {
		return fmt.Errorf("course not found")
	}
	err := courseService.CourseRepository.Delete(id)

	if err != nil {
		return fmt.Errorf("error when deleting : %s", err)
	}

	return nil
}