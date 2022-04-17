package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (repository *TaskRepositoryImpl) SaveTask(ctx context.Context, db gorm.DB, task model.Task) model.Task {

	result := db.WithContext(ctx).Create(task)

	if result.Error != nil {
		panic("Error save task to database: " + result.Error.Error())
	}

	return task

}
