package config

import (
	"github.com/lapkomo2018/goTwitterServices/internal/auth/transport/grpc"
	"github.com/lapkomo2018/goTwitterServices/internal/auth/transport/rest"
	"github.com/lapkomo2018/goTwitterServices/pkg/hash"
	"github.com/lapkomo2018/goTwitterServices/pkg/jwt"
	"github.com/lapkomo2018/goTwitterServices/pkg/validation"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Service struct {
		Name string
		Host string
	}
	Discovery struct {
		Address string
	}
	RestServer rest.Config
	GrpcServer grpc.Config
	JWT        jwt.Config
	Hash       hash.Config
	Validator  validation.Config
}

func LoadConfig() (*Config, error) {
	log.Println("Loading config...")

	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	log.Println("Loaded config")
	return &config, nil
}
