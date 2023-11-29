package repository


import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (courseTrackingrepository *CourseTrackingRepositoryImpl) Create(enrolled *domain.CourseTracking) error {

	tx := courseTrackingrepository.DB.Begin()

	if err := tx.Create(&enrolled).Error; err != nil {
		tx.Rollback()
		return err
	}

	var course domain.Course

	if err := tx.
    Preload("Modules").
    Preload("Modules.SubModules").
    Preload("Modules.Submissions").
    Preload("Modules.Quizzes").
    First(&course, enrolled.CourseID).Error; err != nil {
    return err
    }

	for _, Module := range course.Modules {

		historyModule := domain.HistoryModule{
				ModuleID: Module.ID,
				UserID: enrolled.UserID,
				CourseTrakingID: enrolled.ID,
				Progress: 0,
		}

        if err := tx.Create(&historyModule).Error; err != nil {
			tx.Rollback()
			return err
		}

        for _, SubModule := range Module.SubModules {
			historySubModule := domain.HistorySubModule{
                SubModuleID: SubModule.ID,
				UserID: enrolled.UserID,
                HistoryModuleID: historyModule.ID,
                IsComplete: false,
            }

            if err := tx.Create(&historySubModule).Error; err != nil {
				tx.Rollback()
				return err
			}
        }
		
		for _, Submission := range Module.Submissions {
			historySubModule := domain.SubmissionAnswer{
                UserID: enrolled.UserID,
				SubmissionID: Submission.ID,
				HistoryModuleID: historyModule.ID,
				Status: "not_submitted",
            }

            if err := tx.Create(&historySubModule).Error; err != nil {
				tx.Rollback()
				return err
			}
        }

		for _, Quizzes := range Module.Quizzes {
			historyQuiz:= domain.HistoryQuiz{
                UserID: enrolled.UserID,
				QuizID: Quizzes.ID,
				HistoryModuleID: historyModule.ID,
				Score: 0,
				IsComplete: false,
            }

            if err := tx.Create(&historyQuiz).Error; err != nil {
				tx.Rollback()
				return err
			}
        }
		
    }

	tx.Commit()

	return nil
}