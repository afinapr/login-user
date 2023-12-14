package repository

import (
	"fmt"
	Config "login-user/database"
	Models "login-user/models"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(user *Models.User) (err error) {
	query := "INSERT INTO user_location(username, password, latitude, longitude) VALUES (?,?,?,?)"
	err = Config.DB.Exec(query, user.Username, user.Password, user.Latitude, user.Longitude).Error
	if err != nil {
		return fmt.Errorf("create : %v", err.Error())
	}
	return err
}

func UpdateLocation(user *Models.Location) (err error) {
	query := "UPDATE user_location SET longitude = ? , latitude = ? WHERE username = ?"
	err = Config.DB.Exec(query, user.Username, user.Longitude, user.Latitude).Error

	if err != nil {
		return fmt.Errorf("update : %v", err.Error())
	}
	return err
}

func ReadUsername(username string, passsword string) Models.User {
	var user Models.User
	query := Config.DB.Table("user_location").Where("username = ?", username).Where("password = ? ", passsword)
	err := query.Scan(&user).Error
	if err != nil {
		fmt.Println(err.Error())
		return user
		// return fmt.Errorf("get : %v", err.Error())
	}
	return user
}

// GetUserByNamePassword
// func GetUserByNamePassword(user *Models.User) (err error) {
// 	query := Config.DB.Table("user_location").Where("username = ?", user.Username).Where("password = ?", user.Password)
// 	err = query.Scan(&user).Error
// 	if err != nil {
// 		return fmt.Errorf("get : %v", err.Error())
// 	}
// 	return err
// }
