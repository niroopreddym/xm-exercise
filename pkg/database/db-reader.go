package database

import (
	"context"
)

// DbQuery : To insert/update records.
func (dbService DBHandler) DbQuery(query string, responseRowID *int, args ...interface{}) error {
	DB, err := dbService.InitDbReader()
	if err != nil {
		return err
	}

	err = DB.QueryRow(context.Background(), query, args...).Scan(&responseRowID)
	if err != nil {
		panic(err)
	}

	return nil
}
