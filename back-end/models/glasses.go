package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Glass struct {
	ID              string         `json:"id" gorm:"primary_key"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Price           float64        `json:"price"`
	ImageURL        string         `json:"img_url"`
	HeightAvailable float64        `json:"height_available"`
	WidthAvailable  float64        `json:"width_available"`
	Category        string         `json:"category"`
	Type            string         `json:"type"`
	GlassSheets     int            `json:"glass_sheets"`
	Quantity        int            `json:"quantity"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (glass *Glass) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", -1)

	scope.Statement.SetColumn("ID", id)

	return nil
}
