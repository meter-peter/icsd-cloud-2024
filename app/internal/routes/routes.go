package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/meter-peter/icsd-cloud-2024/internal/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", handlers.GetUsers(db))
	r.GET("/users/:id", handlers.GetUser(db))
	r.POST("/users", handlers.CreateUser(db))
	r.PUT("/users/:id", handlers.UpdateUser(db))
	r.DELETE("/users/:id", handlers.DeleteUser(db))
	r.GET("/users/search", handlers.SearchUsers(db))
}
