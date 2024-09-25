package agora_chat

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cheynewallace/tabby"
	"github.com/spf13/viper"
	auth "github.com/ycj3/agora-chat-cli/agora-chat/auth"
)

func (apps *Apps) RunQuestionnaire() error {
	var newApp App
	err := survey.Ask(questions(), &newApp)
	if err != nil {
		return err
	}
	err = apps.Add(newApp)
	if err != nil {
		return err
	}
	fmt.Println("Application successfully added. 🚀")
	return nil
}

func (apps *Apps) Add(newApp App) error {
	if len(apps.Apps) == 0 {
		apps.Active = newApp.AppID
	}

	for _, app := range apps.Apps {
		if newApp.Name == app.Name {
			return fmt.Errorf("[%s]Application name already exists", newApp.Name)
		}
		if !auth.HasEnvToken() {
			if newApp.AppID == app.AppID {
				return fmt.Errorf("[%s]Application appID already exists", newApp.AppID)
			}
		}
	}

	apps.Apps = append(apps.Apps, newApp)

	viper.Set("active", apps.Active)
	viper.Set("apps", apps.Apps)
	return viper.WriteConfig()
}

func (apps *Apps) ListAllApps() error {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	t := tabby.NewCustom(w)
	t.AddHeader("", "App Name", "App ID", "App Certificate", "BaseURL")
	for _, app := range apps.Apps {
		def := ""

		if app.Name == apps.Active {
			def = "(active)"
		}

		appID := app.AppID
		if appID == "" {
			appID = "--"
		}

		appCert := app.AppCertificate
		if appCert == "" {
			appCert = "--"
		} else {
			appCert = fmt.Sprintf("**************%v", app.AppCertificate[len(app.AppCertificate)-4:])
		}

		t.AddLine(def, app.Name, appID, appCert, app.BaseURL)
	}
	t.Print()
	return nil
}

func (apps *Apps) Remove(appName string) error {
	var (
		idx   int
		found bool
	)
	for i, app := range apps.Apps {
		if appName == app.Name {
			found = true
			idx = i
			break
		}
	}
	if !found {
		return fmt.Errorf("[%s]Application doesn't exist", appName)
	}

	if apps.Active == appName {
		apps.Active = ""
	}

	apps.Apps = append(apps.Apps[:idx], apps.Apps[idx+1:]...)

	viper.Set("active", apps.Active)
	viper.Set("apps", apps.Apps)
	return viper.WriteConfig()
}

func (apps *Apps) Use(appName string) error {
	if apps.Active == appName {
		// if already active, early return
		return nil
	}

	var found bool
	for _, app := range apps.Apps {
		if appName == app.Name {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("[%s]Application doesn't exist", appName)
	}
	apps.Active = appName
	viper.Set("active", apps.Active)
	return viper.WriteConfig()
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
				Message: "What is your App Base URL? (e.g., https://a61.chat.agora.io/61717166/1069763)",
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
