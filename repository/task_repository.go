package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	SaveTask(ctx context.Context, db gorm.DB, task model.Task) model.Task
}
