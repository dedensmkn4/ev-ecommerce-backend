package postgresdb

import (
	"database/sql"
	"fmt"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/config"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type PgDb struct {
	config *config.Config
	Pg *sql.DB `name:"pg"`
}

func NewPgDb(cfg *config.Config) *PgDb {
	return &PgDb{
		config: cfg,
		Pg: openPostgres(cfg),
	}
}

func openPostgres(p *config.Config) *sql.DB {
	conn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		p.DatabaseCfg[0].DBUser, p.DatabaseCfg[0].DBPass, p.DatabaseCfg[0].Host, p.DatabaseCfg[0].Port, p.DatabaseCfg[0].DBName,
	)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		logrus.Fatalf("postgres: %s", err.Error())
	}
	maxOpenConnes, _ := strconv.Atoi(p.DatabaseCfg[0].MaxOpenConns)
	maxIdleConnes, _ := strconv.Atoi(p.DatabaseCfg[0].MaxIdleConns)
	connMaxLifetime, _ := strconv.Atoi(p.DatabaseCfg[0].ConnMaxLifetime)

	db.SetMaxOpenConns(maxOpenConnes)
	db.SetMaxIdleConns(maxIdleConnes)
	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) *time.Second)


	if err = db.Ping(); err != nil {
		logrus.Fatalf("postgres: %s", err.Error())
	}

	return db
}