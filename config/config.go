package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppName    string `envconfig:"APP_NAME" required:"true"`
	AppPort    string `envconfig:"APP_PORT" required:"true"`
	AppVersion string `envconfig:"APP_VERSION" required:"true"`

	DbHost               string `envconfig:"DATABASE_HOST" required:"true"`
	DbPort               int    `envconfig:"DATABASE_PORT" required:"true"`
	DbUser               string `envconfig:"DATABASE_USER" required:"true"`
	DbPassword           string `envconfig:"DATABASE_PASSWORD" required:"true"`
	DbName               string `envconfig:"DATABASE_NAME" required:"true"`
	DbIdle               int    `envconfig:"DATABASE_IDLE" required:"true"`
	DbMaxConnection      int    `envconfig:"DATABASE_MAX_CONNECTION" required:"true"`
	DbLifetimeConnection int    `envconfig:"DATABASE_LIFETIME_CONNECTION" required:"true"`

	JwtSigningKey      string `envconfig:"JWT_SIGNING_KEY" required:"true"`
	JwtTokenExp        int    `envconfig:"JWT_TOKEN_EXPIRATION" required:"true"`
	JwtRefreshTokenExp string `envconfig:"JWT_REFRESH_TOKEN_EXPIRATION" required:"true"`
	JwtEncKey          string `envconfig:"JWT_ENC_KEY" required:"true"`
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env")
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	return &cnf, err
}
