package fmg

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisOptions struct {
	Options []*redisSingleOptions `yaml:"options"`
}
type redisSingleOptions struct {
	Name     string `yaml:"name"`
	Addr     string `yaml:"addr" mapstructure:"addr"`
	Password string `yaml:"password" mapstructure:"password"`
	Db		 int    `yaml:"db"`
	MaxIdle  int    `yaml:"max_idle"`
	MaxActive  int    `yaml:"max_active"`
}

type redisPool struct {
	*redis.Pool
}

var redisMap map[string]*redisPool

func InitRedis(options *RedisOptions) {
	redisMap = make(map[string]*redisPool)
	for _, option := range options.Options {
		redisCache, err := NewRedis(option)
		if err != nil {
			panic(err)
		}
		redisMap[option.Name] = redisCache
	}
}

func NewRedis(options *redisSingleOptions) (*redisPool, error) {
	return &redisPool{
		&redis.Pool{
			MaxIdle:     options.MaxIdle,
			MaxActive:   options.MaxActive,
			IdleTimeout: 240 * time.Second,
			Dial: func () (redis.Conn, error) {
				c, err := redis.Dial("tcp", options.Addr)
				if err != nil {
					return nil, err
				}
				if options.Password != "" {
					if _, err := c.Do("AUTH", options.Password); err != nil {
						c.Close()
						return nil, err
					}
				}
				if options.Db != 0 {
					if _, err := c.Do("SELECT", options.Db); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, nil
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
		},
	}, nil
}

func GetRedis(redisName string) *redisPool {
	return redisMap[redisName]
}
