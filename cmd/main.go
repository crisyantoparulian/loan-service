package main

import (
	"fmt"

	"github.com/crisyantoparulian/loansvc/config"
	"github.com/crisyantoparulian/loansvc/driver"
	"github.com/crisyantoparulian/loansvc/generated"
	"github.com/crisyantoparulian/loansvc/handler"
	custMiddleware "github.com/crisyantoparulian/loansvc/middleware"
	"github.com/crisyantoparulian/loansvc/repository"
	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()

	e := echo.New()

	var server generated.ServerInterface = newServer(cfg)

	generated.RegisterHandlers(e, server)
	e.Use(middleware.Logger())
	e.Use(custMiddleware.AuthMiddleware)
	e.Use(custMiddleware.RoleMiddleware())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.App.Port)))
}

func newServer(cfg *config.Config) *handler.Server {
	validator := validator.New()

	repo := repository.NewRepository(repository.NewRepositoryOptions{
		DB: driver.NewGormDatabase(cfg.Database),
	})

	opts := handler.NewServerOptions{
		Repository: repo,
		Validator:  validator,
		Config:     cfg,
	}

	return handler.NewServer(opts)
}
