package fmg

import (
	"github.com/fpay/pos-api-go/contracts"
	"github.com/spf13/viper"
)

const (
	DefaultConfigFile = "config.yaml"
)

var Config *ConfigOptions

type ConfigOptions struct {
	Viper *viper.Viper
	*Options
}

type EnvOptions struct {
	AppEnv        string `yaml:"app_env" mapstructure:"app_env"`
	StaticHost    string `yaml:"app_env" mapstructure:"static_host"`
	StaticVersion string `yaml:"app_env" mapstructure:"static_version"`
}

type LogOptions struct {
	Template string `yaml:"template" mapstructure:"template"`
	Root     string `yaml:"root" mapstructure:"root"`
}

type Options struct {
	Env   *EnvOptions
	Log   *LogOptions
	Db    *DbOptions
	Redis *RedisOptions
	Rsa   *contracts.RsaOptions
}

func InitConfig(cfgFile string) {
	if Config != nil {
		return
	}

	if cfgFile == "" {
		cfgFile = DefaultConfigFile
	}

	Config = new(ConfigOptions)
	Config.Viper = viper.New()
	Config.Options = new(Options)

	Config.Viper.SetConfigFile(cfgFile)

	if err := Config.Viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := Config.Viper.Unmarshal(Config.Options); err != nil {
		panic(err)
	}
}
