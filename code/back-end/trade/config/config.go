package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"trade/utils"
)

type Config struct {
	GinConfig struct {
		Bind string `yaml:"bind"`
		Port string `yaml:"port"`
	} `yaml:"gin_config"`
	GormConfig struct {
		Mysql struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			DBName   string `yaml:"dbname"`
		} `yaml:"mysql"`
		Redis struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			DB       int    `yaml:"db"`
		} `yaml:"redis"`
	} `yaml:"gorm_config"`
	Routers struct {
		Login      bool `yaml:"login"`
		FileServer bool `yaml:"file_server"`
		FairLaunch bool `yaml:"fair_launch"`
		Ping       bool `yaml:"ping"`
	} `yaml:"routers"`
	ApiConfig struct {
		Lnd struct {
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			TlsCertPath  string `yaml:"tls_cert_path"`
			MacaroonPath string `yaml:"macaroon_path"`
		} `yaml:"lnd"`
		Tapd struct {
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			TlsCertPath  string `yaml:"tls_cert_path"`
			MacaroonPath string `yaml:"macaroon_path"`
		} `yaml:"tapd"`
		Litd struct {
			Host         string `yaml:"host"`
			Port         int    `yaml:"port"`
			TlsCertPath  string `yaml:"tls_cert_path"`
			MacaroonPath string `yaml:"macaroon_path"`
		} `yaml:"litd"`
		CustodyAccount struct {
			MacaroonDir string `yaml:"macaroon_dir"`
		} `yaml:"custody_account"`
	} `yaml:"api_config"`
	Bolt struct {
		DbPath          string `yaml:"db_path"`
		DbMode          uint32 `yaml:"db_mode"`
		DbTimeoutSecond int64  `yaml:"db_timeout_second"`
	} `yaml:"bolt"`
	AdminUsers []BasicAuth `yaml:"admin_users"`
}

type BasicAuth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var (
	config Config
)

func GetConfig() *Config {
	return &config
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func GetLoadConfig() *Config {
	loadConfig, err := LoadConfig("config.yaml")
	if err != nil {
		utils.LogError("[ERROR] Failed to load config", err)
		return &Config{}
	}
	return loadConfig
}
