package app

import (
	"github.com/personal_proj/hotel_management/controllers/offer"
)

func mapUrls() {

	router.POST("/offer", offer.InsertOffer)
	// router.POST("/users", user.Create)
	// router.GET("/users/:user_id", user.Get)
	// //router.GET("/users/search", controllers.SearchUser)
	// router.PUT("/users/:user_id", user.Update)
	// router.PATCH("/users/:user_id", user.Update)
	// router.DELETE("/users/:user_id", user.Delete)
	// router.GET("/internal/users/search", user.Search)
}
