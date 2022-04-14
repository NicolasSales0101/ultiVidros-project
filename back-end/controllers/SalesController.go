package controllers

import (
	"log"

	"github.com/NicolasSales0101/ultiVidros-project/back-end/database"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/database/dbUtils"
	"github.com/NicolasSales0101/ultiVidros-project/back-end/models"
	"github.com/gofiber/fiber/v2"
)

type Result struct {
	SaleId       string  `json:"id"`
	ClientId     string  `json:"client_id"`
	ChatId       string  `json:"chat_id"`
	TbProductsId string  `json:"products_id"`
	ProductId    string  `json:"product_id"`
	ProductPrice float64 `json:"product_price"`
	ProductQty   int     `json:"product_quantity"`
	CodeSupport  string  `json:"code_support"`
}

type Test struct {
	Id string `json:"tb_products_id"`
}

func ShowSales(fctx *fiber.Ctx) error {
	db := database.GetDatabase()

	var result []Result
	var test []Test

	err := db.Raw("SELECT sales.id, sales.client_id, sales.chat_id, products.id, products.product_id, products.product_price, products.product_qty, products.sale_id FROM sales LEFT JOIN products ON products.sale_id = sales.id WHERE sales.deleted_at IS NULL").Scan(&result).Error
	db.Raw("SELECT products.id FROM products LEFT JOIN sales ON products.sale_id = sales.id WHERE products.sale_id = sales.id").Scan(&test)
	log.Println(test)
	if err != nil {
		log.Println("Error in method Get ShowSales:", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot list sales: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(result)
}

func ShowSale(fctx *fiber.Ctx) error {
	id := fctx.Params("id")

	db := database.GetDatabase()

	var sale models.Sale

	err := db.First(&sale, "id = ?", id).Error
	if err != nil {
		log.Println("Error in method Get ShowSale (specific id in url params):", err)
		return fctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot find sale: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(sale)
}

func CreateSale(fctx *fiber.Ctx) error {
	var sale models.Sale

	if err := fctx.BodyParser(&sale); err != nil {
		log.Println("Error in method Post CreateSale:", err)
		return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
			"error": "cannot create sale: " + err.Error(),
		})
	}

	for _, v := range sale.Products {

		err, productQty := dbUtils.GetTotalProductQty(v.ProductID)
		if err != nil {
			log.Println("Error in method Post CreateSale:", err)
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error: " + err.Error(),
			})
		}

		if v.ProductQty > productQty {
			log.Println(v.ProductQty, productQty)
			return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
				"error": "Product quantity in request is big than quantity in stock",
				"data":  sale,
			})
		}
	}

	db := database.GetDatabase()
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
	var sale models.Sale

	if err := fctx.BodyParser(&sale); err != nil {
		log.Println("Error in method Put UpdateSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot update sale: " + err.Error(),
		})
	}

	for _, v := range sale.Products {

		err, productQty := dbUtils.GetTotalProductQty(v.ProductID)
		if err != nil {
			log.Println("Error in method Put UpdateSale:", err)
			return fctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error: " + err.Error(),
			})
		}

		if v.ProductQty > productQty {
			return fctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
				"error": "Product quantity in request is big than quantity in stock",
				"data":  sale,
			})
		}
	}

	db := database.GetDatabase()

	err := db.Omit("CreatedAt").Save(&sale).Error
	if err != nil {
		log.Println("Error in method Put UpdateSale:", err)
		return fctx.Status(fiber.StatusNotModified).JSON(fiber.Map{
			"error": "cannot update sale: " + err.Error(),
		})
	}

	return fctx.Status(fiber.StatusOK).JSON(sale)
}

func DeleteSale(fctx *fiber.Ctx) error {
	id := fctx.Params("id")

	db := database.GetDatabase()

	err := db.Delete(&models.Sale{}, "id = ?", id).Error
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
