package q

import (
	"context"
	"goat/updation"
	"goat/updation/builder"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateWorker(updateCh <-chan *UpdateTask, collection *mongo.Collection, updateWg *sync.WaitGroup) {
	defer updateWg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for iocData := range updateCh {
		query := builder.NewSetQueryBuilder()
		if updation.NewMalwareFamily(*iocData.IncomingIOC.Malware[0].Family, *iocData.ExistingIOC.Malware[0].Family) {
			result, err := collection.InsertOne(ctx, iocData.IncomingIOC)
			if err != nil {
				log.Fatal("error in inserting record with new malware family")
			}
			log.Printf("inserted new documented while updating when malware family is found to be new with id %s", result.InsertedID)
		}

		query.AddForUpdate("firstSeen", updation.FirstSeenUpdation(iocData.IncomingIOC.FirstSeen, iocData.ExistingIOC.FirstSeen))
		query.AddForUpdate("lastSeen", updation.LastSeenUpdation(iocData.IncomingIOC.LastSeen, iocData.ExistingIOC.LastSeen))
		query.AddForUpdate("tags", updation.UnionTag(iocData.IncomingIOC.Tags, iocData.ExistingIOC.Tags))

		collection.UpdateOne(
			ctx,
			bson.M{"key": iocData.ExistingIOC.Key},
			bson.M{
				"$set": query.Build(),
			},
		)
	}

}

// func UpdateLogic(ioc *, collection *mongo.Collection) {
// 	updation.NewMalwareFamily(ioc.Malware)
// }
