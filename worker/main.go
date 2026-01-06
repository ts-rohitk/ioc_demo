package main

import (
	"context"
	"fmt"
	"goat/config"
	"goat/queue"
	"goat/tasks"
	"goat/utils/mapping"
	"goat/utils/q"
	"goat/utils/requests"
	"goat/utils/requests/convert"
	"log"

	"github.com/hibiken/asynq"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandleIocTask(ctx context.Context, t *asynq.Task) error {

	url := string(t.Payload())

	exec(ctx, url, queue.MongoClient.Database(config.Cfg.Get("db_name")).Collection("iocs"))

	fmt.Println("in Queue")
	return nil
}

func exec(ctx context.Context, urlRaw string, collection *mongo.Collection) {
	res, err := requests.New(urlRaw).
		Post().
		Headers(
			requests.NewHeader().
				ContentTypeJSON().
				Set("Auth-Key", "ef47b34bfff285fd2045a09559d728a823029a1f6cdc0bfc").
				Set("Accept-Encoding", "gzip,deflate"),
		).
		JSONBody(requests.Dict{
			"query": "get_iocs",
			"days":  7,
		}).
		Send(ctx)

	if err != nil {
		log.Printf("Error sending request: %v", err)
	}

	mappedRes, err := convert.ConvertTo[mapping.RawData](res)
	if err != nil {
		log.Println("Error from converting response", err)
	}

	fmt.Println(mappedRes.QueryStatus)
	fmt.Println(mappedRes.Data)
	q.Q(len(mappedRes.Data), mappedRes.Data, collection)
}

func main() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				log.Printf("Task %s failed: %v", task.Type(), err)
			}),
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(tasks.TypeProcessIOC, HandleIocTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
