package middleware

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ys7i/ln-api/db/daocore"
)

var nonAuthUrl = []string{"/register", "/login"}

func SetCookie(ctx context.Context, ec echo.Context, session *daocore.UserSession) (echo.Context, error) {
	cookie := new(http.Cookie)
	cookie.Name = "session"
	cookie.Value = session.ID
	cookie.Expires = session.ExpireDate
	ec.SetCookie(cookie)
	return ec, nil
}

func AuthCheck(db *sql.DB) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ec echo.Context) error {
			path := ec.Path()
			for _, p := range nonAuthUrl {
				if p == path {
					return next(ec)
				}
			}
			cookie, err := ec.Cookie("session")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "no credentils found")
			}
			if cookie.Name == "session" {
				_, err := daocore.UserSessionByID(context.Background(), db, cookie.Value)
				if err != nil {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid access")
				}
				return next(ec)
			}
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid access")
		}
	}
}
