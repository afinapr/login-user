package database

import (
	"fmt"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error

	DB, err = sql.Open("mysql", os.Getenv("MYSQL_DIALECTOR"))

	// DB, err = gorm.Open(mysql.Open(), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}

}
