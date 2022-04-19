package dbUtils

import (
	"fmt"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/interfaces"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"gorm.io/gorm"
)

func GetTotalProductQty(product interfaces.Products, db *gorm.DB) (error, int) {

	var err error

	switch p := product.(type) {
	case *models.Glass:
		err, qty := p.GetTotalProductQty(p.ID, db)
		return err, qty
	case *models.Part:
		err, qty := p.GetTotalProductQty(p.ID, db)
		return err, qty
	default:
		err = fmt.Errorf("unknown model")
	}

	return err, 0

}

func IncreaseQty(product interfaces.Products, qty int, db *gorm.DB) error {

	var err error

	switch p := product.(type) {
	case *models.Glass:
		err = p.IncreaseQty(p.ID, qty, db)
		return err
	case *models.Part:
		err = p.IncreaseQty(p.ID, qty, db)
		return err
	default:
		err = fmt.Errorf("unknown model")
	}

	return err

}

func DecreaseQty(product interfaces.Products, qty int, db *gorm.DB) error {

	var err error

	switch p := product.(type) {
	case *models.Glass:
		err = p.DecreaseQty(p.ID, qty, db)
		return err
	case *models.Part:
		err = p.DecreaseQty(p.ID, qty, db)
		return err
	default:
		err = fmt.Errorf("unknown model")
	}

	return err

}
