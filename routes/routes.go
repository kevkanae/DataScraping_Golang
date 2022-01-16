package routes

import (
	"lead/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	server := gin.Default()

	//CORS
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Authorization", "Origin", "X-Requested-With", "Content-Type", "Accept"}
	config.AllowAllOrigins = true
	server.Use(cors.New(config))

	//GIN Extras
	server.Use(gin.Recovery())

	//Load HTML
	server.LoadHTMLGlob("templates/*")

	//Routes
	// Print all Moves
	server.GET("/", controllers.Index)
	// Get Move
	server.GET("/:code", controllers.Code)

	return server
}
