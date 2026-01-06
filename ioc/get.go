package ioc

import (
	"goat/queue"
	"goat/tasks"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type RawUrl struct {
	Url string `json:"url"`
}

func (h IocHandler) GetIOCs(c *gin.Context) {

	client := queue.NewClient()
	defer client.Close()

	var req RawUrl
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed",
		})
		return
	}

	task, err := tasks.NewIocTasks(req.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "failed",
			"error":  err.Error(),
		})
		return
	}

	_, err = client.Enqueue(task, asynq.MaxRetry(5), asynq.Timeout(30*time.Second), asynq.Queue("critical"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "accepted",
	})
}

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
// 	q.Q(len(mappedRes.Data), mappedRes.Data, ctx, collection)
// }
