package routes

import (
	"lead/controllers"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
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

	//Caching
	memoryStore := persist.NewMemoryStore(3 * time.Minute)
	ttlCacheExpiry := 180 * time.Second

	//Routes with Caching
	// Print all Moves
	server.GET("/", cache.CacheByRequestURI(memoryStore, ttlCacheExpiry), controllers.Index)
	// Get Move
	server.GET("/:code", cache.CacheByRequestURI(memoryStore, ttlCacheExpiry), controllers.Code)

	return server
}
