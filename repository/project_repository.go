package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	Save(ctx context.Context, db gorm.DB, project model.Project) model.Project
}
