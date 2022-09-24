package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/ys7i/ln-api/api"
	"github.com/ys7i/ln-api/db/daocore"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx context.Context, txn *sql.Tx, req *api.PostRegisterRequestBody) error {
	bytePassword := []byte(*req.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, 10)
	if err != nil {
		return err
	}
	user := &daocore.User{
		Name:         *req.Name,
		Email:        *req.Email,
		PasswordHash: string(hashedPassword),
	}
	err = user.Insert(ctx, txn)
	if err != nil {
		return err
	}
	return nil
}

func CreateUserSession(ctx context.Context, txn *sql.Tx, req *api.PostLoginRequestBody) (*daocore.UserSession, error) {
	user, err := daocore.UserByEmail(ctx, txn, *req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(*req.Password))
	if err != nil {
		return nil, err
	}
	uuidObj, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	userSession := daocore.UserSession{ID: uuidObj.String(), ExpireDate: time.Now().Add(3 * time.Hour), UserID: user.ID}
	err = userSession.Upsert(ctx, txn)
	if err != nil {
		return nil, err
	}
	return &userSession, nil
}
