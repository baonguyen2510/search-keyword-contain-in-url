package configs

import (
	"os"
	"search-keyword-service/common"
	"search-keyword-service/pkg/log"

	"github.com/spf13/viper"
)

type ConfigData struct {
	AppEnv        common.AppEnvType `mapstructure:"APP_ENV"`
	ServerName    string            `mapstructure:"SERVER_NAME"`
	HttpHost      string            `mapstructure:"HTTP_HOST"`
	HttpPort      int               `mapstructure:"HTTP_PORT"`
	SwaggerEnable bool              `mapstructure:"SWAGGER_ENABLE"`

	LoggerDebug     bool `mapstructure:"LOGGER_DEBUG"`
	LoggerSensitive bool `mapstructure:"LOGGER_SENSITIVE"`

	DbConnection  string `mapstructure:"DB_CONNECTION"`
	DbHost        string `mapstructure:"DB_HOST"`
	DbPort        int    `mapstructure:"DB_PORT"`
	DbUsername    string `mapstructure:"DB_USERNAME"`
	DbPassword    string `mapstructure:"DB_PASSWORD"`
	DbName        string `mapstructure:"DB_NAME"`
	DbSchema      string `mapstructure:"DB_SCHEMA"`
	DbLogSQL      bool   `mapstructure:"DB_LOGSQL"`
	DbAutoMigrate bool   `mapstructure:"DB_AUTOMIGRATE"`

	RedisCluster bool `mapstructure:"REDIS_CLUSTER"`
	RedisSingle  bool `mapstructure:"REDIS_SINGLE"`

	RedisSentinel            bool   `mapstructure:"REDIS_SENTINEL"`
	RedisSentinelMasterGroup string `mapstructure:"REDIS_SENTINEL_MASTER_GROUP"`

	RedisAddr     string `mapstructure:"REDIS_ADDR"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`

	BasicAuthUser      string `mapstructure:"BASIC_AUTH_USER"`
	BasicAuthPassword  string `mapstructure:"BASIC_AUTH_PASSWORD"`
	ConfigTimeSchedule int    `mapstructure:"CONFIG_TIME_SCHEDULE"`
	SearchEngineAddr   string `mapstructure:"SEARCH_ENGINE_ADDR"`
}

func (cfg *ConfigData) GetServerEnv() common.AppEnvType {
	if cfg.AppEnv == "" {
		return common.AppEnvDev
	}
	return cfg.AppEnv
}

// Config is the configs for the whole application
var Config *ConfigData

func Init() {
	viper.SetConfigFile(configFilePath())
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Info("Config file not found, app will not load configs from file")
		} else {
			log.Errorw("Config file was found but another error was produced", "error", err)
		}
	}

	viper.AutomaticEnv()
	if err := viper.Unmarshal(&Config); err != nil {
		log.Fatal("error on unmarshal config file, err: ", err)
	}
}

const DefaultServerConfigFile = "configs/config.yaml"

func configFilePath() string {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		return DefaultServerConfigFile
	}
	return configFile
}
