package usecase

import (
	"context"
	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/response"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"

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

func (usecase *UsecaseImpl) AddTask(ctx context.Context, task request.TaskRq) response.CommonRs {
	commonRs := response.CommonRs{}

	defer func() {
		if r := recover(); r != nil {
			commonRs = commonRs.SetFailed(r)
		}
	}()

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

	saveTask := usecase.TaskRepo.SaveTask(ctx, usecase.DB.Mysql, task)

	commonRs = commonRs.SetSuccess(saveTask)

	return commonRs

}

func (usecase *UsecaseImpl) AddProject(ctx context.Context, project request.ProjectRq) response.CommonRs {

	commonRs := response.CommonRs{}

	defer func() {
		if r := recover(); r != nil {
			commonRs = commonRs.SetFailed(r)
		}
	}()

	saveProject, err := usecase.ProjectRepo.Save(ctx, usecase.DB.Mysql, project)

	if err != nil {
		panic(err.Error())
	}

	commonRs = commonRs.SetSuccess(saveProject)

	return commonRs
}

func (usecase *UsecaseImpl) GetProjectById(ctx context.Context, id int) response.CommonRs {

	commonRs := response.CommonRs{}

	project, errGetProject := usecase.ProjectRepo.GetProjectById(ctx, usecase.DB.Mysql, id)

	if errGetProject != nil {
		return commonRs.SetFailed(errGetProject.Error())
	}

	if (model.Project{}) == project {
		return commonRs.SetFailed("project not found")
	}

	return commonRs.SetSuccess(project)

}

func (usecase *UsecaseImpl) GetTaskById(ctx context.Context, id int) response.CommonRs {

	commonRs := response.CommonRs{}

	getTask, errGetTask := usecase.TaskRepo.GetTaskById(ctx, usecase.DB.Mysql, id)

	if errGetTask != nil {
		return commonRs.SetFailed(errGetTask.Error())
	}

	if (model.Task{}) == getTask {
		return commonRs.SetFailed("task not found")
	}

	return commonRs.SetSuccess(getTask)
}

func (usecase *UsecaseImpl) GetAllProject(ctx context.Context) response.CommonRs {
	commonRs := response.CommonRs{}
	commonRs = commonRs.SetSuccess(usecase.ProjectRepo.GetAllProject(ctx, usecase.DB.Mysql))
	return commonRs
}

func (usecase *UsecaseImpl) GetAllTasks(ctx context.Context) response.CommonRs {
	commonRs := response.CommonRs{}
	commonRs = commonRs.SetSuccess(usecase.TaskRepo.GetAllTasks(ctx, usecase.DB.Mysql))
	return commonRs
}

func (usecase *UsecaseImpl) UpdateTask(ctx context.Context, id int, taskUpdate request.TaskRq) response.CommonRs {

	commonRs := response.CommonRs{}
	task, errGetTask := usecase.TaskRepo.GetTaskById(ctx, usecase.DB.Mysql, id)

	defer func() {
		if r := recover(); r != nil {
			commonRs = commonRs.SetFailed(r)
		}
	}()

	if errGetTask != nil {
		panic(errGetTask.Error())
	}

	if task == (model.Task{}) {
		panic("task not found")
	}

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

	if len(taskUpdate.Tags) > 0 && task.Tags != taskUpdate.Tags {
		task.Tags = taskUpdate.Tags
	}

	logrus.WithFields(
		logrus.Fields{
			"task": task,
		}).Info("[START UPDATE TASK]")
	updateTask, errUpdateTask := usecase.TaskRepo.UpdateTask(ctx, usecase.DB.Mysql, task)
	logrus.WithFields(
		logrus.Fields{
			"task": updateTask,
		}).Info("[COMPLETE UPDATE TASK]")

	if errUpdateTask != nil {
		panic(errUpdateTask)
	}

	commonRs = commonRs.SetSuccess(updateTask)
	return commonRs
}

func (usecase *UsecaseImpl) UpdateProject(ctx context.Context, id int, project request.ProjectRq) response.CommonRs {

	commonRs := response.CommonRs{}
	existProject, err := usecase.ProjectRepo.GetProjectById(ctx, usecase.DB.Mysql, id)

	defer func() {
		if r := recover(); r != nil {
			commonRs = commonRs.SetFailed(r)
		}
	}()

	if err != nil {
		panic(err)
	}

	if existProject == (model.Project{}) {
		panic("project not found")
	}

	if len(project.Name) > 0 && project.Name != existProject.Name {
		existProject.Name = project.Name
	}

	if len(project.Access) > 0 && project.Access != existProject.Access {
		existProject.Access = project.Access
	}

	if len(project.Category) > 0 && project.Category != existProject.Category {
		existProject.Category = project.Category
	}

	if project.Progress != 0 {
		existProject.Progress = project.Progress
	}

	if project.Tracked != 0 {
		existProject.Tracked = project.Tracked
	}

	updateProject, err := usecase.ProjectRepo.UpdateProject(ctx, usecase.DB.Mysql, existProject)

	if err != nil {
		panic(err)
	}

	commonRs = commonRs.SetSuccess(updateProject)

	return commonRs
}

func (usecase *UsecaseImpl) DeleteTask(ctx context.Context, id int) response.CommonRs {

	commonRs := response.CommonRs{}

	defer func() {
		if r := recover(); r != nil {
			commonRs = commonRs.SetFailed(r)
		}
	}()

	taskExist, err := usecase.TaskRepo.GetTaskById(ctx, usecase.DB.Mysql, id)

	if err != nil {
		panic(err)
	}

	if (model.Task{}) == taskExist {
		panic("task with id: " + strconv.Itoa(id) + " not found")
	}

	isDeleteSucces := usecase.TaskRepo.DeleteTask(ctx, usecase.DB.Mysql, id)

	if !isDeleteSucces {
		commonRs = commonRs.SetFailed("Failed delete task unknown error")
	}

	commonRs = commonRs.SetSuccess("success delete task")

	return commonRs
}

func (usecase *UsecaseImpl) DeleteProject(ctx context.Context, id int) response.CommonRs {

	commonRs := response.CommonRs{}

	defer func() {
		if r := recover(); r != nil {
			commonRs = commonRs.SetFailed(r)
		}
	}()

	projectExist, err := usecase.ProjectRepo.GetProjectById(ctx, usecase.DB.Mysql, id)

	if err != nil {
		panic(err)
	}

	if (model.Project{}) == projectExist {
		panic("project not found")
	}

	isDeleteSuccess := usecase.ProjectRepo.DeleteProject(ctx, usecase.DB.Mysql, id)

	if !isDeleteSuccess {
		panic("failed delete project: unknown error")
	}

	commonRs = commonRs.SetSuccess("success delete project")

	return commonRs
}
