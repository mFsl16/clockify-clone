package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model/request"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (repository *TaskRepositoryImpl) SaveTask(ctx context.Context, db gorm.DB, task request.TaskRq) request.TaskRq {

	result := db.Table("tasks").WithContext(ctx).Create(task)

	if result.Error != nil {
		panic("Error save task to database: " + result.Error.Error())
	}

	return task

}
