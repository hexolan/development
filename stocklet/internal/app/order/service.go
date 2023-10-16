package order

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewOrderService(db *pgxpool.Pool, prod EventProducer) OrderService {
	svc := orderService{
		db: db,
		prod: prod,
	}
	return svc
}


type OrderService interface {

}

type orderService struct {
	db *pgxpool.Pool
	prod EventProducer
}

func (svc orderService) DoSomething() {
	
}