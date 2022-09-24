package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/ys7i/ln-api/api"
	"github.com/ys7i/ln-api/internal/handler"
	"github.com/ys7i/ln-api/internal/middleware"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
	return
}

func run() error {
	e := echo.New()
	e.HideBanner = true
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return err
	}
	server := handler.NewServer(db)
	api.RegisterHandlers(e, server)
	e.Use(echoMiddleware.CORS())
	e.Use(middleware.AuthCheck(db))
	e.Start(fmt.Sprintf("127.0.0.1:%v", 5050))
	return nil
}
