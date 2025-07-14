package models

import "gorm.io/gorm"

type Checking struct {
	gorm.Model
	CheckId              string `gorm:"primaryKey"`
	OrgId                string
	CounterCertificateID string
	CheckDate            string
	ValidDate            string
	CounterCondition     string
	CheckRecordCreate    string
	CheckRecordUpdate    string
	Comment              string
	CheckNextID          string
	CheckType            string
}
