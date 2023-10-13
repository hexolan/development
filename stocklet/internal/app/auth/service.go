package auth

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// todo: change to actual types
func NewAuthService(db *pgxpool.Pool, prod EventProducer) AuthService {
	svc := authService{
		db: db,
		prod: prod,
	}
	return svc
}

type authService struct {
	db *pgxpool.Pool
	prod EventProducer
}

type AuthService interface {
	AuthUserPassword(username string, password string) (*bool, error)
	AddMethodPassword(username string, password string) (*bool, error)
	UpdateMethodPassword(username string, password string) (*bool, error)
	DeleteMethodPassword(username string) (*bool, error)
}

func (svc authService) AuthUserPassword(username string, password string) (*bool, error) {
	return nil, nil
}

func (svc authService) AddMethodPassword(username string, password string) (*bool, error) {
	return nil, nil
}

func (svc authService) UpdateMethodPassword(username string, password string) (*bool, error) {
	return nil, nil
}

func (svc authService) DeleteMethodPassword(username string) (*bool, error) {
	return nil, nil
}
