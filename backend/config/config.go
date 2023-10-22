package config

import (
	"github.com/caarlos0/env/v6"
)

// Configは環境変数から取得する設定を表す
// デフォルトの値は開発環境用の値
type Config struct {
	Env        string `env:"ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"50051"`
	DBHost     string `env:"TUNETRAIL_DB_HOST" envDefault:"tunetrail-db"`
	DBPort     int    `env:"TUNETRAIL_DB_PORT" envDefault:"5432"`
	DBUser     string `env:"TUNETRAIL_DB_USER" envDefault:"tunetrail"`
	DBPassword string `env:"TUNETRAIL_DB_PASSWORD" envDefault:"tunetrail"`
	DBName     string `env:"TUNETRAIL_DB_NAME" envDefault:"tunetrail"`
}

// Newは環境変数から設定を取得してConfigを返す
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
