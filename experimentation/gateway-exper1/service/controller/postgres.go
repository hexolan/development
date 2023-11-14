package controller

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/development/experimentation/gateway-exper1/service"
)

type postgresController struct {
	pCl *pgxpool.Pool
	evtC service.EventController
}

func NewPostgresController(pCl *pgxpool.Pool, evtC service.EventController) service.StorageController {
	return postgresController{pCl: pCl, evtC: evtC}
}