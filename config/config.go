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
	JwtRefreshTokenExp int    `envconfig:"JWT_REFRESH_TOKEN_EXPIRATION" required:"true"`

	RedisHost     string `envconfig:"REDIS_HOST" required:"true"`
	RedisPort     string `envconfig:"REDIS_PORT" required:"true"`
	RedisMaxRetry int    `envconfig:"REDIS_MAX_RETRY" required:"true"`

	CacheExpiration int `envconfig:"CACHE_EXPIRATION" required:"true"`
}

func LoadConfig() (*Config, error) {
	// re := regexp.MustCompile(`^(.*` + "mitra-kirim-management-svc" + `)`)
	// cwd, _ := os.Getwd()
	// rootPath := re.Find([]byte(cwd))
	_ = godotenv.Load(".env")
	cnf := Config{}
	err := envconfig.Process("", &cnf)
	return &cnf, err
}
