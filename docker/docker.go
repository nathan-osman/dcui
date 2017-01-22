package docker

import (
	"golang.org/x/net/context"

	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
)

// Docker represents a connection to the Docker daemon. This is used for
// starting and stopping containers.
type Docker struct {
	config  *Config
	project project.APIProject
}

// New creates a new connection to the Docker daemon.
func New(config *Config) (*Docker, error) {
	project, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{config.Filename},
			ProjectName:  config.ProjectName,
		},
	}, nil)
	if err != nil {
		return nil, err
	}
	d := &Docker{
		config:  config,
		project: project,
	}
	return d, nil
}

// Status returns the status of all services.
func (d *Docker) Status() (project.InfoSet, error) {
	return d.project.Ps(context.Background())
}
