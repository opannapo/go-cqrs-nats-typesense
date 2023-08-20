package repository

import (
	"fmt"
	"gcnt/config"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
	"time"
)

var DbInstance *Db

type Db struct {
	Redis *redis.Pool
}

func (d *Db) InitDatabaseInstance(dbType string) (err error) {
	log.Info().Caller().Msg("InitDatabaseInstance")

	switch dbType {
	case "redis":
		//redis
		_redis := d.redisClient()
		d.Redis = _redis
	case "typesense":
		//redis
		_redis := d.redisClient()
		d.Redis = _redis
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
