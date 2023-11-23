package clients

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"null.hexolan.dev/exp/utils/config"
	"null.hexolan.dev/exp/utils/errors"
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
