package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run(port int) {
	app := fiber.New()
	app.Use(logger.New())

	rest := app.Group("/rest")
	_ = rest.Group("/v1")

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
