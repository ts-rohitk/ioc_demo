package tasks

import "github.com/hibiken/asynq"

func NewIocTasks(url string) (*asynq.Task, error) {
	return asynq.NewTask(TypeProcessIOC, []byte(url)), nil
}

func NewIocUpdateTask() (*asynq.Task, error) {
	return asynq.NewTask(TaskProcesIOCUpdate, nil), nil
}
