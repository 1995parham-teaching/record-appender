package config

func Default() Config {
	return Config{
		Database: Database{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "postgres",
			DBName:   "record-appender",
			Password: "postgres",
			SSLmode:  "disable",
		},
	}
}
