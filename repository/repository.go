package repository

import (
	"database/sql"
	"fmt"
	Models "login-user/models"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	Db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{Db: db}
}

func (r *Repository) CreateUser(user *Models.User) (err error) {

	query := "INSERT INTO user_location(username, password, latitude, longitude) VALUES (?,?,?,?)"
	_, err = r.Db.Query(query, user.Username, user.Password, user.Latitude, user.Longitude)
	if err != nil {
		return fmt.Errorf("create : %v", err.Error())
	}
	return err
}

func (r *Repository) UpdateLocation(user *Models.Location) (err error) {
	query := "UPDATE user_location SET longitude = ? , latitude = ? WHERE username = ?"
	_, err = r.Db.Query(query, user.Longitude, user.Latitude, user.Username)
	if err != nil {
		return fmt.Errorf("update : %v", err.Error())
	}
	return err
}

func (r *Repository) ReadUsername(username string, password string) Models.LoginModel {

	var user Models.LoginModel
	err := r.Db.QueryRow("SELECT username, password FROM user_location WHERE username = ? and password = ?", username, password).Scan(&user.Username, &user.Password)
	if err != nil {
		fmt.Println(err.Error())
	}

	return user
}

// func UpdateLocation(long float64, lat float64, username string) err error {
// 	var user Models.Location
// 	query := "UPDATE user_location SET longitude = ? , latitude = ? WHERE username = ?"
// 	err = Config.DB.Exec(query, long, lat, username).Error

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	return err

// }

// func ReadUsername(username string, passsword string) Models.User {
// 	var user Models.User
// 	query := Config.DB.Query("SELECT * FROM ")("user_location").Where("username = ?", username).Where("password = ? ", passsword)
// 	err := query.Scan(&user).Error
// 	if err != nil {
// 		fmt.Println(err.Error())

// 		// return fmt.Errorf("get : %v", err.Error())
// 	}
// 	return user
// }

// GetUserByNamePassword
// func GetUserByNamePassword(user *Models.User) (err error) {
// 	query := Config.DB.Table("user_location").Where("username = ?", user.Username).Where("password = ?", user.Password)
// 	err = query.Scan(&user).Error
// 	if err != nil {
// 		return fmt.Errorf("get : %v", err.Error())
// 	}
// 	return err
// }
