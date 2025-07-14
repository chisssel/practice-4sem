package models

import "gorm.io/gorm"

type TypeCounter struct {
	gorm.Model
	TypeNumber  string `gorm:"primaryKey"`
	TypeName    string
	CounterName string
}
