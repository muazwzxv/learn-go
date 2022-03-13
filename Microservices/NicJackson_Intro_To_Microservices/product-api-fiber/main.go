package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"muazwzxv/product-api-fiber/handler"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	app := setup()

	go func() {
		if err := app.Listen(":9000"); err != nil {
			log.Printf("Failed to start application: %v", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	_ = <-c

	log.Println("Cleaning up task ...........")
	_ = app.Shutdown()

	// Shutdown process here

	log.Println("Shutdown complete !")
}

func setup() *fiber.App {

	app := fiber.New(fiber.Config{
		ReadTimeout:   5 * time.Second,
		WriteTimeout:  10 * time.Second,
		IdleTimeout:   120 * time.Second,
		ServerHeader:  "Fiber",
		StrictRouting: true,
		CaseSensitive: true,
		AppName:       "Products-API",
	})

	app.Use(logger.New())
	app.Use(cors.New())

	return app
}

func routes(app *fiber.App) {
	log := log.New(os.Stdout, "Product-api", log.LstdFlags)

	productHandler := handler.NewProductHandler(log)
	app.Get("/product", productHandler.GetProducts)
	app.Put("/product/:id", productHandler.UpdateProduct)
	app.Post("product/:id", productHandler.PostProduct)
}
