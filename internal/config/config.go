package config

import (
	"github.com/PetkovaDiana/shop/internal/pkg/psql"
	"github.com/PetkovaDiana/shop/internal/server"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

const (
	envPath = ".env"
)

type App struct {
	ServerConfig *server.Config `yaml:"server"`
	DBConfig     *psql.Config   `yaml:"db"`
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
