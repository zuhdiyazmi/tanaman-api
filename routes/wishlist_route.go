package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zuhdiyazmi/alterra-mini-project/controllers/wishlist_controller"
)

// SetupRoutes menginisialisasi rute dan handler
func SetupWishlistRoutes(r *gin.Engine) {
	r.GET("/api/wishlists", wishlist_controller.Index)
	r.GET("/api/wishlist/:id", wishlist_controller.Show)
	r.POST("/api/wishlist", wishlist_controller.Create)
	r.PUT("/api/wishlist/:id", wishlist_controller.Update)
	r.DELETE("/api/wishlist/:id", wishlist_controller.Delete)
	r.GET("/api/wishlists/search", wishlist_controller.Search)
}
