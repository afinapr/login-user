package database

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Init() (*sql.DB, error) {
	DB, err := sql.Open("mysql", os.Getenv("MYSQL_DIALECTOR"))
	if err != nil {
		return nil, err
	}
	fmt.Println("Connection Established")
	return DB, nil
}
