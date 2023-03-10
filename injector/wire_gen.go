// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injector

import (
	"github.com/MochamadAkbar/ordent-test/handler"
	"github.com/MochamadAkbar/ordent-test/repository"
	"github.com/MochamadAkbar/ordent-test/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Injectors from dependency.go:

func InitializeUserService(conn *pgxpool.Pool, router chi.Router) error {
	userRepository := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUseCase(userRepository)
	error2 := handler.NewUserHandler(userUsecase, router)
	return error2
}
