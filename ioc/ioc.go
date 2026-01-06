package ioc

import (
	"goat/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type IocHandler struct {
	mongo *mongo.Collection
}

func AddIocRouter(r *gin.Engine, mongo *mongo.Client) *gin.RouterGroup {
	collection := mongo.Database(config.Cfg.Get("db_name")).Collection("iocs")

	h := &IocHandler{
		mongo: collection,
	}

	iocRouter := r.Group("/ioc/")
	iocRouter.POST("/", h.GetIOCs)

	return iocRouter
}
