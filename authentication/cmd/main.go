package main

import (
	log "github.com/himmel520/notebook_store/authentication/internal/logger"
	"github.com/himmel520/notebook_store/authentication/internal/server"

	"github.com/BurntSushi/toml"
)

var configPath string = "configs/server.toml"

func main() {
	config := server.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Logger.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Logger.Fatal(err)
	}
}
