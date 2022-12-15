package config

import "github.com/1995parham-teaching/record-appender/internal/db"

func Default() Config {
	return Config{
		Database: db.Database{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			DBName:   "record-appender",
			Password: "postgres",
			SSLmode:  "disable",
		},
	}
}
