package models

import "gorm.io/gorm"

type Counters struct {
	gorm.Model
	CounterCertificateID string `gorm:"primaryKey"`
	TypeNumber           string
	CounterMode          string
	FactoryNumber        string
}
