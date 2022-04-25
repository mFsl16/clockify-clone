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

func (usecase *UsecaseImpl) GetTaskById(ctx context.Context, id int) model.Task {

	return usecase.TaskRepo.GetTaskById(ctx, usecase.DB.Mysql, id)
}

func (usecase *UsecaseImpl) GetAllProject(ctx context.Context) []model.Project {

	return usecase.ProjectRepo.GetAllProject(ctx, usecase.DB.Mysql)
}

func (usecase *UsecaseImpl) GetAllTasks(ctx context.Context) []model.Task {

	return usecase.TaskRepo.GetAllTasks(ctx, usecase.DB.Mysql)
}

func (usecase *UsecaseImpl) UpdateTask(ctx context.Context, id int, taskUpdate request.TaskRq) model.Task {

	task := usecase.GetTaskById(ctx, id)

	if len(taskUpdate.Date) > 0 {
		dateUpdate, err := time.Parse("2006-01-02 15:04:05", taskUpdate.Date)
		task.Date = dateUpdate
		if err != nil {
			panic("Error parse date: " + err.Error())
		}
	}

	if len(taskUpdate.StartTime) > 0 {
		startTimeUpdate, errStartTime := time.Parse("2006-01-02 15:04:05", taskUpdate.StartTime)
		task.StartTime = startTimeUpdate
		if errStartTime != nil {
			panic("Error parse date: " + errStartTime.Error())
		}
	}

	if len(taskUpdate.EndTime) > 0 {
		endDateUpdate, errEndDate := time.Parse("2006-01-02 15:04:05", taskUpdate.EndTime)
		task.EndTime = endDateUpdate
		if errEndDate != nil {
			panic("Error parse date: " + errEndDate.Error())
		}
	}

	if len(taskUpdate.Title) > 0 && task.Title != taskUpdate.Title {
		task.Title = taskUpdate.Project
	}

	if len(task.Project) > 0 && task.Project != taskUpdate.Project {
		task.Project = taskUpdate.Project
	}

	if task.Billable != taskUpdate.Billable {
		task.Billable = taskUpdate.Billable
	}

	if task.Duration != taskUpdate.Duration {
		task.Duration = taskUpdate.Duration
	}

	if len(task.Project) > 0 && task.Project != taskUpdate.Project {
		task.Project = taskUpdate.Project
	}

	if len(task.Tags) > 0 && task.Tags != taskUpdate.Tags {
		task.Tags = taskUpdate.Tags
	}

	return usecase.TaskRepo.UpdateTask(ctx, usecase.DB.Mysql, task)
}
