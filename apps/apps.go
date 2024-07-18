package apps

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	configDir = "agora-chat-cli"
)

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

func (c *Apps) Get(appID string) (*App, error) {
	if len(c.Apps) == 0 {
		return nil, errors.New("no application configured, please run `agorachat apps -c` to add a new one")
	}

	for _, app := range c.Apps {
		if app.AppID == appID {
			return &app, nil
		}
	}
	return nil, fmt.Errorf("application:[%q] doesn't exist", appID)
}

func (c *Apps) GetActiveAppOrExplicit(cmd *cobra.Command) (*App, error) {
	appID := c.Active
	explicit, err := cmd.Flags().GetString("app")
	if err != nil {
		return nil, err
	}
	if explicit != "" {
		appID = explicit
	}

	a, err := c.Get(appID)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (c *Apps) RunQuestionnaire() error {
	var newApp App
	err := survey.Ask(questions(), &newApp)
	if err != nil {
		return err
	}
	err = c.Add(newApp)
	if err != nil {
		return err
	}
	fmt.Println("Application successfully added. 🚀")
	return nil
}

func (c *Apps) Add(newApp App) error {
	if len(c.Apps) == 0 {
		c.Active = newApp.AppID
	}

	for _, app := range c.Apps {
		if newApp.AppID == app.AppID {
			return fmt.Errorf("[%s]Application already exists", newApp.AppID)
		}
	}

	c.Apps = append(c.Apps, newApp)

	viper.Set("active", c.Active)
	viper.Set("apps", c.Apps)
	return viper.WriteConfig()
}

func (c *Apps) ListAllApps() error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	t := tabby.NewCustom(w)
	t.AddHeader("", "App Name", "App ID", "App Certificate", "BaseURL")
	for _, app := range c.Apps {
		def := ""
		if app.AppID == c.Active {
			def = "(active)"
		}
		appCertificate := fmt.Sprintf("**************%v", app.AppCertificate[len(app.AppCertificate)-4:])
		t.AddLine(def, app.Name, app.AppID, appCertificate, app.BaseURL)
	}
	t.Print()
	return nil
}

func (c *Apps) Remove(appID string) error {
	var (
		idx   int
		found bool
	)
	for i, app := range c.Apps {
		if appID == app.AppID {
			found = true
			idx = i
			break
		}
	}
	if !found {
		return fmt.Errorf("[%s]Application doesn't exist", appID)
	}

	if c.Active == appID {
		c.Active = ""
	}

	c.Apps = append(c.Apps[:idx], c.Apps[idx+1:]...)

	viper.Set("active", c.Active)
	viper.Set("apps", c.Apps)
	return viper.WriteConfig()
}

func (c *Apps) Use(appID string) error {
	if c.Active == appID {
		// if already active, early return
		return nil
	}

	var found bool
	for _, app := range c.Apps {
		if appID == app.AppID {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("[%s]Application doesn't exist", appID)
	}
	c.Active = appID
	viper.Set("active", c.Active)
	fmt.Println("Application successfully used as the currently 'active app' ")
	return viper.WriteConfig()
}

func GetApps(cmd *cobra.Command) *Apps {
	apps := &Apps{}
	err := viper.Unmarshal(apps)
	if err != nil {
		cmd.PrintErr("Configuration is malformed: " + err.Error())
		os.Exit(1)
	}

	return apps
}

func GetInitConfig(cmd *cobra.Command, cfgPath *string) func() {
	return func() {
		var configPath string

		if *cfgPath != "" {
			// Use config file from the flag.
			configPath = *cfgPath
		} else {
			// Otherwise use UserConfigDir
			dir, err := os.UserConfigDir()
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}
			configPath = filepath.Join(dir, configDir, "config.yml")
		}

		// fmt.Println("Config Path is :", configPath)
		viper.SetConfigFile(configPath)

		err := viper.ReadInConfig()
		if err != nil && os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(configPath), 0o755)
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}

			f, err := os.Create(configPath)
			if err != nil {
				cmd.PrintErr(err)
				os.Exit(1)
			}

			f.Close()
		}
		if err != nil {
			cmd.PrintErr(err)
			os.Exit(1)
		}
	}
}

func questions() []*survey.Question {
	return []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What is the name of your app? (eg. prod, staging, testing)"},
			Validate: survey.Required,
		},
		{
			Name:     "AppID",
			Prompt:   &survey.Input{Message: "What is your App ID?"},
			Validate: survey.Required,
		},
		{
			Name:     "AppCertificate",
			Prompt:   &survey.Password{Message: "What is your App certificate?"},
			Validate: survey.Required,
		},
		{
			Name: "BaseURL",
			Prompt: &survey.Input{
				Message: "What is your App Base URL?",
			},
			Validate: func(ans interface{}) error {
				u, ok := ans.(string)
				if !ok {
					return errors.New("invalid url")
				}

				_, err := url.ParseRequestURI(u)
				if err != nil {
					return errors.New("invalid url format make sure it matches <scheme>://<host>")
				}
				return nil
			},
		},
	}
}
