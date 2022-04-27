package migrations

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Glass{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Sale{})
	db.AutoMigrate(models.SaleGlassRequest{})
	db.AutoMigrate(models.SalePartRequest{})
	db.AutoMigrate(models.Part{})
}
