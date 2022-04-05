package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Glass struct {
	ID              uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;"`
	Name            string         `json:"name"`
	Description     string         `json:description`
	Price           float64        `json:"price"`
	ImageURL        string         `json:"img_url"`
	HeightAvailable float64        `json:"height_available"`
	WidthAvailable  float64        `json:"width_available"`
	Category        string         `json:"category"`
	Type            string         `json:"type"`
	GlassSheets     int            `json:"glass_sheets"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (glass *Glass) BeforeCreate(scope *gorm.DB) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	scope.Statement.SetColumn("ID", uuid.String())
	return nil
}
