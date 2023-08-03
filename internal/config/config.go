package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config -.
	Config struct {
		App     `yaml:"app"`
		HTTP    `yaml:"http"`
		Log     `yaml:"logger"`
		Mongo   `yaml:"mongo"`
		Elastic `yaml:"elastic"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log -.
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}

	// Mongo
	Mongo struct {
		URI string `env-required:"true" yaml:"uri" env:"MONGO_URI"`
	}

	// Elastic
	Elastic struct {
		Address string `env-required:"true" yaml:"address" env:"ELASTIC_ADDRESS"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
