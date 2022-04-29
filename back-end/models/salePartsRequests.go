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
	ID        string         `json:"id"`
	PartID    string         `json:"part_id"`
	PartPrice float64        `json:"part_price"`
	PartQty   int            `json:"part_quantity"`
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

func (saleReq *SalePartRequest) GetTotalRequestQty(id string, db *gorm.DB) (error, int) {

	type RequestQty struct {
		PartQty int
	}

	var partReqQty RequestQty

	err := db.Model(saleReq).Where("id = ?", id).Find(&partReqQty).Error
	if err != nil {
		return err, 0
	}

	return nil, partReqQty.PartQty

}

func (saleReq *SalePartRequest) IncreaseQty(id string, qty int, db *gorm.DB) error {

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

		err := db.Model(&Part{}).Where("id = ?", id).Find(&p).Error
		if err != nil {
			return err
		}

		if qty < saleReq.PartQty {

			increaseQty = saleReq.PartQty - qty

			newQty = p.Quantity + increaseQty

			if newQty < 0 {
				return fmt.Errorf("negative stock error: quantity should be positive")
			}

			err = db.Model(&Part{}).Where("id = ?", id).Update("quantity", newQty).Error
			if err != nil {
				return err
			}

			return nil
		}

		increaseQty = qty - saleReq.PartQty

		newQty = p.Quantity + increaseQty

		if newQty < 0 {
			return fmt.Errorf("negative stock error: quantity should be positive")
		}

		err = db.Model(&Part{}).Where("id = ?", id).Update("quantity", newQty).Error
		if err != nil {
			return err
		}

		return nil

	}

	err = db.Model(&Part{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty = p.Quantity + qty

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

		err := db.Model(&Part{}).Where("id = ?", id).Find(&p).Error
		if err != nil {
			return err
		}

		requestQty = qty - saleReq.PartQty

		newQty = p.Quantity - requestQty

		if newQty < 0 {
			return fmt.Errorf("negative stock error: quantity should be positive")
		}

		err = db.Model(&Part{}).Where("id = ?", id).Update("quantity", newQty).Error
		if err != nil {
			return err
		}

		return nil

	}

	err = db.Model(&Part{}).Where("id = ?", id).Find(&p).Error
	if err != nil {
		return err
	}

	newQty = p.Quantity - qty

	if newQty < 0 {
		return fmt.Errorf("negative stock error: quantity should be positive")
	}

	err = db.Model(&Part{}).Where("id = ?", id).Update("quantity", newQty).Error
	if err != nil {
		return err
	}

	return nil
}
