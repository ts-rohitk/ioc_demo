package queue

import "github.com/hibiken/asynq"

func NewClient() *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: "localhost:6379",
	})
}
