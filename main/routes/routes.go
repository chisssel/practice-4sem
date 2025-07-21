package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"practice-4sem/controllers"
	"practice-4sem/repositories"
)

//func SetupRoutes() *gin.Engine {
//	r := gin.Default()
//
//	db := config.DB
//
//	counterController := controllers.NewCounterController(db)
//	itemRepo := repositories.NewItemRepository(db)
//	itemController := controllers.NewItemController(itemRepo)
//
//	api := r.Group("/api")
//	{
//		api.GET("/counters", counterController.GetCounters)
//		api.GET("/items", itemController.GetItems)
//		// возможно, надо будет дополнить еще маршрутами
//	}
//
//	return r
//}

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	itemRepo := repositories.NewItemRepository(db)
	itemController := controllers.NewItemController(itemRepo)

	api := r.Group("/api/v1")
	{
		api.GET("/items", itemController.GetItems)
	}

	return r
}
