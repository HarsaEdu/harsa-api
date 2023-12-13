package database

import (
	"fmt"

	"github.com/HarsaEdu/harsa-api/configs"
	"github.com/HarsaEdu/harsa-api/internal/model/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(config *configs.MySQLConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	migration(db)

	return db, nil
}

func migration(db *gorm.DB) {
	db.AutoMigrate(
		&domain.Role{},
		&domain.User{},
		&domain.Category{},
		&domain.UserProfile{},
		&domain.Category{},
		&domain.Course{},
		&domain.Module{},
		&domain.SubModule{},
		&domain.Submissions{},
		&domain.Quizzes{},
		&domain.Questions{},
		&domain.Options{},
		&domain.Faqs{},
		&domain.Feedback{},
		&domain.UserInterest{},
		&domain.SubsPlan{},
		&domain.UserChatTopic{},
		&domain.Submissions{},
		&domain.CourseTracking{},
		&domain.HistorySubModule{},
		&domain.SubmissionAnswer{},
		&domain.HistoryQuiz{},
		&domain.HistoryQuizAnswer{},
		&domain.Subscription{},
		&domain.PaymentHistory{},
		&domain.Certificate{},
	)
}
