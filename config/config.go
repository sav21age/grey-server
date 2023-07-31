package config

import (
	"fmt"
	"time"
	"path/filepath"
	"runtime"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type (
	Config struct {
		Postgres  `yaml:"postgres"`
		Server `yaml:"server"`
		Cipher `yaml:"hash"`
		Log    `yaml:"log"`
	}

	Postgres struct {
		Host   string `env-required:"true" env:"POSTGRES_HOST"`
		Port   string `env-required:"true" env:"POSTGRES_PORT"`
		User     string `env-required:"true" env:"POSTGRES_USER"`
		Password string `env-required:"true" env:"POSTGRES_PASSWORD"`
		DBName   string `env-required:"true" yaml:"db_name"`
		SSLMode  string `env-required:"disable" yaml:"ssl_mode"`
	}

	Server struct {
		Host           string        `env:"SERVER_HOST" env-default:"http://localhost"`
		Port           string        `env:"SERVER_PORT" env-default:"8000"`
		MaxHeaderBytes int           `yaml:"max_header_bytes" env-default:"1"`
		ReadTimeout    time.Duration `yaml:"read_timeout" env-default:"10s"`
		WriteTimeout   time.Duration `yaml:"write_timeout" env-default:"10s"`
	}

	Cipher struct {
		Salt string `env-required:"true" env:"SALT"`
	}

	Log struct {
		Level string `yaml:"level" env-default:"error"`
	}
)

func NewConfig() (*Config, error) {
	_, mainFilePath, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("unable to get self filename path")
	}

	mainDirPath := filepath.Dir(mainFilePath)

	envFilePath := filepath.Join(mainDirPath, "../../.env")
	err := godotenv.Load(envFilePath)
	if err != nil {
		return nil, fmt.Errorf("env load error: %w", err)
	}

	cfg := &Config{}

	yamlFilePath := filepath.Join(mainDirPath, "../../config/config.yaml")
	err = cleanenv.ReadConfig(yamlFilePath, cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	log.Debug().Msg(fmt.Sprintf("%+v", *cfg))

	return cfg, nil
}
