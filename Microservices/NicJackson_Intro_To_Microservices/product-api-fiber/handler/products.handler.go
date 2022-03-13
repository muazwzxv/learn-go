package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"muazwzxv/product-api-fiber/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	log *log.Logger
}

func NewProductHandler(logger *log.Logger) *ProductHandler {
	return &ProductHandler{logger}
}

func (p *ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	p.log.Println("Handle PUT Request")

	product := &entity.Product{}
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	i, _ := strconv.Atoi(ctx.Params("id"))
	entity.UpdateProduct(i, product)
	p.log.Printf("Updated products : %v", product)

	return ctx.Status(http.StatusOK).JSON(product)
}

func (p *ProductHandler) PostProduct(ctx *fiber.Ctx) error {
	p.log.Println("Handle POST Request")

	product := &entity.Product{}

	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}

	entity.AddProduct(product)
	return ctx.Status(http.StatusOK).JSON(product)
}

func (p *ProductHandler) GetProducts(ctx *fiber.Ctx) error {
	p.log.Println("Handle GET Request")

	products := entity.GetProducts()

	return ctx.Status(http.StatusOK).JSON(products)
}
