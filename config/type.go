package config

type EnvMode string

type Config struct {
	Server struct {
		Host    string  `json:"host" yaml:"host"`
		Port    string  `json:"port" yaml:"port"`
		Name    string  `json:"name" yaml:"name"`
		Version string  `json:"version" yaml:"version"`
		Mode    EnvMode `json:"mode" yaml:"mode"`
	} `json:"server" yaml:"server"`

	DB struct {
		Port     string `json:"port" yaml:"port"`
		Username string `json:"username" yaml:"username"`
		Host     string `json:"host" yaml:"host"`
		Password string `json:"password" yaml:"password"`
		Name     string `json:"db_name" yaml:"db_name"`
	} `json:"db" yaml:"db"`

	Jwt struct {
		SigningKey             string `json:"signing_key" yaml:"signing_key"`
		TokenExpiration        int    `json:"token_expiration" yaml:"token_expiration"`
		RefreshTokenExpiration int    `json:"refresh_token_expiration" yaml:"refresh_token_expiration"`
		EncryptionKey          string `json:"encryption_key" yaml:"encryption_key"`
	} `json:"jwt" yaml:"jwt"`
}
