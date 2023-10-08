package app

import (
	"dashboard/cmd/dashboard-rest-server/router"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run(port int) {
	app := fiber.New()
	app.Use(logger.New())

	rest := app.Group("/rest")
	v1 := rest.Group("/v1")

	users := v1.Group("/users")
	projects := v1.Group("/tasks")
	tasks := v1.Group("/tasks")

	router.SetupUsers(users)
	router.SetupProjects(projects)
	router.SetupTasks(tasks)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
