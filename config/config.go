package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
	"github.com/rs/zerolog"
)

type Dapr struct {
	HTTPPort int `env:"HTTP_PORT" envDefault:"3500"`
	GRPCPort int `env:"GRPC_PORT" envDefault:"50001"`
}

type Log struct {
	Level zerolog.Level `env:"LEVEL" envDefault:"info"`
}

type Config struct {
	Dapr        Dapr   `envPrefix:"DAPR_"`
	Log         Log    `envPrefix:"LOG_"`
	Port        int    `env:"PORT" envDefault:"6000"`
	BaseURL     string `env:"BASE_URL" envDefault:"http://localhost:6000"`
	DepositPath string `env:"DEPOSIT_URI" envDefault:"/deposit"`
}

func (c *Config) Address() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *Config) DepositURI() string {
	return c.BaseURL + c.DepositPath
}

func Read() (*Config, error) {
	cfg := new(Config)
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return cfg, err
	}
	return cfg, nil
}
