package main

import (
	"ginjwt/conf"
	"log"
)

func main() {
	cfg := conf.Load("../config.yaml")

	r := initApp(&cfg.DB)

	if err := r.App.Run(cfg.Server.Address); err != nil {
		log.Fatalf("run server at %s failed: %v", cfg.Server.Address, err)
	}
}