package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database/queryUtils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Glass struct {
	ID              string         `json:"id" gorm:"primary_key"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	Price           float64        `json:"price"`
	ImageURL        string         `json:"img_url"`
	HeightAvailable float64        `json:"height_available"`
	WidthAvailable  float64        `json:"width_available"`
	Category        string         `json:"category"`
	Type            string         `json:"type"`
	GlassSheets     int            `json:"glass_sheets"`
	Quantity        int            `json:"quantity"`
	CreatedAt       time.Time      `json:"created"`
	UpdatedAt       time.Time      `json:"updated"`
	DeletedAt       gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (glass *Glass) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", -1)

	scope.Statement.SetColumn("ID", id)

	return nil
}

func (g *Glass) GetTotalProductQty(id string, db *gorm.DB) (error, int) {

	var p queryUtils.ProductQty

	err := db.Model(g).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err, 0
	}

	return nil, p.Quantity
}

func (g *Glass) IncreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var newQty int

	newQty = g.Quantity + qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err := db.Model(g).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil

}

func (g *Glass) DecreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var newQty int

	newQty = g.Quantity - qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err := db.Model(g).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *Glass) GetWidthAndHeightAvailableOfProduct(id string, db *gorm.DB) (error, map[string]float64) {

	var p queryUtils.ProductArea

	err := db.Model(g).Where("id = ?").Find(&p).Error
	if err != nil {
		return err, nil
	}

	return nil, map[string]float64{
		"width":  p.WidthAvailable,
		"height": p.HeightAvailable,
	}
}

func (g *Glass) IncreaseArea(id string, width, height float64, db *gorm.DB) error {

	if g.Category == "tempered" {
		return nil
	}

	if width < 0 || height < 0 {
		return fmt.Errorf("width or height is negative")
	}

	var (
		newWidth  float64
		newHeight float64
	)

	newWidth = g.WidthAvailable + width
	newHeight = g.HeightAvailable + height

	err := db.Model(g).Where("id = ?", id).Updates(map[string]interface{}{"width_available": newWidth, "height_available": newHeight}).Error
	if err != nil {
		return err
	}

	return nil
}

func (g *Glass) DecreaseArea(id string, width, height float64, db *gorm.DB) error {

	if g.Category == "tempered" {
		return nil
	}

	var (
		newWidth  float64
		newHeight float64
	)

	if g.WidthAvailable < width || g.HeightAvailable < height {
		return fmt.Errorf("width or height is large than in stock")
	}

	newWidth = g.WidthAvailable - width
	newHeight = g.HeightAvailable - height

	if newWidth < 0 || newHeight < 0 {
		return fmt.Errorf("negative stock error: width or height should be positive or zero")
	}

	err := db.Model(g).Where("id = ?", id).Updates(map[string]interface{}{"width_available": newWidth, "height_available": newHeight}).Error
	if err != nil {
		return err
	}

	return nil
}
