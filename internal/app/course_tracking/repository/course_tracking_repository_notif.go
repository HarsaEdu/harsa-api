package repository

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"github.com/HarsaEdu/harsa-api/internal/model/web"
)

func (courseTrackingrepository *CourseTrackingRepositoryImpl) NotifEnrolled(userId uint, courseId uint) (*web.NotificationPersonal, error) {

    user := domain.User{}
    course := domain.Course{}

    // Mengambil data pengguna (user) berdasarkan ID
    if err := courseTrackingrepository.DB.First(&user, userId).Preload("UserProfile").Error; err != nil {
        return nil, err
    }

    // Mengambil data kursus (course) berdasarkan ID
    if err := courseTrackingrepository.DB.First(&course, courseId).Error; err != nil {
        return nil, err
    }

    // Membuat pesan notifikasi
    title := "Pendaftaran Kelas Baru"
    message := fmt.Sprintf("Halo %s! Selamat, Anda baru saja mendaftar untuk kelas %s.", user.UserProfile.FirstName, course.Title)

    // Membuat objek notifikasi personal
    notification := web.NotificationPersonal{
        Title:             title,
        Message:           message,
        RegistrationToken: user.RegistrationToken,
    }

    notif := domain.Notification{
        UserID: user.ID,
        Title: title,
        Content: message,
    }

    result := courseTrackingrepository.DB.Create(&notif)
	if result.Error != nil {
		return  nil, result.Error
	}

    return &notification, nil
}

func (courseTrackingrepository *CourseTrackingRepositoryImpl) NotifEnrolledWeb(userId uint, courseId uint) (*web.NotificationPersonal, error) {

    user := domain.User{}
    course := domain.Course{}

    // Mengambil data pengguna (user) berdasarkan ID
    if err := courseTrackingrepository.DB.First(&user, userId).Preload("UserProfile").Error; err != nil {
        return nil, err
    }

    // Mengambil data kursus (course) berdasarkan ID
    if err := courseTrackingrepository.DB.First(&course, courseId).Error; err != nil {
        return nil, err
    }

    // Membuat pesan notifikasi
    title := "Pendaftaran Kelas Baru"
    message := fmt.Sprintf("Halo %s! Selamat, Anda baru saja didaftarkan untuk mengikuti kelas %s.", user.UserProfile.FirstName, course.Title)

    // Membuat objek notifikasi personal
    notification := web.NotificationPersonal{
        Title:             title,
        Message:           message,
        RegistrationToken: user.RegistrationToken,
    }

    notif := domain.Notification{
        UserID: user.ID,
        Title: title,
        Content: message,
    }

    result := courseTrackingrepository.DB.Create(&notif)
	if result.Error != nil {
		return  nil, result.Error
	}

    return &notification, nil
}

