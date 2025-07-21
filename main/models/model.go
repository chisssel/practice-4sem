package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	CheckId              string `gorm:"primaryKey" json:"checkId"`
	OrgId                string `gorm:"column:org" json:"orgId"`
	CounterCertificateID string `gorm:"column:counter" json:"certificateId"`
	// Добавить другие поля. Эти пока что для проверки
}
