package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/auth_service/config"
	"github.com/auth_service/storage"
)

type Store struct {
	db   *pgxpool.Pool
	user *UserRepo
}

func NewPostgres(ctx context.Context, cfg config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:   pool,
		user: NewUserRepo(pool),
	}, err
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

// func (s *Store) Book() storage.BookRepoI {

// 	if s.book == nil {
// 		s.book = NewBookRepo(s.db)
// 	}

// 	return s.book
// }

// func (s *Store) Order() storage.OrderRepoI {

// 	if s.order == nil {
// 		s.order = NewOrderRepo(s.db)
// 	}

// 	return s.order
// }
