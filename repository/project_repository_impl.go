package repository

import (
	"context"
	"strconv"

	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
	"gorm.io/gorm"
)

type ProjectRepositoryImpl struct {
}

func NewProjectRepository() ProjectRepository {
	return &ProjectRepositoryImpl{}
}

func (repository *ProjectRepositoryImpl) Save(ctx context.Context, db gorm.DB, project request.ProjectRq) request.ProjectRq {

	result := db.Table("projects").WithContext(ctx).Create(project)
	if result.Error != nil {
		panic("Error insert to database: " + result.Error.Error())
	}

	return project
}

func (repository *ProjectRepositoryImpl) GetProjectById(ctx context.Context, db gorm.DB, id int) model.Project {

	project := model.Project{}
	result := db.Find(&project, id)

	if result.Error != nil {
		panic("Error get project by Id: " + strconv.Itoa(id) + " | message: " + result.Error.Error())
	}

	return project
}

func (repository *ProjectRepositoryImpl) GetAllProject(ctx context.Context, db gorm.DB) []model.Project {

	projects := []model.Project{}
	result := db.Find(&projects)

	if result.Error != nil {
		panic("Error get project: " + result.Error.Error())
	}

	return projects
}
