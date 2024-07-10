package config

import (
	"github.com/PetkovaDiana/shop/internal/pkg/pgx"
	"github.com/PetkovaDiana/shop/internal/server"
	"github.com/PetkovaDiana/shop/internal/service"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

const (
	envPath = ".env"
)

type App struct {
	ServerConfig *server.Config     `yaml:"server"`
	DBConfig     *pgx.Config        `yaml:"db"`
	ItemsArgon   service.ItemsArgon `yaml:"password_hash"`
}

func NewAppConfig() (*App, error) {
	if err := godotenv.Load(envPath); err != nil {
		return nil, err
	}

	cfgApp := new(App)

	if err := cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), cfgApp); err != nil {
		return nil, err
	}

	return cfgApp, nil
}
