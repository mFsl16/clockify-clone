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
}
