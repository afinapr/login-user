package main

import (
	"fmt"
	"login-user/database"
	"login-user/repository"
	router "login-user/router"
	"login-user/usecase"
	"os"

	g "github.com/incubus8/go/pkg/gin"
	"github.com/rs/zerolog/log"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	db, err := database.Init()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer db.Close()
	r := repository.NewRepository(db)
	uc := usecase.NewUsecase(r)

	addr := fmt.Sprintf("%s:%s", os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT"))
	conf := g.Config{
		ListenAddr: addr,
		Handler:    router.EndPoint(uc),
		OnStarting: func() {
			log.Info().Msg("Your service is up and running at " + addr)
		},
	}

	g.Run(conf)
}
