package repository

import (
	"fmt"
	"gcnt/config"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
	"github.com/typesense/typesense-go/typesense"
	"time"
)

var DbInstance *Db

type Db struct {
	Redis     *redis.Pool
	TypeSense *typesense.Client
}

func (d *Db) InitDatabaseInstance(dbType string) {
	log.Info().Caller().Msg("InitDatabaseInstance")

	switch dbType {
	case "redis":
		d.Redis = d.redisClient()
	case "typesense":
		d.TypeSense = d.typeSenseClient()
	default:
	}

	DbInstance = d
	return
}

func (d *Db) redisClient() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", config.Instance.RedisHost, config.Instance.RedisPort),
				redis.DialPassword(config.Instance.RedisPassword),
			)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func (d *Db) typeSenseClient() *typesense.Client {
	client := typesense.NewClient(
		typesense.WithServer(config.Instance.TypesenseAddress),
		typesense.WithAPIKey(config.Instance.TypesenseKey))
	return client
}
