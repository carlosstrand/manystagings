package actions

import (
	"errors"
	"net/url"

	"context"

	"github.com/carlosstrand/manystagings/cli/ms/client"
	"github.com/carlosstrand/manystagings/cli/ms/utils/msconfig"
	"github.com/carlosstrand/manystagings/models"
	"github.com/manifoldco/promptui"
)

func validateRequiredString(input string) error {
	if len(input) == 0 {
		return errors.New("Field is required")
	}
	return nil
}

func PromptConfigureHostUrl() (string, error) {
	validate := func(input string) error {
		_, err := url.ParseRequestURI(input)
		if err != nil {
			return errors.New("Invalid host url")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Host URL",
		Validate: validate,
	}
	return prompt.Run()
}

func PromptConfigureUsername() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Username",
		Validate: validateRequiredString,
	}
	return prompt.Run()
}

func PromptConfigurePassword() (string, error) {
	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validateRequiredString,
		Mask:     '*',
	}
	return prompt.Run()
}

func PromptSelectEnvironment(envs []*models.Environment) (*models.Environment, error) {
	envOptions := make([]string, 0)
	for _, env := range envs {
		envOptions = append(envOptions, env.Name)
	}
	prompt := promptui.Select{
		Label: "Select an Environment",
		Items: envOptions,
	}
	idx, _, err := prompt.Run()
	return envs[idx], err
}

func (a *Actions) Configure() error {
	hostURL, err := PromptConfigureHostUrl()
	if err != nil {
		return err
	}
	username, err := PromptConfigureUsername()
	if err != nil {
		return err
	}
	password, err := PromptConfigurePassword()
	if err != nil {
		return err
	}
	client := client.NewClient(hostURL)
	token, err := client.Auth(context.TODO(), username, password)
	if err != nil {
		return err
	}
	client.SetAuthToken(token.Value)
	envOptions, err := client.GetEnvironments(context.TODO())
	if err != nil {
		return err
	}
	env, err := PromptSelectEnvironment(envOptions.Data)
	if err != nil {
		return err
	}
	info, err := client.GetInfo(context.TODO())
	if err != nil {
		return err
	}
	config := &msconfig.ManyStagingsConfig{
		HostURL:              hostURL,
		Token:                token.Value,
		LogLevel:             "debug",
		EnvironmentID:        env.ID,
		OrchestratorProvider: info.OrchestratorProvider,
		OrchestratorSettings: info.OrchestratorSettings,
	}
	msconfig.SaveConfig(config)
	return nil
}
