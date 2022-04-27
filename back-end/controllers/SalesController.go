package controllers

import (
	"fmt"
	"log"
	"reflect"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/database/dbUtils"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Result struct {
	SaleData          models.Sale               `json:"sale_data"`
	SaleGlassRequests []models.SaleGlassRequest `json:"sale_glass_requests_data"`
	SalePartRequests  []models.SalePartRequest  `json:"sale_part_requests_data"`
}

// TODO: test everything

func ShowSales(fctx *fiber.Ctx) error {

	db := database.GetDatabase()

	var sales []models.Sale
	var salesGlassesRequests []models.SaleGlassRequest
	var salesPartsRequests []models.SalePartRequest

	err := db.Find(&sales).Error
	if err != nil {
		log.Println("Error in method Get ShowSales:", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot list sales: " + err.Error(),
		})
	}

	err = db.Find(&salesGlassesRequests).Error
	if err != nil {
		log.Println("Error in method Get ShowSales:", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot list sales (glasses requests): " + err.Error(),
		})
	}

	err = db.Find(&salesPartsRequests).Error
	if err != nil {
		log.Println("Error in method Get ShowSales:", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot list sales (parts requests): " + err.Error(),
		})
	}

	var result = make([]Result, len(sales))

	for i, v := range sales {
		for _, value := range salesGlassesRequests {
			if v.ID == value.SaleID {
				result[i].SaleData = v
				result[i].SaleGlassRequests = append(result[i].SaleGlassRequests, value)
			}
		}

		for _, value := range salesPartsRequests {
			if v.ID == value.SaleID {
				// result[i].SaleData = v
				result[i].SalePartRequests = append(result[i].SalePartRequests, value)
			}
		}

	}

	return fctx.Status(fiber.StatusOK).JSON(result)
}

func ShowSale(fctx *fiber.Ctx) error {

	id := fctx.Params("id")

	db := database.GetDatabase()

	var sale models.Sale
	var salesGlassesRequests []models.SaleGlassRequest
	var salesPartsRequests []models.SalePartRequest

	err := db.First(&sale, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method Get ShowSale (specific id in url params):", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find sale: " + err.Error(),
		})
	}

	err = db.Find(&salesGlassesRequests, "sale_id = ?", sale.ID).Error
	if err != nil {
		log.Println("Error in method Get ShowSale (specific id in url params):", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find sale (glasses requests): " + err.Error(),
		})
	}

	err = db.Find(&salesPartsRequests, "sale_id = ?", sale.ID).Error
	if err != nil {
		log.Println("Error in method Get ShowSale (specific id in url params):", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find sale (parts requests): " + err.Error(),
		})
	}

	var result Result

	for _, value := range salesGlassesRequests {
		result.SaleData = sale
		result.SaleGlassRequests = append(result.SaleGlassRequests, value)
	}

	for _, value := range salesPartsRequests {
		// result.SaleData = sale
		result.SalePartRequests = append(result.SalePartRequests, value)
	}

	return fctx.Status(fiber.StatusOK).JSON(result)
}

// TODO: test CreateSale and UpdateSale

func CreateSale(fctx *fiber.Ctx) error {

	var sale models.Sale

	if err := fctx.BodyParser(&sale); err != nil {
		log.Println("Error in method Post CreateSale:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create sale: " + err.Error(),
		})
	}

	db := database.GetDatabase()

	for _, v := range sale.GlassRequests {

		err, productQty := v.GetTotalProductQty(v.GlassID, db)
		if err != nil {
			log.Println("Error in method Post CreateSale:", err)
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error: " + err.Error(),
			})
		}

		if v.GlassQty > productQty {
			return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
				"error": "Glass quantity in request is big than quantity in stock",
				"data":  sale,
			})
		}

		err, areaAvailable := v.GetWidthAndHeightAvailableOfProduct(v.ID, db)
		if err != nil {
			log.Println("Error in method Put UpdateSale:", err)
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error: " + err.Error(),
			})
		}

		if v.RequestWidth > areaAvailable["width"] || v.RequestHeight > areaAvailable["height"] {
			return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
				"error": "Product width or height is big than width and height avalible in stock",
				"data":  sale,
			})
		}

		err = v.DecreaseQty(v.GlassID, v.GlassQty, db)
		if err != nil {
			return err
		}

		err = v.ReduceArea(v.GlassID, v.RequestWidth, v.RequestHeight, db)
		if err != nil {
			return err
		}
	}

	for _, v := range sale.PartRequests {

		err, productQty := v.GetTotalProductQty(v.PartID, db)
		if err != nil {
			log.Println("Error in method Post CreateSale:", err)
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error: " + err.Error(),
			})
		}

		if v.PartQty > productQty {
			return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
				"error": "Part quantity in request is big than quantity in stock",
				"data":  sale,
			})
		}

		err = v.DecreaseQty(v.PartID, v.PartQty, db)
		if err != nil {
			return err
		}

	}

	err := db.Create(&sale).Error
	if err != nil {
		log.Println("Error in method Post CreateSale:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create sale: " + err.Error(),
			"data":  sale,
		})
	}

	return fctx.Status(fiber.StatusCreated).JSON(sale)
}

func UpdateSale(fctx *fiber.Ctx) error {

	var (
		result Result
		db     *gorm.DB = database.GetDatabase()
	)

	if err := fctx.BodyParser(&result); err != nil {
		log.Println("Error in method Put UpdateSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot bind sale: " + err.Error(),
			"data":  result,
		})
	}

	for _, v := range result.SaleRequests {

		if v.ID == "" || v.SaleID == "" {
			log.Println("Error in method Put UpdateSale: empty JSON in body")
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error: empty JSON in body",
				"data":  result,
			})
		}

		fmt.Println("--->", reflect.TypeOf(v))

		switch t := v.Product.(type) {

		case *models.Glass:

			err, productQty := dbUtils.GetTotalProductQty(v.Product, db)
			if err != nil {
				log.Println("Error in method Put UpdateSale:", err)
				return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "internal server error: " + err.Error(),
					"data":  result,
				})
			}

			if v.ProductQty > productQty {
				return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
					"error": "Product quantity in request is big than quantity in stock",
					"data":  result,
				})
			}

			err, areaAvailable := t.GetWidthAndHeightAvailableOfProduct(t.ID, db)
			if err != nil {
				log.Println("Error in method Put UpdateSale:", err)
				return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "internal server error: " + err.Error(),
					"data":  result,
				})
			}

			if v.RequestWidth > areaAvailable["width"] || v.RequestHeight > areaAvailable["height"] {
				return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
					"error": "Product width or height is big than width and height avalible in stock",
					"data":  result,
				})
			}

			err = dbUtils.DecreaseQty(t, v.ProductQty, db)
			if err != nil {
				return err
			}

			err = t.ReduceArea(t.ID, v.RequestWidth, v.RequestHeight, db)
			if err != nil {
				return err
			}

		case *models.Part:

			err, productQty := dbUtils.GetTotalProductQty(v.Product, db)
			if err != nil {
				log.Println("Error in method Put UpdateSale:", err)
				return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "internal server error: " + err.Error(),
					"data":  result,
				})
			}

			if v.ProductQty > productQty {
				return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
					"error": "Product quantity in request is big than quantity in stock",
					"data":  result,
				})
			}

			err = dbUtils.DecreaseQty(t, v.ProductQty, db)
			if err != nil {
				return err
			}

		}
	}

	err := db.Omit("CreatedAt").Save(&result.SaleData).Error
	if err != nil {
		log.Println("Error in method Put UpdateSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot update sale: " + err.Error(),
			"data":  result,
		})
	}

	for _, v := range result.SaleRequests {

		log.Println(v)

		err := db.Omit("CreatedAt").Save(&v).Error
		if err != nil {
			log.Println("Error in method Put UpdateSale:", err)
			return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
				"error": "cannot update sale: " + err.Error(),
				"data":  result,
			})
		}

	}

	return fctx.Status(fiber.StatusOK).JSON(result)
}

func DeleteSale(fctx *fiber.Ctx) error {

	id := fctx.Params("id")

	db := database.GetDatabase()

	var saleRequests []models.SaleRequest
	err := db.Find(&saleRequests, "sale_id = ?", id).Error
	if err != nil {
		log.Println("Error in method Delete DeleteSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot delete sale: " + err.Error(),
		})
	}

	for _, v := range saleRequests {
		err := db.Delete(&v, "sale_id = ?", id).Error
		if err != nil {
			log.Println("Error in method Delete DeleteSale:", err)
			return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
				"error": "cannot delete sale: " + err.Error(),
			})
		}
	}

	err = db.Delete(&models.Sale{}, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method Delete DeleteSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot delete sale: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "Sale removed successfully",
	})
}
