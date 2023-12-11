package test

import (
	"testing"
	"time"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserRepository_UserAvailable(t *testing.T) {
	mockRepo := new(MockUserRepository)

	// Pengaturan expectasi panggilan fungsi UserAvailable
	expectedUser := &domain.User{
		ID:                1,
		Email:             "example@example.com",
		Username:          "example_user",
		Password:          "example_password",
		RoleID:            1,
		RegistrationToken: "registration_token",
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
		DeletedAt:         gorm.DeletedAt{},
		Role:              domain.Role{ID: 1, Name: "student"},
		UserProfile: domain.UserProfile{ID: 1,
			UserID:      1,
			ImageUrl:    "https://robohash.org/74.png?size=200x200",
			FirstName:   "John",
			LastName:    "Doe",
			DateBirth:   time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			Bio:         "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			Gender:      "m",
			PhoneNumber: "123456789",
			City:        "Example City",
			Address:     "123 Example Street",
			Job:         "Software Developer",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now()},
	}
	t.Run("User Available", func(t *testing.T) {
		mockRepo.On("UserAvailable", "username", "email").Return(expectedUser, nil)

		// Pengujian
		user, err := mockRepo.UserAvailable("username", "email")

		// Verifikasi panggilan fungsi telah sesuai dengan ekspektasi
		mockRepo.AssertExpectations(t)

		// Periksa hasil pengujian
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
	})
	t.Run("User Available By ID", func(t *testing.T) {
		mockRepo.On("UserAvailableByID", uint(1)).Return(expectedUser, nil)

		user, err := mockRepo.UserAvailableByID(1)

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, expectedUser, user)
	})
	t.Run("User Profile Create", func(t *testing.T) {
		userProfile := &domain.UserProfile{}
		mockRepo.On("UserProfileCreate", userProfile).Return(nil)

		err := mockRepo.UserProfileCreate(userProfile)

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
	})

	t.Run("User Create", func(t *testing.T) {
		mockRepo.On("UserCreate", expectedUser).Return(expectedUser, nil)
		createdUser, err := mockRepo.UserCreate(expectedUser)

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, expectedUser, createdUser)
	})
	t.Run("User Update", func(t *testing.T) {
		mockRepo.On("UserUpdate", expectedUser).Return(nil)
		err := mockRepo.UserUpdate(expectedUser)

		mockRepo.AssertExpectations(t)

		assert.NoError(t, err)
	})
}
