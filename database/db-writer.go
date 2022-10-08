package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// DbExecuteScalar : To insert/update records.
func (dbService DBHandler) DbExecuteScalar(query string, args ...interface{}) (pgx.Rows, error) {
	var err error
	DB, err = dbService.InitDbWriter()
	if err == nil {
		rows, err := DB.Query(context.Background(), query, args...)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
	return nil, err
}

// DbExecuteScalarReturningID : To insert/update records returns ids.
func (dbService DBHandler) DbExecuteScalarReturningID(query string, args ...interface{}) (int, error) {
	var err error
	DB, err = dbService.InitDbWriter()
	returningid := 0
	if err == nil {
		err = DB.QueryRow(context.Background(), query, args...).Scan(&returningid)
		if err != nil {
			return 0, err
		}
		return returningid, nil
	}
	return 0, err
}
