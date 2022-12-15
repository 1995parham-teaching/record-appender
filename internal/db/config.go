package db

import "fmt"

type Database struct {
	Host     string `koan:"host"`
	Port     string `koanf:"port"`
	User     string `koanf:"user"`
	DBName   string `koanf:"dbname"`
	Password string `koanf:"password"`
	SSLmode  string `koanf:"sslmode"`
}

func (d Database) Cstring() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=%s",
		d.Host, d.Port, d.User, d.DBName, d.Password, d.SSLmode)
}
