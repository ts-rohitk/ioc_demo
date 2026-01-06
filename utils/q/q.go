package q

import (
	"fmt"
	"goat/utils/mapping"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

func Q(taskLength int, data []mapping.RawData, collection *mongo.Collection) {
	var wg sync.WaitGroup
	var batchWg sync.WaitGroup

	numWorker := 10
	numTasks := taskLength

	tasks := make(chan Task, numTasks)
	results := make(chan Result, numTasks)
	processedIOC := make(chan *mapping.IOC, numTasks)

	for w := 0; w < numWorker; w++ {
		wg.Add(1)
		go Workers(w, tasks, results, &wg)
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

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("task id: %d , data: %+v \n", result.Task.Id, result.Data)
		processedIOC <- result.Data
	}

	go func() {
		batchWg.Wait()
		close(processedIOC)
	}()

	fmt.Println("All Processing Completed...!")
}
