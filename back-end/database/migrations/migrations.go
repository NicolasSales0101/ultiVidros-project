package migrations

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Glass{})
}