package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
}

func NewProjectRepository() ProjectRepository {
	return &ProjectRepositoryImpl{}
}

func (repository *ProjectRepositoryImpl) Save(ctx context.Context, db gorm.DB, project model.Project) model.Project {

	result := db.WithContext(ctx).Create(project)
	if result.Error != nil {
		panic("Error insert to database: " + result.Error.Error())
	}

	return project
}
