package orchestrator

import (
	"context"
	"fmt"
	"strings"
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

func NewDeploymentDockerImageFromString(image string) *DeploymentDockerImage {
	parts := strings.Split(image, ":")
	name := parts[0]
	tag := parts[1]
	if tag == "" {
		tag = "latest"
	}
	return &DeploymentDockerImage{
		Name: name,
		Tag:  tag,
	}
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

	// Port to exposte as service
	Port int32

	// ContainerPort is the port that the container will listen
	ContainerPort int32

	// Env is an environment variables map (key -> value)
	Env map[string]string
}

type DeploymentStatus struct {
	// Deployment related to status
	Deployment *Deployment `json:"deployment"`

	// Status. Possible values: PENDING, RUNNING, FAILED, UNKNOWN
	Status string `json:"status"`
}

type PublicURLOptions struct {
	// Host of the URL (eg mysite.com)
	Host string `json:"host"`

	// Subdomain of the URL (eg. my-staging)
	Subdomain string `json:"subdomain"`
}

type Orchestrator interface {
	// Provider name (e.g. kubernetes)
	Provider() string

	// Get the current orchestror settings
	Settings() map[string]interface{}

	// Create a namespace. Every namespace should be unique for environment
	CreateNamespace(ctx context.Context, namespace string) error

	// Create a deployment
	CreateDeployment(ctx context.Context, deployment *Deployment, recreate bool) error

	// Create a public url
	CreatePublicURL(ctx context.Context, deployment *Deployment, opts PublicURLOptions) (string, error)

	// Delete a public url
	DeletePublicURL(ctx context.Context, deployment *Deployment) error

	// Delete a deployment
	DeleteDeployment(ctx context.Context, deployment *Deployment) error

	// Retrieve a list of deployment and their status
	DeploymentStatuses(ctx context.Context, namespace string) ([]DeploymentStatus, error)
}
