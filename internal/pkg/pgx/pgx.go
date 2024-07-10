package pgx

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"time"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	DBName   string `yaml:"dbname"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

func NewDB(ctx context.Context, cfg *Config, maxAttempts int) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	for attempts := 0; attempts < maxAttempts; attempts++ {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		config, err := pgxpool.ParseConfig(dsn)
		if err != nil {
			return nil, fmt.Errorf("unable to parse database URL: %v", err)
		}

		pool, err = pgxpool.NewWithConfig(ctx, config)
		if err == nil {
			return pool, nil
		}

		log.Printf("failed to connect to postgres (attempt %d/%d): %v", attempts+1, maxAttempts, err)
		time.Sleep(2 * time.Second) // небольшая пауза перед следующей попыткой
	}

	return nil, fmt.Errorf("unable to connect to database after %d attempts: %v", maxAttempts, err)
}
