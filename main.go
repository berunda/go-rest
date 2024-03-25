package main

import (
	"github.com/berunda/go-rest/db"
	"github.com/berunda/go-rest/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}
