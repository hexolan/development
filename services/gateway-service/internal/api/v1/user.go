package v1

import (
	"time"
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/rpc"
	"github.com/hexolan/panels/gateway-service/internal/rpc/userv1"
	"github.com/hexolan/panels/gateway-service/internal/api/handlers"
)

type userSignupForm struct {
	Username string
	Password string
}

func getUserByUsername(username string) (*userv1.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	user, err := rpc.Svcs.GetUserSvc().GetUserByName(
		ctx,
		&userv1.GetUserByNameRequest{Username: username},
	)

	return user, err
}

func GetUserByUsername(c *fiber.Ctx) error {
	user, err := getUserByUsername(c.Params("username"))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

func GetCurrentUser(c *fiber.Ctx) error {
	// access token claims
	tokenClaims, err := handlers.GetTokenClaims(c)
	if err != nil {
		return err
	}

	// make RPC call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	user, err := rpc.Svcs.GetUserSvc().GetUser(
		ctx,
		&userv1.GetUserByIdRequest{Id: tokenClaims.Subject},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": user})
}

func DeleteCurrentUser(c *fiber.Ctx) error {
	// access token claims
	tokenClaims, err := handlers.GetTokenClaims(c)
	if err != nil {
		return err
	}

	// make RPC call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_, err = rpc.Svcs.GetUserSvc().DeleteUser(
		ctx,
		&userv1.DeleteUserByIdRequest{Id: tokenClaims.Subject},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "user deleted"})
}

func UserSignup(c *fiber.Ctx) error {
	// Parse the body data
	form := new(userSignupForm)
	if err := c.BodyParser(form); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}
	
	// todo: defer this logic away from gateway-service in future
	// Attempt to create the user
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	user, err := rpc.Svcs.GetUserSvc().CreateUser(
		ctx,
		&userv1.CreateUserRequest{
			Data: &userv1.UserMutable{
				Username: &form.Username,
			},
		},
	)
	if err != nil {
		return err
	}

	// Attempt to set password auth method for the user
	err = setAuthMethod(user.Id, form.Password)
	if err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		_, _ = rpc.Svcs.GetUserSvc().DeleteUser(
			ctx,
			&userv1.DeleteUserByIdRequest{Id: user.Id},
		)
		return err
	}

	// Signup success - attempt to get auth token
	token, _ := authWithPassword(user.Id, form.Password)
	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"user": user,
			"token": token,
		},
	})
}