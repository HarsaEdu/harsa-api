package service

import (
	"fmt"
	"testing"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
	"github.com/HarsaEdu/harsa-api/tests/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAllByUserId_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)

	// Set up mock expectations for GetDashBoardIntructur
	mockCourseRepo.On("GetDashBoardIntructur", offset, limit, search, userID).Return(&web.DashboardIntructur{
		// Set dashboard details
	}, int64(5), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByUserId(offset, limit, search, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByUserId_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)

	// Set up mock expectations for GetDashBoardIntructur
	mockCourseRepo.On("GetDashBoardIntructur", offset, limit, search, userID).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByUserId(offset, limit, search, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByUserId_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)

	// Set up mock expectations for GetDashBoardIntructur
	mockCourseRepo.On("GetDashBoardIntructur", offset, limit, search, userID).Return(nil, int64(0), fmt.Errorf("error getting courses"))

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByUserId(offset, limit, search, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "error getting courses", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllCourseByUserId_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)

	// Set up mock expectations for GetAllCourseDashBoardIntructur
	mockCourseRepo.On("GetAllCourseDashBoardIntructur", offset, limit, search, userID).Return(&web.DashboardAllCourseIntructur{
		// Set dashboard details
	}, int64(5), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllCourseByUserId(offset, limit, search, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllCourseByUserId_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)

	// Set up mock expectations for GetAllCourseDashBoardIntructur
	mockCourseRepo.On("GetAllCourseDashBoardIntructur", offset, limit, search, userID).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllCourseByUserId(offset, limit, search, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllCourseByUserId_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	userID := uint(1)

	// Set up mock expectations for GetAllCourseDashBoardIntructur
	mockCourseRepo.On("GetAllCourseDashBoardIntructur", offset, limit, search, userID).Return(nil, int64(0), fmt.Errorf("error getting courses"))

	// Call the function you want to test
	result, pagination, err := courseService.GetAllCourseByUserId(offset, limit, search, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "error getting courses", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAll_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "programming"

	course := domain.Course{
		ID: uint(10),
		Title: "Asdasd",
	}
	course2 := domain.Course{
		ID: uint(11),
		Title: "hgjhgj",
	}

	// Set up mock expectations for GetAll
	mockCourseRepo.On("GetAll", offset, limit, search, category).Return([]domain.Course{
		course,course2,
	}, int64(5), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAll(offset, limit, search, category)

	fmt.Println(result)
	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAll_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "programming"

	// Set up mock expectations for GetAll
	mockCourseRepo.On("GetAll", offset, limit, search, category).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAll(offset, limit, search, category)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "programming"

	// Set up mock expectations for GetAll
	mockCourseRepo.On("GetAll", offset, limit, search, category).Return(nil, int64(0), fmt.Errorf("error getting courses"))

	// Call the function you want to test
	result, pagination, err := courseService.GetAll(offset, limit, search, category)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "error getting courses", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByCategory_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := uint(1)

	course := domain.Course{
		ID: uint(10),
		Title: "Asdasd",
	}
	course2 := domain.Course{
		ID: uint(11),
		Title: "hgjhgj",
	}

	// Set up mock expectations for GetAllByCategory
	mockCourseRepo.On("GetAllByCategory", offset, limit, search, category).Return([]domain.Course{
		course,course2,
	}, int64(5), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByCategory(offset, limit, search, category)

	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByCategory_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := uint(1)

	// Set up mock expectations for GetAllByCategory
	mockCourseRepo.On("GetAllByCategory", offset, limit, search, category).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByCategory(offset, limit, search, category)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByCategory_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := uint(1)

	// Set up mock expectations for GetAllByCategory
	mockCourseRepo.On("GetAllByCategory", offset, limit, search, category).Return(nil, int64(0), fmt.Errorf("error getting courses"))

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByCategory(offset, limit, search, category)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "error getting courses", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}


func TestGetAllMyCourse_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "category"
	userID := uint(1)

	// Set up mock expectations for GetAllMyCourse
	mockCourseRepo.On("GetAllbyUserID", offset, limit, search, category, userID).Return([]domain.Course{
		// Set course details
	}, int64(5), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllMyCourse(offset, limit, search, category, userID)

	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllMyCourse_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "category"
	userID := uint(1)

	// Set up mock expectations for GetAllMyCourse
	mockCourseRepo.On("GetAllbyUserID", offset, limit, search, category, userID).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllMyCourse(offset, limit, search, category, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllMyCourse_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "category"
	userID := uint(1)

	// Set up mock expectations for GetAllMyCourse
	mockCourseRepo.On("GetAllbyUserID", offset, limit, search, category, userID).Return(nil, int64(0), fmt.Errorf("error getting courses"))

	// Call the function you want to test
	result, pagination, err := courseService.GetAllMyCourse(offset, limit, search, category, userID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "error getting courses", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByRating_Success(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "category"

	// Set up mock expectations for GetAllByRating
	mockCourseRepo.On("GetAllCourseByRating", offset, limit, search, category).Return([]domain.Course{
		// Set course details
	}, int64(5), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByRating(offset, limit, search, category)

	// Assert the result
	assert.NoError(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, pagination)

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByRating_CourseNotFound(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "category"

	// Set up mock expectations for GetAllByRating
	mockCourseRepo.On("GetAllCourseByRating", offset, limit, search, category).Return(nil, int64(0), nil)

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByRating(offset, limit, search, category)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "course not found", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}

func TestGetAllByRating_Error(t *testing.T) {
	mockCourseRepo := new(mocks.CourseRepository)

	// Create a CourseServiceImpl with the mock repository
	courseService := NewCourseService(mockCourseRepo, nil, nil)

	// Define test data
	offset := 0
	limit := 10
	search := "example"
	category := "category"

	// Set up mock expectations for GetAllByRating
	mockCourseRepo.On("GetAllCourseByRating", offset, limit, search, category).Return(nil, int64(0), fmt.Errorf("error getting courses"))

	// Call the function you want to test
	result, pagination, err := courseService.GetAllByRating(offset, limit, search, category)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, pagination)
	assert.Equal(t, "error getting courses", err.Error())

	// Assert that the expected calls were made
	mockCourseRepo.AssertExpectations(t)
}