package actions

import (
	"errors"
	"fmt"
	"net/url"

	"context"

	"github.com/carlosstrand/manystagings/cli/manystagings/client"
	"github.com/carlosstrand/manystagings/models"
	"github.com/manifoldco/promptui"
)

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

func PromptConfigureToken() (string, error) {
	validate := func(input string) error {
		if len(input) != 32 {
			return errors.New("Invalid token. It must be 32 characters")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Token",
		Validate: validate,
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

func ConfiguteAction() error {
	hostURL, err := PromptConfigureHostUrl()
	if err != nil {
		return err
	}
	client := client.NewClient(hostURL + "/api")
	envs, err := client.GetEnvironments(context.TODO())
	if err != nil {
		return err
	}
	env, err := PromptSelectEnvironment(envs.Data)
	if err != nil {
		return err
	}
	fmt.Println(env)
	return nil
}
