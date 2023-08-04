package config

import (
	"os"
	"strconv"
)

// Loads cofnig from env or default values
func Load() {
	var loadedConfig Config
	var err error

	loadedConfig.DbFile = os.Getenv("DBFILE")
	if loadedConfig.DbFile == "" {
		loadedConfig.DbFile = "./6obcy.db"
	}

	sport := os.Getenv("PORT")
	port, err := strconv.Atoi(sport)
	loadedConfig.Port = uint16(port)
	if err != nil {
		loadedConfig.Port = 8080
	}

	loadedConfig.AllowOrigins = os.Getenv("ALLOWORIGINS")
	if loadedConfig.AllowOrigins == "" {
		loadedConfig.AllowOrigins = "*"
	}

	Conf = &loadedConfig
}
