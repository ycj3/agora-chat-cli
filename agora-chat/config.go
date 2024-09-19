/*
Copyright © 2024 Carlson <carlsonyuandev@gmail.com>
*/
package agora_chat

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const configDir = "agora-chat-cli"

type Config interface {
	GetActiveApp() (*App, error)
	GetApps() (Apps, error)
}

type config struct {
	Path string
	Apps Apps
}

type Apps struct {
	Active string `yaml:"active" mapstructure:"active"`
	Apps   []App  `yaml:"apps" mapstructure:"apps"`
}

type App struct {
	Name           string `yaml:"name" mapstructure:"name"`
	AppID          string `yaml:"app-id" mapstructure:"app-id"`
	AppCertificate string `yaml:"app-certificate" mapstructure:"app-certificate"`
	BaseURL        string `yaml:"base-url" mapstructure:"base-url"`
	AppTokenExp    uint32 `yaml:"app-token-expire" mapstructure:"app-token-expire"`
}

func NewConfig() (Config, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("error getting user config directory: %v", err)
	}
	cfgPath := filepath.Join(dir, configDir, "config.yml")
	// fmt.Println("Config Path is :", cfgPath)
	viper.SetConfigFile(cfgPath)

	if err := viper.ReadInConfig(); err != nil && os.IsNotExist(err) {

		if err := os.MkdirAll(filepath.Dir(cfgPath), 0o755); err != nil {
			return nil, fmt.Errorf("error making dir for config file: %v", err)
		}

		f, err := os.Create(cfgPath)
		if err != nil {
			return nil, fmt.Errorf("error creating config file: %v", err)
		}
		f.Close()
	}

	var apps Apps
	if err := viper.Unmarshal(&apps); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	return &config{Path: cfgPath, Apps: apps}, nil
}

func (cfg *config) GetActiveApp() (*App, error) {
	apps := cfg.Apps
	if len(apps.Apps) == 0 {
		return nil, fmt.Errorf("no app exists")
	}

	for _, app := range apps.Apps {
		if app.Name == apps.Active {
			return &app, nil
		}
	}
	return nil, fmt.Errorf("failed to find active app")
}

func (cfg *config) GetApps() (Apps, error) {
	apps := cfg.Apps
	if len(apps.Apps) == 0 {
		return Apps{}, fmt.Errorf("no app exists")
	}
	return apps, nil
}
