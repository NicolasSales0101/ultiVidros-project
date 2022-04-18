package models

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaleRequest struct {
	ID            string         `json:"id"`
	ProductID     string         `json:"product_id"`
	ProductPrice  float64        `json:"product_price"`
	ProductQty    int            `json:"product_quantity"`
	RequestHeight float64        `json:"request_height"`
	RequestWidth  float64        `json:"request_width"`
	SaleID        string         `json:"sale_id" gorm:"size:191"`
	CreatedAt     time.Time      `json:"created"`
	UpdatedAt     time.Time      `json:"updated"`
	DeletedAt     gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (saleRequest *SaleRequest) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", 0)

	scope.Statement.SetColumn("ID", id)

	return nil
}
