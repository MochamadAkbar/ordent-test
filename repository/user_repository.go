// Package repository is describe all action to database
package repository

import (
	"context"

	"github.com/MochamadAkbar/ordent-test/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	Register(ctx context.Context, user *entity.User) (entity.User, bool)
	Login(ctx context.Context, user *entity.User) (entity.User, bool)
}

type UserRepositoryImpl struct {
	Conn *pgxpool.Pool
}

func (repository *UserRepositoryImpl) Register(ctx context.Context, user *entity.User) (entity.User, bool) {
	var row entity.User
	statement := `INSERT INTO "users" ("name", "email", "password") VALUES($1, $2, $3) RETURNING "id";`

	err := repository.Conn.QueryRow(ctx, statement, user.Name, user.Email, user.Password).Scan(&row.ID)
	if err != nil {
		panic(err.Error())
	}

	return row, true
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, user *entity.User) (entity.User, bool) {
	var result entity.User
	statement := `SELECT "id", "email", "password" FROM "users" WHERE "email" = $1;`

	err := repository.Conn.QueryRow(ctx, statement, user.Email).
		Scan(&result.ID, &result.Email, &result.Password)
	if err != nil {
		panic(err.Error())
	}

	return result, true
}
