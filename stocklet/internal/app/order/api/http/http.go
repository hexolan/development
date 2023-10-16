package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/stocklet/internal/app/order"
)

func NewHttpAPI(svc order.OrderRepository) {
	app := fiber.New()

	// todo: error middleware
	// todo: improved json marshaller

	app.Post("/", createOrderHandler(svc))
	app.Get("/:orderId", getOrderHandler(svc))
	app.Patch("/:orderId", updateOrderHandler(svc))
	app.Delete("/:orderId", deleteOrderHandler(svc))

	app.Listen(":3000")
}
