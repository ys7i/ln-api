package handler

import (
	"context"
	"net/http"

	"github.com/ys7i/ln-api/api"
	"github.com/ys7i/ln-api/internal/middleware"
	"github.com/ys7i/ln-api/internal/service"
	"github.com/ys7i/ln-api/internal/validator"
	"github.com/ys7i/ln-api/pkg/echoutil"

	"github.com/labstack/echo/v4"
)

func (s *Server) PostRegister(ec echo.Context) error {
	ctx := context.Background()
	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	req, err := validator.ExtractPostRegisterRequest(ec)

	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	err = service.CreateUser(ctx, txn, req)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	txn.Commit()
	ok := "ok"
	return ec.JSON(http.StatusOK, &api.Message{Message: &ok})
}

func (s *Server) PostLogin(ec echo.Context) error {
	ctx := context.Background()
	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()

	req, err := validator.ExtractPostLoginRequest(ec)
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	userSession, err := service.CreateUserSession(ctx, txn, req)
	if err != nil {
		return err
	}
	ec, err = middleware.SetCookie(ctx, ec, userSession)
	txn.Commit()
	ok := "ok"
	return ec.JSON(http.StatusOK, &api.Message{Message: &ok})
}
