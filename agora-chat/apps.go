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
)

// type Apps interface {
// 	SomeMethod(x int64, y string)
// }

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
		if newApp.AppID == app.AppID {
			return fmt.Errorf("[%s]Application already exists", newApp.AppID)
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
		if app.AppID == apps.Active {
			def = "(active)"
		}
		appCertificate := fmt.Sprintf("**************%v", app.AppCertificate[len(app.AppCertificate)-4:])
		t.AddLine(def, app.Name, app.AppID, appCertificate, app.BaseURL)
	}
	t.Print()
	return nil
}

func (apps *Apps) Remove(appID string) error {
	var (
		idx   int
		found bool
	)
	for i, app := range apps.Apps {
		if appID == app.AppID {
			found = true
			idx = i
			break
		}
	}
	if !found {
		return fmt.Errorf("[%s]Application doesn't exist", appID)
	}

	if apps.Active == appID {
		apps.Active = ""
	}

	apps.Apps = append(apps.Apps[:idx], apps.Apps[idx+1:]...)

	viper.Set("active", apps.Active)
	viper.Set("apps", apps.Apps)
	return viper.WriteConfig()
}

func (apps *Apps) Use(appID string) error {
	if apps.Active == appID {
		// if already active, early return
		return nil
	}

	var found bool
	for _, app := range apps.Apps {
		if appID == app.AppID {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("[%s]Application doesn't exist", appID)
	}
	apps.Active = appID
	viper.Set("active", apps.Active)
	fmt.Println("Application successfully used as the currently 'active app' ")
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
