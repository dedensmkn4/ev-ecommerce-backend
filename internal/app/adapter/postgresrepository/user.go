package postgresrepository

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain"
	"github.com/dedensmkn4/ev-ecommerce-backend/internal/app/domain/port"
	"github.com/typical-go/typical-rest-server/pkg/sqkit"
	log "github.com/sirupsen/logrus"

)

type userRepositoryImpl struct {
	pg *sql.DB
}

func NewUserRepository(pg *sql.DB) port.UserRepository {
	return &userRepositoryImpl{
		pg: pg,
	}
}

func (u userRepositoryImpl) Find(ctx context.Context, opts ...sqkit.SelectOption) (users []*domain.User, err error) {
	builder := sq.
		Select(
			domain.UserTable.ID,
			domain.UserTable.Name,
			domain.UserTable.Email,
			domain.UserTable.Address,
		).
		From(domain.UserTableName).
		PlaceholderFormat(sq.Dollar).
		RunWith(u.pg)


	for _, opt := range opts {
		builder = opt.CompileSelect(builder)
	}

	rows, err := builder.QueryContext(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	users = make([]*domain.User, 0)
	for  rows.Next(){
		user := new(domain.User)
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Address,
		); err != nil {
			log.Error(err)
			return
		}
		users = append(users, user)
	}
	return
}