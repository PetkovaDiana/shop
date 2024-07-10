package main

import (
	"context"
	"github.com/PetkovaDiana/shop/internal/config"
	"github.com/PetkovaDiana/shop/internal/handler"
	"github.com/PetkovaDiana/shop/internal/pkg/pgx"
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

	ctx := context.Background()
	db, err := pgx.NewDB(ctx, cfg.DBConfig, 3)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	defer db.Close()

	repos := repository.NewRepository(db)
	domain := service.NewService(ctx, repos, cfg.ItemsArgon)
	routes := handler.NewHandler(domain)
	srv := server.NewServer(cfg.ServerConfig)

	if err = srv.Run(routes.InitRoutes()); err != nil {
		panic(err)
	}
}
