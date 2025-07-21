package repositories

import (
	"gorm.io/gorm"
	"practice-4sem/models"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) GetPaginatedItems(page, pageSize int) ([]models.Item, error) {
	var items []models.Item

	if pageSize > 200 {
		pageSize = 200
	}

	offset := (page - 1) * pageSize

	result := r.db.
		Order("id ASC").
		Offset(offset).
		Limit(pageSize).
		Find(&items)

	if result.Error != nil {
		return nil, result.Error
	}

	return items, nil
}

func (r *ItemRepository) GetTotalCount() (int64, error) {
	var count int64
	result := r.db.Model(&models.Item{}).Count(&count)
	return count, result.Error
}
