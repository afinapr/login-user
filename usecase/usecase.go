package usecase

import (
	"fmt"

	// mod "login_user/mod"
	mod "login_user/models"
	repo "login_user/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user mod.User
	var resp mod.Response
	c.BindJSON(&user)
	err := repo.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		resp.Status = "200"
		resp.Message = "Sucess retrieve data"
		c.JSON(http.StatusOK, resp)
	}
}

// UpdateUser ... Update the user information
// func UpdateLocation(c *gin.Context) {
// 	var loc mod.Location
// 	var resp mod.Response
// 	c.BindJSON(&loc)
// 	err := repo.UpdateLocation(&loc)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"data": loc})

// 	} else {
// 		resp.Status = "200"
// 		resp.Message = "Sucess retrieve data"
// 		c.JSON(http.StatusOK, loc)
// 	}
// }

func Login(c *gin.Context) {
	var login mod.LoginModel
	var resp mod.ResponseAll
	c.BindJSON(&login)

	username := login.Username
	password := login.Password

	result := repo.ReadUsername(username, password)

	if result.Username != username {
		resp.Status = "400"
		resp.Message = "Username/Password tidak valid"
		c.JSON(http.StatusBadRequest, gin.H{"Status": resp.Status, "Message": resp.Message, "data": result})
		return
	}

	if result.Password == password {
		resp.Status = "200"
		resp.Message = "Login Berhasil"
		c.JSON(http.StatusOK, gin.H{"Status": resp.Status, "Message": resp.Message, "data": result})
	} else {
		resp.Status = "400"
		resp.Message = "Username/Password tidak valid"
		c.JSON(http.StatusBadRequest, gin.H{"Status": resp.Status, "Message": resp.Message, "data": result})
		return

	}

}