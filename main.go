package main

import (
	"goat/config"
	"goat/db"
	"goat/ioc"

	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Load()
	client := db.Connect()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	ioc.AddIocRouter(r, client)

	r.Run(":3000")
}

// func startWorker() {

// 	go func() {
// 		for url := range q {
// 			exec(url, collection)
// 		}
// 	}()
// }

// func exec(urlRaw string, collection *mongo.Collection) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	res, err := requests.New(urlRaw).
// 		Post().
// 		Headers(
// 			requests.NewHeader().
// 				ContentTypeJSON().
// 				Set("Auth-Key", "ef47b34bfff285fd2045a09559d728a823029a1f6cdc0bfc").
// 				Set("Accept-Encoding", "gzip,deflate"),
// 		).
// 		JSONBody(requests.Dict{
// 			"query": "get_iocs",
// 			"days":  7,
// 		}).
// 		Send(ctx)

// 	if err != nil {
// 		log.Printf("Error sending request: %v", err)
// 	}

// 	mappedRes, err := convert.ConvertTo[mapping.RawData](res)
// 	if err != nil {
// 		log.Println("Error from converting response", err)
// 	}

// 	fmt.Println(mappedRes.QueryStatus)
// 	fmt.Println(mappedRes.Data)
// 	q.Q(len(mappedRes.Data), mappedRes.Data, collection)
// }
