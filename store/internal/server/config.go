package server

type Config struct {
	Port        string `toml:"port"`
	DatabaseURL string `toml:"database_url"`
	NatsURL     string `toml:"nats_url"`
}

func NewConfig() *Config {
	return &Config{}
}
