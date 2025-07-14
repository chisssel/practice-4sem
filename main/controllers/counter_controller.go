package controllers

import (
	"net/http"

	"practice-4sem/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CounterController struct {
	DB *gorm.DB
}

func NewCounterController(db *gorm.DB) *CounterController {
	return &CounterController{DB: db}
}

func (cc *CounterController) GetCounters(c *gin.Context) {
	var counters []models.Counters
	if result := cc.DB.Find(&counters); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, counters)
}
