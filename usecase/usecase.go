package usecase

import (
	"context"
	"github.com/mFsl16/clockify-clone/model/response"

	"github.com/mFsl16/clockify-clone/model/request"
)

type Usecase interface {
	AddTask(ctx context.Context, task request.TaskRq) response.CommonRs
	AddProject(ctx context.Context, project request.ProjectRq) response.CommonRs
	GetProjectById(ctx context.Context, id int) response.CommonRs
	GetTaskById(ctx context.Context, id int) response.CommonRs
	GetAllProject(ctx context.Context) response.CommonRs
	GetAllTasks(ctx context.Context) response.CommonRs
	UpdateTask(ctx context.Context, id int, taskUpdate request.TaskRq) response.CommonRs
	UpdateProject(ctx context.Context, id int, project request.ProjectRq) response.CommonRs
	DeleteTask(ctx context.Context, id int) response.CommonRs
	DeleteProject(ctx context.Context, id int) response.CommonRs
}
