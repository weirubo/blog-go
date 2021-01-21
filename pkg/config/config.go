package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	v *viper.Viper
}

type ServerConfig struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DatabaseConfig struct {
	UserName    string
	Password    string
	Host        string
	DBName      string
	TablePrefix string
	Charset     string
	ParseTime   bool
	Loc         string
}

func NewViper() (*Config, error) {
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../configs/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("NewViper error:", err)
		return nil, err
	}
	config := &Config{
		v: viper,
	}
	return config, nil
}

func (c *Config) ReadConfig(key string, value interface{}) error {
	err := c.v.UnmarshalKey(key, value)
	if err != nil {
		return err
	}
	return nil
}
