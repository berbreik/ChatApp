package main

import (
	"chatapp/v/internal/auth"
	"chatapp/v/internal/db"
	"chatapp/v/internal/projects"
	"chatapp/v/internal/proposals"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database := db.ConnectPostgres()

	auth.NewHandler(database).RegisterRoutes(r)
	projects.NewHandler(database).RegisterRoutes(r)
	proposals.NewHandler(database).RegisterRoutes(r)

	r.Run(":8080")
}
