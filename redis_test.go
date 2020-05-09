package fmg

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetRedis(t *testing.T) {
	InitConfig("")
	InitRedis(Config.Redis)
	for _, v := range Config.Redis.Options {
		pool := GetRedis(v.Name)
		r, e := pool.Get().Do("PING")
		if assert.NoError(t, e) {
			assert.Equal(t, "PONG", r)
		}
	}
}
