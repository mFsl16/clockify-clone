package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Save(ctx context.Context, db gorm.DB, project request.ProjectRq) (request.ProjectRq, error)
	GetProjectById(ctx context.Context, db gorm.DB, id int) (model.Project, error)
	GetAllProject(ctx context.Context, db gorm.DB) []model.Project
	UpdateProject(ctx context.Context, db gorm.DB, project model.Project) (model.Project, error)
	DeleteProject(ctx context.Context, mysql gorm.DB, id int) bool
}
