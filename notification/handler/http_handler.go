package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/personal_proj/notification/kafka"

	"github.com/gin-gonic/gin"
)

//OfferHandlerInterface ...
type OfferHandlerInterface interface {
	PostDataToKafka(c *gin.Context)
}
type offerHandler struct {
}

//NewHandler ...
func NewHandler() OfferHandlerInterface {
	return &offerHandler{}
}

func (album offerHandler) PostDataToKafka(ctx *gin.Context) {
	parent := context.Background()
	defer parent.Done()

	// form := &struct {
	// 	Message map[string]interface{} `json:"message,omitempty"`
	// }{}

	var form map[string]interface{}
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while marshalling json: %s", err.Error()),
			},
		})

	}
	//ctx.Bind(form)
	formInBytes, err := json.Marshal(form)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while marshalling json: %s", err.Error()),
			},
		})

		ctx.Abort()
		return
	}

	err = kafka.Push(parent, nil, formInBytes)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"error": map[string]interface{}{
				"message": fmt.Sprintf("error while push message into kafka: %s", err.Error()),
			},
		})

		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "success push data into kafka",
		"data":    form,
	})
}
