package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database/queryUtils"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/interfaces"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ interfaces.Products = &SaleGlassRequest{}

type SaleGlassRequest struct {
	ID            string  `json:"id"`
	GlassID       string  `json:"glass_id"`
	GlassPrice    float64 `json:"glass_price"`
	GlassQty      int     `json:"glass_quantity"`
	RequestHeight float64 `json:"request_height"`
	RequestWidth  float64 `json:"request_width"`
	//	Product       interfaces.Products `json:"-" gorm:"-:all"`
	SaleID    string         `json:"sale_id" gorm:"size:191"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (saleRequest *SaleGlassRequest) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", 0)

	scope.Statement.SetColumn("ID", id)

	return nil
}

func (saleReq *SaleGlassRequest) GetTotalProductQty(id string, db *gorm.DB) (error, int) {

	var p queryUtils.ProductQty

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err, 0
	}

	return nil, p.Quantity

}

func (saleReq *SaleGlassRequest) IncreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var p queryUtils.ProductQty

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty := p.Quantity + qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err = db.Model(&Glass{}).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}

func (saleReq *SaleGlassRequest) DecreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var p queryUtils.ProductQty

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty := p.Quantity - qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err = db.Model(&Glass{}).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *SaleGlassRequest) GetWidthAndHeightAvailableOfProduct(id string, db *gorm.DB) (error, map[string]float64) {

	var p queryUtils.ProductArea

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err, nil
	}

	return nil, map[string]float64{
		"width":  p.WidthAvailable,
		"height": p.HeightAvailable,
	}
}

func (g *SaleGlassRequest) ReduceArea(id string, width, height float64, db *gorm.DB) error {

	var p queryUtils.ProductArea

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return nil
	}

	if p.WidthAvailable < width || p.HeightAvailable < height {
		return fmt.Errorf("width or height is large than in stock")
	}

	err = db.Model(&Glass{}).Where("id = ?", id).Updates(map[string]interface{}{"width_available": width, "height_available": height}).Error
	if err != nil {
		return err
	}

	return nil

}
