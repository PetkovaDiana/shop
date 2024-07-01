package main

import (
	"github.com/PetkovaDiana/shop/internal/config"
	"github.com/PetkovaDiana/shop/internal/handler"
	"github.com/PetkovaDiana/shop/internal/pkg/psql"
	"github.com/PetkovaDiana/shop/internal/repository"
	"github.com/PetkovaDiana/shop/internal/server"
	"github.com/PetkovaDiana/shop/internal/service"
	"log"
)

func main() {
	cfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := psql.NewDB(cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}

	repos := repository.NewRepository(db)
	domain := service.NewService(repos)
	routes := handler.NewHandler(domain)
	srv := server.NewServer(cfg.ServerConfig)

	if err = srv.Run(routes.InitRoutes()); err != nil {
		panic(err)
	}
}
