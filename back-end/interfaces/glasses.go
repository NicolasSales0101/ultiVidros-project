package interfaces

import "gorm.io/gorm"

type Glass interface {
	GetWidthAndHeightAvailableOfProduct(id string, db *gorm.DB) (error, map[string]float64)
	DecreaseArea(id string, width, height float64, db *gorm.DB) error
	IncreaseArea(id string, width, height float64, db *gorm.DB) error
}
