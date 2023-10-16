package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/stocklet/internal/app/order"
	order_v1 "github.com/hexolan/stocklet/internal/pkg/protobuf/order/v1"
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

func getOrderHandler(svc order.OrderRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}
}

func updateOrderHandler(svc order.OrderRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// todo: some sort of validation on request data 
		reqData := new(order_v1.UpdateOrderRequestData)
		if err := c.BodyParser(reqData); err != nil {
			return err // todo: custom error?
		}

		req := &order_v1.UpdateOrderRequest{
			OrderId: c.Params("order_id"),
			UserId: "0",  // todo: get current user from request context
			Data: reqData,
		}

		// Handle the request and return the resposne
		order, err := svc.UpdateOrder(req)
		if err != nil {
			return err
		}

		return c.JSON(order)
		// return c.JSON(fiber.Map{"status": "success", "data": order})
	}
}

func deleteOrderHandler(svc order.OrderRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}
}

func createOrderHandler(svc order.OrderRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	}
}