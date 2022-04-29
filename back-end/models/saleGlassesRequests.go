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

func (saleReq *SaleGlassRequest) GetTotalProductQty(id string, db *gorm.DB) (error, int) {

	var p queryUtils.ProductQty

	err := db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err, 0
	}

	return nil, p.Quantity

}

func (saleReq *SaleGlassRequest) GetTotalRequestQty(id string, db *gorm.DB) (error, int) {

	type RequestQty struct {
		GlassQty int
	}

	var glassReqQty RequestQty

	err := db.Model(saleReq).Where("id = ?", id).Find(&glassReqQty).Error
	if err != nil {
		return err, 0
	}

	return nil, glassReqQty.GlassQty

}

func (saleReq *SaleGlassRequest) IncreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var (
		p           queryUtils.ProductQty
		pid         queryUtils.ProductId
		increaseQty int
		newQty      int
	)

	err := db.Model(saleReq).Where("id = ?", saleReq.ID).Find(&pid).Error

	if err != nil {
		return err
	}

	if pid.ID != "" {

		err = db.Model(&Glass{}).Where("id = ?", id).Find(&p).Error
		if err != nil {
			return err
		}

		if qty < saleReq.GlassQty {

			increaseQty = saleReq.GlassQty - qty

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
		pid        queryUtils.ProductId
		requestQty int
		newQty     int
	)

	err := db.Model(saleReq).Where("id = ?", saleReq.ID).Find(&pid).Error

	if err != nil {
		return err
	}

	if pid.ID != "" {

		if qty == 0 {
			return nil
		}

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

func (g *SaleGlassRequest) GetWidthAndHeightOfRequest(id string, db *gorm.DB) (error, map[string]float64) {

	type RequestArea struct {
		RequestWidth  float64
		RequestHeight float64
	}

	var glassReqArea RequestArea

	err := db.Model(g).Where("id = ?", id).Find(&glassReqArea).Error
	if err != nil {
		return err, nil
	}

	return nil, map[string]float64{
		"width":  glassReqArea.RequestWidth,
		"height": glassReqArea.RequestHeight,
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
		pid       queryUtils.ProductId
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

	err = db.Model(g).Where("id = ?", g.ID).Find(&pid).Error

	if err != nil {
		return err
	}

	if pid.ID != "" {

		var requestWidth float64
		var requestHeight float64

		if height != 0 && width != 0 {

			if width > g.RequestWidth {
				requestWidth = width - g.RequestWidth
			}

			if height > g.RequestHeight {
				requestHeight = height - g.RequestHeight
			}

			if width < g.RequestWidth {
				requestWidth = g.RequestWidth - width
			}

			if height < g.RequestHeight {
				requestHeight = g.RequestHeight - height
			}

			newWidth = p.WidthAvailable + requestWidth
			newHeight = p.HeightAvailable + requestHeight

			if newWidth < 0 || newHeight < 0 {
				return fmt.Errorf("negative stock error: width or height should be positive or zero")
			}

			err = db.Model(&Glass{}).Where("id = ?", id).Updates(map[string]interface{}{"width_available": newWidth, "height_available": newHeight}).Error
			if err != nil {
				return err
			}

			return nil
		}

		if height == 0 && width != 0 {

			if width > g.RequestWidth {
				requestWidth = width - g.RequestWidth
			}

			if width < g.RequestWidth {
				requestWidth = g.RequestWidth - width
			}

			newWidth = p.WidthAvailable + requestWidth

			if newWidth < 0 {
				return fmt.Errorf("negative stock error: width or height should be positive or zero")
			}

			err = db.Model(&Glass{}).Where("id = ?", id).Update("width_available", newWidth).Error
			if err != nil {
				return err
			}

			return nil
		}

		if height != 0 && width == 0 {

			if height > g.RequestHeight {
				requestHeight = height - g.RequestHeight
			}

			if height < g.RequestHeight {
				requestHeight = g.RequestHeight - height
			}

			newHeight = p.HeightAvailable + requestHeight

			if newHeight < 0 {
				return fmt.Errorf("negative stock error: width or height should be positive or zero")
			}

			err = db.Model(&Glass{}).Where("id = ?", id).Update("height_available", newHeight).Error
			if err != nil {
				return err
			}

			return nil
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
		pid       queryUtils.ProductId
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

	err = db.Model(g).Where("id = ?", g.ID).Find(&pid).Error
	if err != nil {
		return err
	}

	if pid.ID != "" {

		if height != 0 && width != 0 {

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
		}

		if height == 0 && width != 0 {

			requestWidth := width - g.RequestWidth

			newWidth = p.WidthAvailable - requestWidth

			if newWidth < 0 {
				return fmt.Errorf("negative stock error: width or height should be positive or zero")
			}

			err = db.Model(&Glass{}).Where("id = ?", id).Update("width_available", newWidth).Error
			if err != nil {
				return err
			}

		}

		if height != 0 && width == 0 {

			requestHeight := height - g.RequestHeight

			newHeight = p.HeightAvailable - requestHeight

			if newHeight < 0 {
				return fmt.Errorf("negative stock error: width or height should be positive or zero")
			}

			err = db.Model(&Glass{}).Where("id = ?", id).Update("height_available", newHeight).Error
			if err != nil {
				return err
			}
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
