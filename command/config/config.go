package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

var Instance *appConfig

type appConfig struct {
	//app
	AppName string `mapstructure:"APP_NAME" validate:"required"`
	AppEnv  string `mapstructure:"APP_ENV" validate:"required"`
	AppMode string `mapstructure:"APP_MODE" validate:"required"`

	//storage
	MySqlHost     string `mapstructure:"MYSQL_HOST" validate:"required"`
	MySqlPort     int    `mapstructure:"MYSQL_PORT" validate:"required"`
	MySqlDb       string `mapstructure:"MYSQL_DB" validate:"required"`
	MySqlUsername string `mapstructure:"MYSQL_USERNAME" validate:"required"`
	MySqlPassword string `mapstructure:"MYSQL_PASSWORD" validate:"required"`
	RedisHost     string `mapstructure:"REDIS_HOST" validate:"required"`
	RedisPort     int    `mapstructure:"REDIS_PORT" validate:"required"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`

	//services
	NatsAddress        string `mapstructure:"NATS_ADDRESS" validate:"required"`
	CommandServiceHost string `mapstructure:"COMMAND_SERVICE_HOST" validate:"required"`
	CommandServicePort int    `mapstructure:"COMMAND_SERVICE_PORT" validate:"required"`
	QueryServiceHost   string `mapstructure:"QUERY_SERVICE_HOST" validate:"required"`
	QueryServicePort   int    `mapstructure:"QUERY_SERVICE_PORT" validate:"required"`
}

func InitConfigInstance() (err error) {
	cfg := appConfig{}

	err = godotenv.Load()
	if err != nil {
		panic(err)
	}

	//app
	viper.Set(getOsEnv("APP_NAME"))
	viper.Set(getOsEnv("APP_ENV"))
	viper.Set(getOsEnv("APP_MODE"))

	//storage
	viper.Set(getOsEnv("MYSQL_HOST"))
	viper.Set(getOsEnv("MYSQL_PORT"))
	viper.Set(getOsEnv("MYSQL_DB"))
	viper.Set(getOsEnv("MYSQL_USERNAME"))
	viper.Set(getOsEnv("MYSQL_PASSWORD"))
	viper.Set(getOsEnv("REDIS_HOST"))
	viper.Set(getOsEnv("REDIS_PORT"))
	viper.Set(getOsEnv("REDIS_PASSWORD"))

	//services
	viper.Set(getOsEnv("NATS_ADDRESS"))
	viper.Set(getOsEnv("COMMAND_SERVICE_HOST"))
	viper.Set(getOsEnv("COMMAND_SERVICE_PORT"))
	viper.Set(getOsEnv("QUERY_SERVICE_HOST"))
	viper.Set(getOsEnv("QUERY_SERVICE_PORT"))

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(fmt.Errorf("err Failed to load config: %+v", err))
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		panic(fmt.Errorf("err invalid configuration %s", err))
	}

	Instance = &cfg
	log.Printf("Config %+v", cfg)

	return
}

func getOsEnv(name string) (key string, value interface{}) {
	key = name
	value = os.Getenv(name)
	return
}
