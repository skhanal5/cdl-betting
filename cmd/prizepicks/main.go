package main

import (
	"github.com/caarlos0/env/v11"
	"github.com/skhanal5/prizepicks/internal/common"
	"github.com/skhanal5/prizepicks/internal/database"
)

func main() {
	var cfg common.Config
	err := env.Parse(&cfg)
	if err != nil {
		panic("Failed to read environment variables")
	}
	
	db := database.New(cfg)
	db.GetPlayer("Dashy")
}
