package config

type Config struct {
	// PostgreSQL
	DBName string `mapstructure:"db-name"`
	Stage  string `mapstructure:"stage"`
}
