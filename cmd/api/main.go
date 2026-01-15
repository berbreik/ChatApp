package main

import (
	"chatapp/v/internal/auth"
	"chatapp/v/internal/chat"
	"chatapp/v/internal/db"
	"chatapp/v/internal/projects"
	"chatapp/v/internal/proposals"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.ConnectPostgres()
	streamClient := chat.InitStreamClient()
	r := gin.Default()

	auth.RegisterRoutes(r, db)
	projects.RegisterRoutes(r, db)
	proposals.RegisterRoutes(r, db)
	chat.RegisterRoutes(r, streamClient)
	r.Run(":8080")
}
