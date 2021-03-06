package repository

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
}

func NewTaskRepository() TaskRepository {
	return &TaskRepositoryImpl{}
}

func (repository *TaskRepositoryImpl) SaveTask(ctx context.Context, db gorm.DB, task request.TaskRq) request.TaskRq {

	query := db.Table("tasks").WithContext(ctx).Create(task)

	if query.Error != nil {
		panic("Error save task to database: " + query.Error.Error())
	}

	return task

}

func (repository *TaskRepositoryImpl) GetTaskById(ctx context.Context, db gorm.DB, id int) (model.Task, error) {

	task := model.Task{}
	query := db.Find(&task, id)

	if query.Error != nil {
		return task, query.Error
	}

	return task, nil
}

func (repository *TaskRepositoryImpl) GetAllTasks(ctx context.Context, db gorm.DB) []model.Task {

	tasks := []model.Task{}
	query := db.Find(&tasks)

	if query.Error != nil {
		panic("error get task: " + query.Error.Error())
	}

	return tasks
}

func (repository *TaskRepositoryImpl) UpdateTask(ctx context.Context, db gorm.DB, task model.Task) (model.Task, error) {

	query := db.Save(&task)

	if query.Error != nil {
		return model.Task{}, query.Error
	}

	return task, nil
}

func (repository *TaskRepositoryImpl) DeleteTask(ctx context.Context, mysql gorm.DB, id int) bool {

	task := model.Task{}
	query := mysql.Delete(&task, id)

	if query.Error != nil {
		panic("error delete task: " + query.Error.Error())
	}

	if query.RowsAffected == 0 {
		return false
	}

	return true
}
