package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Save(ctx context.Context, db gorm.DB, project request.ProjectRq) request.ProjectRq
	GetProjectById(ctx context.Context, db gorm.DB, id int) model.Project
	GetAllProject(ctx context.Context, db gorm.DB) []model.Project
	UpdateProject(ctx context.Context, db gorm.DB, project model.Project) model.Project
	DeleteProject(ctx context.Context, mysql gorm.DB, id int) bool
}
