package models

import "time"

type ApplicationEnvVar struct {
	Base
	ApplicationID string       `json:"application_id"`
	Application   *Application `json:"application"`
	Key           string       `json:"key"`
	Value         string       `json:"value"`
}

type Application struct {
	Base
	EnvironmentID      string              `json:"environment_id"`
	Environment        *Environment        `json:"environment"`
	Name               string              `json:"name"`
	DockerImageName    string              `json:"docker_image_name"`
	DockerImageTag     string              `json:"docker_image_tag"`
	ShellCommand       string              `json:"shell_command"`
	ApplicationEnvVars []ApplicationEnvVar `json:"application_env_vars"`
	Port               int32               `json:"port"`
	ContainerPort      int32               `json:"container_port"`
	PublicUrlEnabled   bool                `json:"public_url_enabled"`
	PublicUrl          string              `json:"public_url"`
	StartedAt          *time.Time          `json:"started_at"`
}
