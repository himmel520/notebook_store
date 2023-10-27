package main

import (
	"authentication/internal/server"
	"log"

	"github.com/BurntSushi/toml"
)

var configPath string = "configs/server.toml"

func main() {
	config := server.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
