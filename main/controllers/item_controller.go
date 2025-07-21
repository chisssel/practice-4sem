package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "practice-4sem/models"
	"practice-4sem/repositories"
)

type ItemController struct {
	repo *repositories.ItemRepository
}

func NewItemController(repo *repositories.ItemRepository) *ItemController {
	return &ItemController{repo: repo}
}

func (c *ItemController) GetItems(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("page_size", "20"))
	if err != nil || pageSize < 1 || pageSize > 200 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Page size must be between 1 and 200"})
		return
	}

	items, err := c.repo.GetPaginatedItems(page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	total, err := c.repo.GetTotalCount()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total count"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       items,
		"page":       page,
		"page_size":  pageSize,
		"total":      total,
		"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}
