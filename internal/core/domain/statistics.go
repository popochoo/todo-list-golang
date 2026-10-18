package domain

import "time"

type Statistics struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	taskCreated int,
	taskCompleted int,
	taskCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreated:               taskCreated,
		TasksCompleted:             taskCompleted,
		TasksCompletedRate:         taskCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
