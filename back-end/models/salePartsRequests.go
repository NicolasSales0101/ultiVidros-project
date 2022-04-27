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

var _ interfaces.Products = &SalePartRequest{}

type SalePartRequest struct {
	ID        string  `json:"id"`
	PartID    string  `json:"glass_id"`
	PartPrice float64 `json:"product_price"`
	PartQty   int     `json:"glass_quantity"`
	//	Product       interfaces.Products `json:"-" gorm:"-:all"`
	SaleID    string         `json:"sale_id" gorm:"size:191"`
	CreatedAt time.Time      `json:"created"`
	UpdatedAt time.Time      `json:"updated"`
	DeletedAt gorm.DeletedAt `json:"deleted" gorm:"index"`
}

func (saleRequest *SalePartRequest) BeforeCreate(scope *gorm.DB) (err error) {

	newId := uuid.New()
	id := strings.Replace(newId.String(), "-", "", 0)

	scope.Statement.SetColumn("ID", id)

	return nil
}

func (saleReq *SalePartRequest) GetTotalProductQty(id string, db *gorm.DB) (error, int) {

	var p queryUtils.ProductQty

	err := db.Model(&Part{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err, 0
	}

	return nil, p.Quantity

}

func (saleReq *SalePartRequest) IncreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var p queryUtils.ProductQty

	err := db.Model(&Part{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty := p.Quantity + qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err = db.Model(&Part{}).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}

func (saleReq *SalePartRequest) DecreaseQty(id string, qty int, db *gorm.DB) error {

	if qty < 0 {
		return fmt.Errorf("negative stock error: quantity number should be positive")
	}

	var p queryUtils.ProductQty

	err := db.Model(&Part{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty := p.Quantity - qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err = db.Model(&Part{}).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}