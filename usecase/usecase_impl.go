package usecase

import (
	"context"
	"time"

	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
	"github.com/mFsl16/clockify-clone/repository"
)

type UsecaseImpl struct {
	DB          *repository.Database
	ProjectRepo repository.ProjectRepository
	TaskRepo    repository.TaskRepository
}

func NewUsecase(db *repository.Database, projectRepo repository.ProjectRepository, taskRepo repository.TaskRepository) Usecase {
	return &UsecaseImpl{
		DB:          db,
		ProjectRepo: projectRepo,
		TaskRepo:    taskRepo,
	}
}

func (usecase *UsecaseImpl) AddTask(ctx context.Context, task request.TaskRq) request.TaskRq {

	startTime, errStartTime := time.Parse("2006-01-02 15:04:05", task.StartTime)

	if errStartTime != nil {
		panic("error parsing start time: " + errStartTime.Error())
	}

	endTime, errorEndTime := time.Parse("2006-01-02 15:04:05", task.EndTime)

	if errorEndTime != nil {
		panic("error parsing end time: " + errorEndTime.Error())
	}

	duration := endTime.UnixMilli() - startTime.UnixMilli()

	task.Duration = duration

	return usecase.TaskRepo.SaveTask(ctx, usecase.DB.Mysql, task)
}

func (usecase *UsecaseImpl) AddProject(ctx context.Context, project request.ProjectRq) request.ProjectRq {

	return usecase.ProjectRepo.Save(ctx, usecase.DB.Mysql, project)
}

func (usecase *UsecaseImpl) GetProjectById(ctx context.Context, id int) model.Project {

	return usecase.ProjectRepo.GetProjectById(ctx, usecase.DB.Mysql, id)
}
