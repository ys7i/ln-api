package handler

import (
	"database/sql"

	"github.com/ys7i/ln-api/api"
)

type Server struct {
	db *sql.DB
}

var _ api.ServerInterface = &Server{}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}
