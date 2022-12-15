package config

import (
	"log"
	"strings"

	"github.com/1995parham-teaching/record-appender/internal/db"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

type Config struct {
	Database db.Database `koanf:"database"`
}

func New() Config {
	// Global koanf instance. Use . as the key path delimiter. This can be / or anything.
	k := koanf.New(".")
	// Load default configuration from struct
	if err := k.Load(structs.Provider(Default(), "konaf"), nil); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	// Load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Println("No config file provided")
	}

	// Prefix indicates environments variables prefix
	const Prefix = "record_appender_"

	// Load environment variables
	if err := k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "__", ".")
	}), nil); err != nil {
		log.Println("No env variable provided")
	}

	var out Config

	if err := k.Unmarshal("", &out); err != nil {
		log.Fatalf("Error unmarshalling config: %s", err)
	}

	return out
}
