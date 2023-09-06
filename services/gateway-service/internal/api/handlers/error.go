package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
    c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)
    code := fiber.StatusInternalServerError
	msg := err.Error()

    // Retrieval of codes from fiber.Errors
    var e *fiber.Error
    if errors.As(err, &e) {
        code = e.Code
    }

	// Retrival of codes from gRPC errors.
	status, ok := status.FromError(err)
	if ok {
		msg = status.Message()

		switch status.Code() {
		case codes.NotFound:
			code = fiber.StatusNotFound
		case codes.InvalidArgument:
			code = fiber.StatusUnprocessableEntity
		case codes.AlreadyExists:
			code = fiber.StatusConflict
		case codes.PermissionDenied:
			code = fiber.StatusUnauthorized
		case codes.Internal:
			code = fiber.StatusInternalServerError
		default:
			code = fiber.StatusBadGateway
		}
	}
    
    return c.Status(code).JSON(fiber.Map{"status": "failure", "msg": msg})
}