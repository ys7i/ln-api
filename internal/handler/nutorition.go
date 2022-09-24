package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ys7i/ln-api/api"
	"github.com/ys7i/ln-api/internal/service"
	"github.com/ys7i/ln-api/pkg/echoutil"
)

func (s *Server) GetNutoritions(ec echo.Context) error {
	ctx := context.Background()
	txn, err := s.db.Begin()
	if err != nil {
		return echoutil.ErrInternal(ec, err)
	}
	defer txn.Rollback()
	nutoritions, err := service.GetNutoritions(ctx, txn)
	if err != nil {
		return err
	}
	var responses []api.Nutorition
	for _, n := range nutoritions {
		responses = append(responses, api.Nutorition{Id: &n.ID, Effect: &n.Effect.String, ImageUrl: &n.ImageURL.String})
	}
	return ec.JSON(http.StatusOK, responses)
}
