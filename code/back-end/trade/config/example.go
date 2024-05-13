package config

import (
	"gopkg.in/yaml.v3"
	"path/filepath"
	"trade/utils"
)

func generateConfig() string {
	var emptyConfig = Config{AdminUsers: make([]BasicAuth, 0)}
	emptyConfig.AdminUsers = append(emptyConfig.AdminUsers, BasicAuth{})
	conf, _ := yaml.Marshal(&emptyConfig)
	return string(conf)
}

func WriteConfigExample(dir string) {
	utils.CreateFile(filepath.Join(dir, "config.yaml.example"), generateConfig())
}
