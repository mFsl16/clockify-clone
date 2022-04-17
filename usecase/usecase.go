package usecase

import (
	"context"

	"github.com/mFsl16/clockify-clone/model"
)

type Usecase interface {
	AddTask(ctx context.Context, task model.Task) model.Task
	AddProject(ctx context.Context, project model.Project) model.Project
}
