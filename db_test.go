package fmg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDb(t *testing.T) {
	InitConfig("")
	InitDb(Config.Db)
	for _, v := range Config.Db.Options {
		assert.NoError(t, GetDb(v.Name).Error)
	}
}
