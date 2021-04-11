package models

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
	ApplicationEnvVars []ApplicationEnvVar `json:"application_env_vars"`
}
