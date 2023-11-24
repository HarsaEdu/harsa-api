package repository

import "github.com/HarsaEdu/harsa-api/internal/model/domain"

func (repositoryOptions *OptionsRepositoryImpl) Delete(option *domain.Options) error {

	if err := repositoryOptions.DB.Delete(&option).Error; err != nil {
		return  err
	}

	return nil
}