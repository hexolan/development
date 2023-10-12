package auth

import (
	
)

// todo: change to actual types
func NewAuthService(cfg bool) AuthService {
	svc := authService{}
	return svc
}

type authService struct {
	Kafka bool
	Postgres bool
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
