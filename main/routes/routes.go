package routes

import (
	"practice-4sem/config"
	"practice-4sem/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	db := config.DB

	counterController := controllers.NewCounterController(db)

	api := r.Group("/api")
	{
		api.GET("/counters", counterController.GetCounters)
		// Добавьте другие маршруты
	}

	return r
}
