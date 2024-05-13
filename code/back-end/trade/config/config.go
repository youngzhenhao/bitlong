package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"trade/utils"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`
	Routers struct {
		Login      bool `yaml:"login"`
		FileServer bool `yaml:"file_server"`
		FairLaunch bool `yaml:"fair_launch"`
	} `yaml:"routers"`
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
}

var config Config

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

func GetConfig() *Config {
	loadConfig, err := LoadConfig("config.yaml")
	if err != nil {
		fmt.Println(utils.GetTimeNow(), "[ERROR] Failed to load config: "+err.Error())
		return &Config{}
	}
	return loadConfig
}
