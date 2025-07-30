// Package server contains the dependency wiring and HTTP handler setup for the API server.
// The wire function initializes repositories, use cases, controllers, and HTTP routes.
package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nicobellanich/order-management-api/config"
	"github.com/nicobellanich/order-management-api/internal/controllers"
	"github.com/nicobellanich/order-management-api/internal/platform/repository"
	"github.com/nicobellanich/order-management-api/internal/services"
)

// wire sets up all dependencies and returns the HTTP handler mux.
func wire() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Chequea si es un *fiber.Error
			if e, ok := err.(*fiber.Error); ok {
				return c.Status(e.Code).JSON(fiber.Map{
					"error": e.Message,
					"code":  e.Code,
				})
			}

			// Otros errores (no Fiber)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
				"code":  fiber.StatusInternalServerError,
			})
		},
	})

	//Load configuration (environment, etc.)
	conf := config.Load()

	// Infrastructure: initialize repositories based on environment
	ordersRepository, err := repository.NewOrdersRepository(conf)
	if err != nil {
		panic(err)
	}

	// Services : different services uses to help Usecases
	ordersService := services.NewOrderService(ordersRepository)

	// // Use Cases: business logic

	// // Controllers: HTTP handlers
	ordersController := controllers.NewOrderController(ordersService)

	// // HTTP Routes
	app.Get("/order", ordersController.GetAllOrders)
	app.Get("/order/:id", ordersController.GetOrderByID)
	// app.Post("/order", ordersController.Create)
	// app.Get("/user/timeline", ordersController.GetTimeline)
	// app.Post("/user/publish", ordersController.AddPublication)
	// app.Post("/user/following", ordersController.AddFollowing)

	return app
}
