package config

import (
	"context"

	"github.com/spf13/viper"
)

const (
	app string = "app"
	env string = "env"
)

type Config struct {
	Environment environment `mapstructure:"ENV"`
	Host        string      `mapstructure:"HOST"`
	Port        string      `mapstructure:"PORT"`
}

// Depedencies
type Deps struct {
	Config
}

func GetConfig(path string) (Config, error) {
	var cfg Config

	viper.AddConfigPath(path)
	viper.SetConfigName(app)
	viper.SetConfigType(env)

	// Load env variables from a .env file if present
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil

}

func GetDeps(ctx context.Context, cfg Config) (Deps, error) {
	var deps Deps
	deps.Config = cfg
	return deps, nil
}
