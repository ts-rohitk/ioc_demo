package q

import (
	"context"
	"fmt"
	"goat/utils/mapping"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	BatchSize    = 100
	BatchTimeout = 6 * time.Second
)

func BatchWorker(results <-chan *mapping.IOC, collection *mongo.Collection, wg *sync.WaitGroup) {
	defer wg.Done()

	batch := make([]any, 0, BatchSize)
	ticker := time.NewTicker(BatchTimeout)
	defer ticker.Stop()

	totalInserted := 0
	batchCount := 0

	flush := func() {
		if len(batch) == 0 {
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, err := collection.InsertMany(ctx, batch)
		if err != nil {
			fmt.Println("error from batch process", err)
		} else {
			totalInserted += 1
			batchCount += len(batch)
			fmt.Printf("Batch: %d , Inserted %d documents (total: %d)\n", batchCount, len(batch), totalInserted)
		}

		batch = batch[:0]
	}

	for {
		select {
		case ioc, ok := <-results:
			if !ok {
				flush()
				fmt.Printf("Batch Process Finished %d documents \n", batchCount)
				return
			}

			batch = append(batch, ioc)
			if len(batch) >= BatchSize {
				flush()
			}
		case <-ticker.C:
			if len(batch) > 0 {
				flush()
				fmt.Printf("%d documents are created \n", batchCount)
			}
		}
	}
}
