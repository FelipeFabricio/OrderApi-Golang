package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	DBDriver         string `mapstructure:"DB_DRIVER"`
	DBUser           string `mapstructure:"DB_USER"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBName           string `mapstructure:"DB_NAME"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBPort           string `mapstructure:"DB_PORT"`
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	RabbitMQUser     string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPassword string `mapstructure:"RABBITMQ_PASSWORD"`
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`
}

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, nil
}
