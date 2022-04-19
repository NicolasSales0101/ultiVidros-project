package interfaces

import (
	"gorm.io/gorm"
)

type Products interface {
	GetTotalProductQty(id string, db *gorm.DB) (error, int)
	IncreaseQty(id string, qty int, db *gorm.DB) error
	DecreaseQty(id string, qty int, db *gorm.DB) error
}
