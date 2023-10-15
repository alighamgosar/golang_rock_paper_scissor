package models

import (
	"time"

	"gorm.io/gorm"
)

type Tournament struct {
	gorm.Model
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	// Players in the tournament
	Players []Player `gorm:"foreignKey:id"`

	// Matches in the tournament
	Game []Game `gorm:"foreignKey:id"`
}
