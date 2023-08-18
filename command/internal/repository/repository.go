package repository

import (
	"fmt"
	"gcnt/config"
	"github.com/gomodule/redigo/redis"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Db struct {
	Mysql *gorm.DB
	Redis *redis.Pool
}

func (d *Db) InitDatabase() (err error) {
	log.Info().Caller().Msg("InitDatabase")

	//mysql
	_mysql, err := d.mysqlClient()
	if err != nil {
		log.Err(err).Send()
		return
	}
	d.Mysql = _mysql

	//redis
	_redis := d.redisClient()
	d.Redis = _redis

	return
}

func (d *Db) mysqlClient() (db *gorm.DB, err error) {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.Instance.MySqlUsername,
		config.Instance.MySqlPassword,
		config.Instance.MySqlHost,
		config.Instance.MySqlPort,
		config.Instance.MySqlDb,
	)
	db, err = gorm.Open(mysql.New(mysql.Config{DSN: dns}), &gorm.Config{SkipDefaultTransaction: false, DisableAutomaticPing: false})
	if err != nil {
		log.Err(err).Send()
		return nil, err
	}

	return db, err
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
