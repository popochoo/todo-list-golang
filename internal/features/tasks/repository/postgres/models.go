package tasks_postgres_repository

import (
	"time"

	"github.com/popochoo/todo-list-golang/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func taskDomainFromModel(taskModel TaskModel) domain.Task {
	return domain.NewTask(
		taskModel.ID,
		taskModel.Version,
		taskModel.Title,
		taskModel.Description,
		taskModel.Completed,
		taskModel.CreatedAt,
		taskModel.CompletedAt,
		taskModel.AuthorUserID,
	)
}

func tasksDomainsFromModels(tasks []TaskModel) []domain.Task {
	tasksDomains := make([]domain.Task, len(tasks))

	for i, task := range tasks {
		tasksDomains[i] = domain.NewTask(
			task.ID,
			task.Version,
			task.Title,
			task.Description,
			task.Completed,
			task.CreatedAt,
			task.CompletedAt,
			task.AuthorUserID,
		)
	}

	return tasksDomains
}
