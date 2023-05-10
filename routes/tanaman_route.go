package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zuhdiyazmi/alterra-mini-project/controllers/tanaman_controller"
)

// SetupRoutes menginisialisasi rute dan handler
func SetupTanamanRoutes(r *gin.Engine) {
	r.GET("/api/tanamans", tanaman_controller.Index)
	r.GET("/api/tanaman/:id", tanaman_controller.Show)
	r.POST("/api/tanaman", tanaman_controller.Create)
	r.PUT("/api/tanaman/:id", tanaman_controller.Update)
	r.DELETE("/api/tanaman/:id", tanaman_controller.Delete)
	r.GET("/api/tanamans/search", tanaman_controller.Search)
}
