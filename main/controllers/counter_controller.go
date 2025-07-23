package controllers

import (
	"net/http"
	"strconv"

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	const maxPageSize = 200
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize

	var counters []models.Counters
	var total int64

	if err := cc.DB.Model(&models.Counters{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := cc.DB.Offset(offset).Limit(pageSize).Find(&counters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": counters,
		"meta": gin.H{
			"page":       page,
			"page_size":  pageSize,
			"total":      total,
			"total_page": (int(total) + pageSize - 1) / pageSize,
		},
	})
}
