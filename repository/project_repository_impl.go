package repository

import (
	"context"
	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
}

func NewProjectRepository() ProjectRepository {
	return &ProjectRepositoryImpl{}
}

func (repository *ProjectRepositoryImpl) Save(ctx context.Context, db gorm.DB, project request.ProjectRq) (request.ProjectRq, error) {

	result := db.Table("projects").WithContext(ctx).Create(project)
	if result.Error != nil {
		return project, result.Error
	}

	return project, nil
}

func (repository *ProjectRepositoryImpl) GetProjectById(ctx context.Context, db gorm.DB, id int) (model.Project, error) {

	project := model.Project{}
	result := db.Find(&project, id)

	if result.Error != nil {
		return project, result.Error
	}

	return project, nil
}

func (repository *ProjectRepositoryImpl) GetAllProject(ctx context.Context, db gorm.DB) []model.Project {

	projects := []model.Project{}
	result := db.Find(&projects)

	if result.Error != nil {
		panic("Error get project: " + result.Error.Error())
	}

	return projects
}

func (repository *ProjectRepositoryImpl) UpdateProject(ctx context.Context, db gorm.DB, project model.Project) (model.Project, error) {

	query := db.Save(&project)

	if query.Error != nil {
		return model.Project{}, query.Error
	}

	return project, nil
}

func (repository *ProjectRepositoryImpl) DeleteProject(ctx context.Context, db gorm.DB, id int) bool {

	project := model.Project{}
	query := db.Delete(&project, id)

	if query.Error != nil {
		panic("error delete project: " + query.Error.Error())
	}

	if query.RowsAffected == 0 {
		return false
	}

	return true
}
