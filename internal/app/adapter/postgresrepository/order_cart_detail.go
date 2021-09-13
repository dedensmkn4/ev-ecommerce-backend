package postgresrepository

import (
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
)

type orderCartDetailRepositoryImpl struct {
	pg *sql.DB
}

func NewOrderCartDetailRepository(pg *sql.DB)  port.OrderCartDetailRepository{
	return &orderCartRepositoryImpl{
		pg: pg,
	}
}