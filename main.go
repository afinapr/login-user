package main

import (
	config "login-user/database"
	router "login-user/router"

	g "github.com/incubus8/go/pkg/gin"

	"github.com/rs/zerolog/log"
	"github.com/subosito/gotenv"
)

var err error

func main() {
	gotenv.Load()
	// Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	// if err != nil {
	// 	fmt.Println("Status:", err)
	// }

	// r := ro.EndPoint()
	// //running
	// r.Run()

	addr := config.Config.ServiceHost + ":" + config.Config.ServicePort
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.EndPoint(),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
