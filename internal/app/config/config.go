package config

import (
	"os"
	"strconv"
)

type Config struct {
	Address      	string
	ReadTimeout  	string
	WriteTimeout 	string
	Debug        	bool
	DatabaseCfg   	[]DatabaseCfg
}

type DatabaseCfg struct {
	DBName 	string
	DBUser 	string
	DBPass 	string
	Host   	string
	Port   	string
	SslMode	string

	MaxOpenConns    string
	MaxIdleConns    string
	ConnMaxLifetime string
}

func NewConfig() *Config {
	config := &Config{}

	config.Address = os.Getenv("APP_ADDRESS")
	config.ReadTimeout = os.Getenv("APP_READ_TIMEOUT")
	config.WriteTimeout= os.Getenv("APP_WRITE_TIMEOUT")
	config.Debug, _ = strconv.ParseBool(os.Getenv("APP_DEBUG"))

	config.DatabaseCfg = append(config.DatabaseCfg, DatabaseCfg{
		DBName: os.Getenv("PG_DBNAME"),
		DBUser: os.Getenv("PG_DBUSER"),
		DBPass: os.Getenv("PG_DBPASS"),
		Host: os.Getenv("PG_HOST"),
		Port: os.Getenv("PG_PORT"),
		SslMode: os.Getenv("PG_SSL_MODE"),
		MaxOpenConns: os.Getenv("PG_MAX_IDLE_CONNS"),
		MaxIdleConns: os.Getenv("PG_MAX_IDLE_CONNS"),
		ConnMaxLifetime: os.Getenv("PG_CONN_MAX_LIFETIME"),
	})
	return config
}