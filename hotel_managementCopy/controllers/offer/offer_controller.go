package offer

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/personal_proj/hotel_management/services"
	"github.com/personal_proj/hotel_management/utils/errors"
)

func InsertOffer(c *gin.Context) {
	var offer map[string]interface{}

	if err := c.ShouldBindJSON(&offer); err != nil {
		restErr := errors.NewBadRequestError("Invalid Json Body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.OfferService.InsertOffer(offer)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		//TODO: Handle user creation error
		return
	}
	c.JSON(http.StatusCreated, result)
}
