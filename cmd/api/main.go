package main

import (
	"chatapp/v/internal/auth"
	"chatapp/v/internal/chat"
	"chatapp/v/internal/db"
	"chatapp/v/internal/middleware"
	"chatapp/v/internal/projects"
	"chatapp/v/internal/proposals"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database := db.ConnectPostgres()
	streamClient := chat.NewStreamClient()

	// Public routes
	auth.NewHandler(database).RegisterRoutes(r)

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	projects.NewHandler(database).RegisterRoutes(protected)
	proposals.NewHandler(database).RegisterRoutes(protected)
	chat.NewHandler(streamClient).RegisterRoutes(protected)

	r.Run(":8080")
}
