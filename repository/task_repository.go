package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model/request"
	"gorm.io/gorm"
)

type TaskRepository interface {
	SaveTask(ctx context.Context, db gorm.DB, task request.TaskRq) request.TaskRq
}
