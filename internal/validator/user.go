package validator

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/ys7i/ln-api/api"
)

func ExtractPostRegisterRequest(ec echo.Context) (*api.PostRegisterRequestBody, error) {
	req := &api.PostRegisterRequestBody{}
	err := ec.Bind(req)
	if err != nil {
		return nil, err
	}
	if req.Email == nil {
		return nil, errors.New("email is required")
	}
	if req.Name == nil {
		return nil, errors.New("name is required")
	}
	if req.Password == nil || *req.Password == "" {
		return nil, errors.New("password is required")
	}
	return req, nil
}

func ExtractPostLoginRequest(ec echo.Context) (*api.PostLoginRequestBody, error) {
	req := &api.PostLoginRequestBody{}
	err := ec.Bind(req)
	if err != nil {
		return nil, err
	}
	if req.Password == nil {
		return nil, errors.New("name is required")
	}
	if req.Email == nil {
		return nil, errors.New("email is required")
	}
	return req, nil
}
