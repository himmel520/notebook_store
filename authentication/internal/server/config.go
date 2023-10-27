package server

type RedisConf struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type Config struct {
	Port        string    `toml:"port"`
	DatabaseURL string    `toml:"database_url"`
	NatsURL     string    `toml:"nats_url"`
	Redis       RedisConf `toml:"redis"`
}

func NewConfig() *Config {
	return &Config{}
}
