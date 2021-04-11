package actions

import (
	"errors"
	"net/url"

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

func PromptSelectEnvironment() (string, error) {
	prompt := promptui.Select{
		Label: "Select a Environment",
		Items: []string{},
	}
	_, result, err := prompt.Run()
	return result, err
}

func ConfiguteAction() error {
	hostURL, err := PromptConfigureHostUrl()
	if err != nil {
		return err
	}
	token, err := PromptConfigureToken()
	if err != nil {
		return err
	}
	environmentId, err := PromptSelectEnvironment()
	if err != nil {
		return err
	}
	return nil
}
