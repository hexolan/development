package v1

import (
	"time"
	"context"

	"github.com/gofiber/fiber/v2"

	"github.com/hexolan/panels/gateway-service/internal/rpc"
	"github.com/hexolan/panels/gateway-service/internal/rpc/panelv1"
)

func CreatePanel(c *fiber.Ctx) error {
	// todo: check user can create panels

	newPanel := new(panelv1.PanelMutable)
	if err := c.BodyParser(newPanel); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().CreatePanel(
		ctx,
		&panelv1.CreatePanelRequest{Data: newPanel},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}

func GetPanelByName(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().GetPanelByName(
		ctx,
		&panelv1.GetPanelByNameRequest{Name: c.Params("name")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}

func UpdatePanelByName(c *fiber.Ctx) error {
	// todo: check user can update panels

	patchData := new(panelv1.PanelMutable)
	if err := c.BodyParser(patchData); err != nil {
		fiber.NewError(fiber.StatusBadRequest, "malformed request")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().UpdatePanelByName(
		ctx, 
		&panelv1.UpdatePanelByNameRequest{Name: c.Params("name"), Data: patchData},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "data": panel})
}

func DeletePanelByName(c *fiber.Ctx) error {
	// todo: check user can delete panels

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := rpc.Svcs.GetPanelSvc().DeletePanelByName(
		ctx,
		&panelv1.DeletePanelByNameRequest{Name: c.Params("name")},
	)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"status": "success", "msg": "panel deleted"})
}

func getPanelIDFromName(panelName string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	panel, err := rpc.Svcs.GetPanelSvc().GetPanelByName(
		ctx,
		&panelv1.GetPanelByNameRequest{Name: panelName},
	)
	if err != nil {
		return "", err
	}
	
	return panel.GetId(), nil
}