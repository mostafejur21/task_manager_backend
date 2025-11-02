package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mostafejur21/task_manager_backend/config"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s", cnf.User, cnf.Password, cnf.Host, cnf.Port, cnf.Name)

	if !cnf.EnableSSLMode {
		connString += " sslmode=disable"
	}
	return connString
}

func NewDBConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)

	if err != nil {
		return nil, err
	}
	return dbCon, nil

}
