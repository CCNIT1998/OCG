package main

import (
	// "fmt"
	"github.com/CCNIT1998/OCG/REST_API/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
	})

	app.Static("/public", "./public", fiber.Static{ //http://localhost:3000/public OR http://localhost:3000/public/dog.jpeg
		Compress:  true,
		ByteRange: true,
		Browse:    true,
		MaxAge:    3600,
	})

	userRouter := app.Group("/api/user")
	routes.ConfigUserRouter(&userRouter) //http://localhost:3000/api/user

	ProductRouter := app.Group("/api/product")
	routes.ConfigProductRouter(&ProductRouter) //http://localhost:3000/api/product

	CategoryRouter := app.Group("/api/category")
	routes.ConfigCategoryRouter(&CategoryRouter) //http://localhost:3000/api/category

	app.Listen(":3000")

}
