package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mode string `mapstructure:"GIN_MODE"`
	Port string `mapstructure:"GIN_PORT"`
}

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

func LoadConfig(path string) (*Config, error) {
	var config Config
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

// LoadAWSConfig carrega apenas as configurações da AWS
func LoadAWSConfig(path string) (*AWSConfig, error) {
	var awsConfig AWSConfig
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&awsConfig); err != nil {
		return nil, err
	}

	return &awsConfig, nil
}
