package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or env vars
type Config struct {
	Environment         string        `mapstructure:"ENVIRONMENT"`
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	RedisAddress        string        `mapstructure:"REDIS_ADDRESS"`
	MigrationURL        string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress   string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GrpcServerAddress   string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuarion time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	SmtpSenderName      string        `mapstructure:"SMTP_SENDER_NAME"`
	SmtpSenderAddress   string        `mapstructure:"SMTP_SENDER_ADDRESS"`
	SmtpUsername        string        `mapstructure:"SMTP_USERNAME"`
	SmtpPassword        string        `mapstructure:"SMTP_PASSWORD"`
	SmtpHost            string        `mapstructure:"SMTP_HOST"`
	SmtpPort            int           `mapstructure:"SMTP_PORT"`
}

// LoadConfig reads ocnfiguration from path
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
