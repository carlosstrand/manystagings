package orchestrator

import (
	"context"
	"fmt"
)

type DeploymentDockerImage struct {
	// Docker Image Name (e.g. nginx)
	Name string

	// Docker Image Tag (e.g. latest)
	Tag string
}

func (i *DeploymentDockerImage) ToString() string {
	tag := i.Tag
	if tag == "" {
		tag = "latest"
	}
	return fmt.Sprintf("%s:%s", i.Name, tag)
}

type Deployment struct {
	// Name is a unique deployment name (e.g. node-api)
	Name string

	// Namespace is the unique namespace used by the environment
	Namespace string

	/* DockerImage is the docker imaged used in deployment
	Example:
	DeploymentDockerImage{Name: "nginx", Tag: "latest"}
	*/
	DockerImage DeploymentDockerImage

	// Env is an environment variables map (key -> value)
	Env map[string]string
}

type Orchestrator interface {
	// Create a namespace. Every namespace should be unique for environment
	CreateNamespace(ctx context.Context, namespace string) error

	// Create a deployment
	CreateDeployment(ctx context.Context, deployment *Deployment) error

	// Delete a deployment
	DeleteDeployment(ctx context.Context, deployment *Deployment) error
}
