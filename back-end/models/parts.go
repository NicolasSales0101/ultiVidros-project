package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database/queryUtils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Part struct {
	ID          string         `json:"id" gorm:"primary_key"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	ImageURL    string         `json:"img_url"`
	Category    string         `json:"category"`
	Quantity    int            `json:"quantity"`
	CreatedAt   time.Time      `json:"created"`
	UpdatedAt   time.Time      `json:"updated"`
	DeletedAt   gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (part *Part) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", -1)

	scope.Statement.SetColumn("ID", id)

	return nil
}

func (part *Part) GetTotalProductQty(id string, db *gorm.DB) (error, int) {

	var p queryUtils.ProductQty

	err := db.Model(part).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err, 0
	}

	return nil, p.Quantity
}

func (part *Part) IncreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var p queryUtils.ProductQty

	err := db.Model(part).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty := p.Quantity + qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err = db.Model(part).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}

func (part *Part) DecreaseQty(id string, qty int, db *gorm.DB) error {

	var p queryUtils.ProductQty

	err := db.Model(part).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty := p.Quantity - qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity in request is large than quantity in stock")
	}

	err = db.Model(part).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}
