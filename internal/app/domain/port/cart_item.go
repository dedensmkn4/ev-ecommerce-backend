package port

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
)

type CartItemRepository interface {
	Find(ctx context.Context, opts ...sqkit.SelectOption) (cartItems []*domain.CartItem, err error)
	Insert(ctx context.Context,  dbTx *sql.Tx, ent *domain.CartItem) (int, error)
	Delete(ctx context.Context,  dbTx *sql.Tx, opts ...sqkit.DeleteOption) (int, error)
}
