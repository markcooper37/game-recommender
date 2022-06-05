package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port int    `envconfig:"PORT" default:"8080"`
	DSN  string `envconfig:"DATABASE_URL" default:"postgresql://user:password@localhost:5432/game-recommender?sslmode=disable"`
}

func Load() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}
	return &c, nil
}
