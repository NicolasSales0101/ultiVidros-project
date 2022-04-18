package dbUtils

import (
	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
)

func GetTotalProductQty(id string) (error, int) {

	var glass models.Glass

	db := database.GetDatabase()

	if err := db.First(&glass, "id = ?", id).Error; err != nil {
		return err, 0
	}

	return nil, glass.Quantity

}

func GetWidthAndHeightAvailableOfProduct(id string) (error, map[string]float64) {
	var glass models.Glass

	db := database.GetDatabase()

	if err := db.First(&glass, "id = ?", id).Error; err != nil {
		return err, nil
	}

	area := map[string]float64{
		"width":  glass.WidthAvailable,
		"height": glass.HeightAvailable,
	}

	return nil, area
}
