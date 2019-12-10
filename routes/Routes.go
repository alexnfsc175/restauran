package routes

import (
	"os"

	Controllers "github.com/alexnfsc175/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter : describe what this function does
func SetupRouter() *gin.Engine {

	if isDev := os.Getenv("APP_ENV"); isDev == "dev" {
		gin.SetMode(gin.DebugMode)
	}

	if isProd := os.Getenv("APP_ENV"); isProd == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("todo", Controllers.GetTodos)
		v1.POST("todo", Controllers.CreateATodo)
		v1.GET("todo/:id", Controllers.GetATodo)
		v1.PUT("todo/:id", Controllers.UpdateATodo)
		v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}

	return r
}
