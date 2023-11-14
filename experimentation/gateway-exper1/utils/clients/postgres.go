package clients

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/hexolan/development/experimentation/gateway-exper1/utils/config"
	"github.com/hexolan/development/experimentation/gateway-exper1/utils/errors"
)

func NewPostgresConn(conf *config.PostgresConfig) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*45)
	defer cancel()

	pCl, err := pgxpool.New(ctx, conf.GetDSN())
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeExtService, "failed to connect to postgres", err)
	}
	return pCl, nil
}
