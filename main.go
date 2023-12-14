package main

import (
	"fmt"
	Config "login_user/database"
	ro "login_user/router"

	"github.com/jinzhu/gorm"
	"github.com/subosito/gotenv"
)

var err error

func main() {
	gotenv.Load()
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	r := ro.EndPoint()
	//running
	r.Run()
}
