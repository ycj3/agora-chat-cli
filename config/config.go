/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const configDir = "agora-chat-cli"

type Apps struct {
	Active string `yaml:"active" mapstructure:"active"`
	Apps   []App  `yaml:"apps" mapstructure:"apps"`
}

type App struct {
	Name           string `yaml:"name" mapstructure:"name"`
	AppID          string `yaml:"app-id" mapstructure:"app-id"`
	AppCertificate string `yaml:"app-certificate" mapstructure:"app-certificate"`
	BaseURL        string `yaml:"base-url" mapstructure:"base-url"`
}

func LoadConfig() (*Apps, error) {
	var configPath string
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("error getting user config directory: %v", err)
	}
	configPath = filepath.Join(dir, configDir, "config.yml")

	// fmt.Println("Config Path is :", configPath)
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil && os.IsNotExist(err) {

		if err := os.MkdirAll(filepath.Dir(configPath), 0o755); err != nil {
			return nil, fmt.Errorf("error making dir for config file: %v", err)
		}

		f, err := os.Create(configPath)
		if err != nil {
			return nil, fmt.Errorf("error creating config file: %v", err)
		}
		f.Close()
	}

	var apps Apps
	if err := viper.Unmarshal(&apps); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return &apps, nil
}
