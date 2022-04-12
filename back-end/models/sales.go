package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Continue here

type Product struct {
	ProductID    string  `json:"product_id"`
	ProductPrice float64 `json:"product_price"`
}

type Sale struct {
	ID        string         `json:"id" gorm:"primary_key"`
	ClientID  string         `json:"client_id"`
	ChatID    string         `json:"chat_id"`
	Products  []Product      `json:"products"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (sale *Sale) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", -1)

	scope.Statement.SetColumn("ID", id)

	return nil
}
