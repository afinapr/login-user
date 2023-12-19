package usecase

import (
	"fmt"

	mod "login-user/models"
	repo "login-user/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usecase struct {
	Repo repo.Repository
}

func NewUsecase(repo *repo.Repository) *Usecase {
	return &Usecase{Repo: *repo}
}

// CreateUser ... Create User
func (u *Usecase) CreateUser(c *gin.Context) {
	var user mod.User
	var resp mod.Response
	c.BindJSON(&user)
	err := u.Repo.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		resp.Status = "500"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
	} else {
		resp.Status = "200"
		resp.Message = "Sucess retrieve data"
		c.JSON(http.StatusOK, resp)
	}
}

// UpdateUser ... Update the user information
func (u *Usecase) UpdateLocation(c *gin.Context) {
	var loc mod.Location
	var resp mod.Response
	c.BindJSON(&loc)
	err := u.Repo.UpdateLocation(&loc)
	if err != nil {
		fmt.Println(err.Error())
		resp.Status = "500"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
	} else {
		resp.Status = "200"
		resp.Message = "Sucess retrieve data"
		c.JSON(http.StatusOK, resp)
	}
}

func (u *Usecase) Login(c *gin.Context) {
	var login mod.LoginModel
	var resp mod.Response
	c.BindJSON(&login)

	// username := login.Username
	// password := login.Password

	_, err := u.Repo.ReadUsername(login.Username, login.Password)
	if err != nil {
		fmt.Println(err.Error())
		resp.Status = "500"
		resp.Message = err.Error()
		c.JSON(http.StatusInternalServerError, resp)
	} else {
		// result.Username == username && result.Password == password {
		resp.Status = "200"
		resp.Message = "Login Berhasil"
		c.JSON(http.StatusOK, resp)
		// } else if result.Username != username && result.Password != password {
		// 	resp.Status = "400"
		// 	resp.Message = "Username/Password tidak valid"
		// 	c.JSON(http.StatusBadRequest, resp)

		// }
	}
}
