package app

import (
	"github.com/personal_proj/notification/handler"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication ...
func StartApplication() {
	handler := handler.NewHandler()
	router.POST("/sendnotification", handler.PostDataToKafka)
	router.Run(":9000")

}
