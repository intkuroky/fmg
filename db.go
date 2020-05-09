package fmg

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbOptions struct {
	Options []*dbSingleOptions `yaml:"options"`
}

type dbSingleOptions struct {
	Name    string `yaml:"name"`
	Driver  string `yaml:"driver"`
	Dsn     string `yaml:"dsn"`
	MaxIdle int    `yaml:"max_idle"`
	MaxActive int    `yaml:"max_active"`
}

type db struct {
	*gorm.DB
}

var dbMap map[string]*db

func InitDb(options *DbOptions) {
	dbMap = make(map[string]*db)
	for _, option := range options.Options {
		gormDb, err := newDb(option)
		if err != nil {
			panic(err)
		}
		dbMap[option.Name] = gormDb
	}
}

func newDb(options *dbSingleOptions) (*db, error) {
	gormDb, err := gorm.Open(options.Driver, options.Dsn)
	if err != nil {
		return nil, err
	}
	gormDb.DB().SetMaxIdleConns(options.MaxIdle)
	gormDb.DB().SetMaxIdleConns(options.MaxActive)
	return &db{gormDb}, nil
}

func GetDb(dbName string) *db {
	return dbMap[dbName]
}

type Account struct {
	gorm.Model
	DeviceNo string `gorm:"column:device_no; type:varchar(100); unique_index"`
	Token    string `gorm:"column:token; type:varchar(100); unique_index"`
	Secret   string `gorm:"column:secret; type:varchar(100)"`
}
