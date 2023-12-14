package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	if DB != nil {
		return
	}
	var err error

	DB, err = gorm.Open(mysql.Open(os.Getenv("MYSQL_DIALECTOR")), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}

}

// // DBConfig represents db configuration
// type DBConfig struct {
// 	Host     string
// 	User     string
// 	DBName   string
// 	Password string
// 	Port     string
// }

// func BuildDBConfig() *DBConfig {
// 	dbConfig := DBConfig{
// 		Host:     os.Getenv("DB_HOST"),
// 		User:     os.Getenv("DB_USER"),
// 		Password: os.Getenv("DB_PASSWORD"),
// 		DBName:   os.Getenv("DB_NAME"),
// 		Port:     os.Getenv("DB_PORT"),
// 	}
// 	return &dbConfig
// }

// Host:     "localhost",
// 		User:     "root",
// 		Password: "",
// 		DBName:   "user_login",
// 		Port:     3306,

// func DbURL(dbConfig *DBConfig) string {
// 	return fmt.Sprintf(
// 		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
// 		dbConfig.User,
// 		dbConfig.Password,
// 		dbConfig.Host,
// 		dbConfig.Port,
// 		dbConfig.DBName,
// 	)
// }
