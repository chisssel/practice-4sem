package models

import "gorm.io/gorm"

type Organization struct {
	gorm.Model
	OrgID   string `gorm:"primaryKey"`
	OrgName string
}
