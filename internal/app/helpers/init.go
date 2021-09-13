package helpers

import (
	"context"
	"database/sql"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/config"
	"time"
)

type HelperInterface interface {
	InitPostgres(cfg *config.Config) *sql.DB
	BeginTrx(ctx context.Context, opts *sql.TxOptions) error
	CommitTrx()
	RollbackTrx()
	QueryContext(ctx context.Context, query string) (rows *sql.Rows, err error)
	QueryRowContext(ctx context.Context, query string) (row *sql.Row)
	SetLogMaxAge(age time.Duration)
	InitLogger()
	CustomRequestLogger()
}

// Helpers ...
type Helpers struct {
	logMaxAge time.Duration
	dbConn    *sql.DB
	dbTrx     *sql.Tx
}

// NewHelper New making new instance of this library
func NewHelper() *Helpers {
	h := &Helpers{}
	return h
}

// SetLogMaxAge is set the maximum log file keep in the path log
func (h *Helpers) SetLogMaxAge(age time.Duration) {
	h.logMaxAge = age
}
