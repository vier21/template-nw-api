package config

import (
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	MongoDBURL             string `mapstructure:"MONGODB_URL"`
	MongoUserDBName        string `mapstructure:"MONGODB_DB_USER_NAME"`
	MongoDeviceDBName      string `mapstructure:"MONGODB_DB_DEVICE_NAME"`
	MongoDBMain            string `mapstructure:"MONGODB_DB_MAIN_NAME"`
	ServerAddr             string `mapstructure:"SERVER_ADDR"`
	PlaybookDir            string `mapstructure:"PLAYBOOK_DIR"`
	PlaybookDirCiscoRouter string `mapstructure:"PLAYBOOK_DIR_CISCO_ROUTER"`
	PlaybookDirCiscoSwitch string `mapstructure:"PLAYBOOK_DIR_CISCO_SWITCH"`
	RSAPubPath             string `mapstructure:"RSA_PUB_PATH"`
	RSAPrivPath            string `mapstructure:"RSA_PRIV_PATH"`
	MongoDBUsername        string `mapstructure:"MONGODB_ADMIN_USERNAME"`
	MongoDBPassword        string `mapstructure:"MONGODB_ADMIN_PASSWORD"`
	PostgresDBURL          string `mapstructure:"POSTGRES_URL"`
	PostgresDBUsername     string `mapstructure:"POSTGRES_ADMIN_USERNAME"`
	PostgresDBPassword     string `mapstructure:"MONGODB_ADMIN_PASSWORD"`
	RedisURL               string `mapstructure:"REDIS_URL"`
	RedisUsername          string `mapstructure:"REDIS_ADMIN_USERNAME"`
	RedisPassword          string `mapstructure:"REDIS_ADMIN_PASSWORD"`
}

var env *Config

func init() {

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	path, err := os.Getwd()
	if err != nil {
		logrus.Errorf("Cannot get root project: %s \n", err)
		return
	}

	rootPath := strings.TrimSuffix(path, "/pkg")
	if err := os.Setenv("APP_PATH", rootPath); err != nil {
		logrus.Errorf("Cannot get root project: %s \n", err)
		return
	}
}

func loadConfig(path string) *Config {
	config := Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("error while read env: %s", err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logrus.Errorf("error while read env: %s", err.Error())
	}

	return &config
}

func InitConfig(path string) {
	env = loadConfig(path)
}

func GetConfig() *Config {
	return env
}
