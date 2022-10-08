package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var (
	// DB : for the database at the global space
	DB *pgx.Conn
	// ReaderConnectionString ...
	ReaderConnectionString string
	// WriterConnectionString ...
	WriterConnectionString string
)

//DBHandler provides the class implementation for DbIface interface
type DBHandler struct {
	DatabaseService DbIface
}

// DBNewHandler ...
func DBNewHandler() *DBHandler {
	SetConnectionStrings()
	return &DBHandler{
		DatabaseService: nil,
	}
}

// SetConnectionStrings ...
func SetConnectionStrings() {
	ReaderConnectionString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", "postgres", "postgres", "localhost", 5432, "companiesdb")
	WriterConnectionString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", "postgres", "postgres", "localhost", 5432, "companiesdb")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// InitDbReader : for the database at the global space
func (dbService DBHandler) InitDbReader() (*pgx.Conn, error) {
	var err error
	DB, err = dbService.CreateConnection(ReaderConnectionString)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

// InitDbWriter : for the database at the global space
func (dbService DBHandler) InitDbWriter() (*pgx.Conn, error) {
	DB, err := dbService.CreateConnection(WriterConnectionString)
	if err != nil {
		return nil, err
	}

	return DB, nil
}

// CreateConnection : Creates the Connection
func (dbService DBHandler) CreateConnection(connectionString string) (*pgx.Conn, error) {
	var err error
	DB, err = pgx.Connect(context.Background(), connectionString)
	if err != nil {
		panic(err)
	}
	return DB, nil
}

//DbClose : Close the DB connectivity.
func (dbService DBHandler) DbClose() {
	err := DB.Close(context.Background())
	if err != nil {
		log.Println(err)
	}
}

//TxQuery : To execute a query and fetch rows. This will typically perform an insert & select (or) a plain select.
func (dbService DBHandler) TxQuery(tx pgx.Tx, query string) (pgx.Rows, error) {
	rows, err := tx.Query(context.Background(), query)
	if err != nil {
		if rberror := tx.Rollback(context.Background()); rberror != nil {
			return nil, rberror
		}
		return nil, err
	}

	return rows, nil
}

//TxCreateTempTable : To execute a query and fetch rows. This will typically perform an insert & select (or) a plain select.
func (dbService DBHandler) TxCreateTempTable(tx pgx.Tx, tblName string, sql string) (*pgconn.StatementDescription, error) {
	statDesc, err := tx.Prepare(context.Background(), tblName, sql)
	if err != nil {
		if rberror := tx.Rollback(context.Background()); rberror != nil {
			return nil, rberror
		}
		return nil, err
	}

	return statDesc, nil
}

//TxBegin : To begin transaction.
func (dbService DBHandler) TxBegin() (pgx.Tx, error) {
	var err error
	DB, err = dbService.CreateConnection(WriterConnectionString)
	if err != nil {
		return nil, err
	}

	tx, err := DB.Begin(context.Background())
	return tx, err
}

// TxComplete : Save Changes to the Database.
func (dbService DBHandler) TxComplete(tx pgx.Tx) error {
	if err := tx.Commit(context.Background()); err != nil {
		return err
	}
	return nil
}
