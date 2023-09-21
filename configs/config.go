package config

import (
	"github.com/spf13/viper"
)

type Configs struct {
	AWS  AWSConfig
	LOGS LogsConfig
	GIN  GinConfig
}

// AWSConfig representa as configurações da AWS
type AWSConfig struct {
	Region    string `mapstructure:"AWS_REGION"`
	AccessKey string `mapstructure:"AWS_ACCESS_KEY"`
	SecretKey string `mapstructure:"AWS_SECRET_KEY"`
	S3Bucket  string `mapstructure:"AWS_S3_BUCKET"`
}

type LogsConfig struct {
	LogOutput string `mapstructure:"LOG_OUTPUT"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`
}

type GinConfig struct {
	Mode string `mapstructure:"GIN_MODE"`
}

// LoadConfig carrega as configurações do arquivo config.yaml
func LoadConfig(path string) (*Configs, error) {
	var config Configs
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
