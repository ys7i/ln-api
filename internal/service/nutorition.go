package service

import (
	"context"
	"database/sql"

	"github.com/ys7i/ln-api/db/daocore"
)

func GetNutoritions(ctx context.Context, txn *sql.Tx) ([]daocore.Nutorition, error) {
	rows, err := txn.Query(
		"SELECT * FROM nutoritions",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nutoritions []daocore.Nutorition
	for rows.Next() {
		n := &daocore.Nutorition{}
		rows.Scan(&n.ID, &n.Effect, &n.ImageURL)
		nutoritions = append(nutoritions, *n)
	}
	return nutoritions, nil
}
