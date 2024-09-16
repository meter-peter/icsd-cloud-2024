package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/users", GetUsers(db))
	r.GET("/users/:id", GetUser(db))
	r.POST("/users", CreateUser(db))
	r.PUT("/users/:id", UpdateUser(db))
	r.DELETE("/users/:id", DeleteUser(db))
}
