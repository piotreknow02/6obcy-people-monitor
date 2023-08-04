package config

var Conf *Config 

type Config struct {
	DbFile string
	Port uint16
	AllowOrigins string
}

