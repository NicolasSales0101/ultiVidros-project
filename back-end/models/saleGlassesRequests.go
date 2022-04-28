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
	ID            string         `json:"id"`
	GlassID       string         `json:"glass_id"`
	GlassPrice    float64        `json:"glass_price"`
	GlassQty      int            `json:"glass_quantity"`
	RequestHeight float64        `json:"request_height"`
	RequestWidth  float64        `json:"request_width"`
	SaleID        string         `json:"sale_id" gorm:"size:191"`
	CreatedAt     time.Time      `json:"created"`
	UpdatedAt     time.Time      `json:"updated"`
	DeletedAt     gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (saleRequest *SaleGlassRequest) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", 0)

	scope.Statement.SetColumn("ID", id)

	return nil
}

func (saleReq *SaleGlassRequest) findIfExist(id string, db *gorm.DB) (error, bool) {

	var p queryUtils.ProductId

	err := db.Model(saleReq).Where("id = ?", id).Find(p).Error
	if err != nil {
		return err, false
	}

	if p.ID == "" {
		return nil, false
	}

	return nil, true
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

	var (
		p           queryUtils.ProductQty
		increaseQty int
		newQty      int
	)

	err, ver := saleReq.findIfExist(id, db)
	if err != nil {
		return err
	}

	if ver == true {

		err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
		if err != nil {
			return err
		}

		increaseQty = qty - saleReq.GlassQty

		newQty = p.Quantity + increaseQty

		if newQty < 0 {
			return fmt.Errorf("negative stock error: quantity should be positive")
		}

		err = db.Model(&Glass{}).Where("id = ?", id).Update("quantity", newQty).Error
		if err != nil {
			return err
		}

		return nil

	}

	err = db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty = p.Quantity + qty

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

	var (
		p          queryUtils.ProductQty
		requestQty int
		newQty     int
	)

	err, ver := saleReq.findIfExist(id, db)
	if err != nil {
		return err
	}

	if ver == true {

		err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
		if err != nil {
			return err
		}

		requestQty = qty - saleReq.GlassQty

		newQty = p.Quantity - requestQty

		if newQty < 0 {
			return fmt.Errorf("negative stock error: quantity should be positive")
		}

		err = db.Model(&Glass{}).Where("id = ?", id).Update("quantity", newQty).Error
		if err != nil {
			return err
		}

		return nil

	}

	err = db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty = p.Quantity - qty

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

func (g *SaleGlassRequest) IncreaseArea(id string, width, height float64, db *gorm.DB) error {

	var pq queryUtils.ProductCategory

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&pq).Error
	if err != nil {
		return err
	}

	if pq.Category == "tempered" {
		return nil
	}

	var (
		p         queryUtils.ProductArea
		newWidth  float64
		newHeight float64
	)

	err = db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	if width < 0 || height < 0 {
		return fmt.Errorf("width or height is negative")
	}

	err, ver := g.findIfExist(id, db)
	if err != nil {
		return err
	}

	if ver == true {

		increaseWidth := width - g.RequestWidth
		increaseHeight := height - g.RequestHeight

		newWidth = p.WidthAvailable + increaseWidth
		newHeight = p.HeightAvailable + increaseHeight

		err = db.Model(&Glass{}).Where("id = ?", id).Updates(map[string]interface{}{"width_available": newWidth, "height_available": newHeight}).Error
		if err != nil {
			return err
		}

		return nil

	}

	newWidth = p.WidthAvailable + width
	newHeight = p.HeightAvailable + height

	err = db.Model(&Glass{}).Where("id = ?", id).Updates(map[string]interface{}{"width_available": newWidth, "height_available": newHeight}).Error
	if err != nil {
		return err
	}

	return nil

}

func (g *SaleGlassRequest) DecreaseArea(id string, width, height float64, db *gorm.DB) error {

	var pq queryUtils.ProductCategory

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&pq).Error
	if err != nil {
		return err
	}

	if pq.Category == "tempered" {
		return nil
	}

	var (
		p         queryUtils.ProductArea
		newWidth  float64
		newHeight float64
	)

	err = db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	if p.WidthAvailable < width || p.HeightAvailable < height {
		return fmt.Errorf("width or height is large than in stock")
	}

	err, ver := g.findIfExist(id, db)
	if err != nil {
		return err
	}

	if ver == true {

		requestWidth := width - g.RequestWidth
		requestHeight := height - g.RequestHeight

		newWidth = p.WidthAvailable - requestWidth
		newHeight = p.HeightAvailable - requestHeight

		if newWidth < 0 || newHeight < 0 {
			return fmt.Errorf("negative stock error: width or height should be positive or zero")
		}

		err = db.Model(&Glass{}).Where("id = ?", id).Updates(map[string]interface{}{"width_available": newWidth, "height_available": newHeight}).Error
		if err != nil {
			return err
		}

		return nil

	}

	newWidth = p.WidthAvailable - width
	newHeight = p.HeightAvailable - height

	if newWidth < 0 || newHeight < 0 {
		return fmt.Errorf("negative stock error: width or height should be positive or zero")
	}

	err = db.Model(&Glass{}).Where("id = ?", id).Updates(map[string]interface{}{"width_available": newWidth, "height_available": newHeight}).Error
	if err != nil {
		return err
	}

	return nil

}
