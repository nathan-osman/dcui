package docker

import (
	"github.com/docker/libcompose/docker/auth"
	"github.com/docker/libcompose/docker/client"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/docker/network"
	"github.com/docker/libcompose/docker/service"
	"github.com/docker/libcompose/docker/volume"
	"github.com/docker/libcompose/project"
)

// Docker represents a connection to the Docker daemon. This is used for
// starting and stopping containers.
type Docker struct {
	config  *Config
	project *project.Project
}

// New initializes the project.
func New(config *Config) (*Docker, error) {
	c := &ctx.Context{
		Context: project.Context{
			ComposeFiles: []string{config.Filename},
			ProjectName:  config.ProjectName,
		},
	}
	// If docker.NewProject would only return a *Project...
	if c.AuthLookup == nil {
		c.AuthLookup = auth.NewConfigLookup(c.ConfigFile)
	}
	if c.ServiceFactory == nil {
		c.ServiceFactory = service.NewFactory(c)
	}
	if c.ClientFactory == nil {
		f, err := client.NewDefaultFactory(client.Options{})
		if err != nil {
			return nil, err
		}
		c.ClientFactory = f
	}
	if c.NetworksFactory == nil {
		n := &network.DockerFactory{
			ClientFactory: c.ClientFactory,
		}
		c.NetworksFactory = n
	}
	if c.VolumesFactory == nil {
		v := &volume.DockerFactory{
			ClientFactory: c.ClientFactory,
		}
		c.VolumesFactory = v
	}
	// Phew, finally create the project!
	d := &Docker{
		config:  config,
		project: project.NewProject(&c.Context, nil, nil),
	}
	if err := d.project.Parse(); err != nil {
		return nil, err
	}
	return d, nil
}
