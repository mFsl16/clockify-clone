package usecase

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
	"github.com/mFsl16/clockify-clone/model/request"
)

type Usecase interface {
	AddTask(ctx context.Context, task request.TaskRq) request.TaskRq
	AddProject(ctx context.Context, project request.ProjectRq) request.ProjectRq
	GetProjectById(ctx context.Context, id int) model.Project
}
