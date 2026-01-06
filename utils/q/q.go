package q

import (
	"context"
	"fmt"
	"goat/utils/mapping"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

func Q(taskLength int, data []mapping.RawData, ctx context.Context, collection *mongo.Collection) {
	var wg sync.WaitGroup
	var batchWg sync.WaitGroup
	var updateWg sync.WaitGroup

	numWorker := 10
	numTasks := taskLength

	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)
	processedIOC := make(chan *mapping.IOC, numTasks)
	updateIOC := make(chan *UpdateTask)

	for w := 0; w < numWorker; w++ {
		wg.Add(1)
		go Workers(w, tasks, results, updateIOC, &wg, ctx, collection)
	}

	for t := 0; t < numTasks; t++ {
		tasks <- Task{
			Id:   t,
			Data: data[t],
		}
	}
	close(tasks)

	batchWg.Add(1)
	go BatchWorker(processedIOC, collection, &batchWg)

	updateWg.Add(1)
	go UpdateWorker(updateIOC, collection, &updateWg)

	go func() {
		wg.Wait()
		close(results)
		close(updateIOC)
	}()

	for result := range results {
		fmt.Printf("task id: %d , data: %+v \n", result.Task.Id, result.Data)
		processedIOC <- result.Data
	}
	close(processedIOC)

	batchWg.Wait()
	updateWg.Wait()

	fmt.Println("All Processing Completed...!")
}
