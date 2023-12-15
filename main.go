package main

import (
	"fmt"
	db "login-user/database"
	router "login-user/router"
	"os"

	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"

	"github.com/subosito/gotenv"
)

var err error

func main() {
	gotenv.Load()
	db.Init()
	// Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	// if err != nil {
	// 	fmt.Println("Status:", err)
	// }

	// r := ro.EndPoint()
	// //running
	// r.Run()

	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.EndPoint(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
