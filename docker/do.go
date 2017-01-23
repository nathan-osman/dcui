package docker

import (
	"errors"

	"github.com/docker/libcompose/project/options"
	"golang.org/x/net/context"
)

// Do applies the specified actions to the specified service. If service is
// empty, the action will be applied to all services.
func (d *Docker) Do(action string, service string) error {
	var (
		services []string
		err      error
	)
	if len(service) != 0 {
		services = append(services, service)
	}
	switch action {
	case "build":
		err = d.project.Build(
			context.Background(),
			options.Build{},
			services...,
		)
	case "create":
		err = d.project.Create(
			context.Background(),
			options.Create{},
			services...,
		)
	case "start":
		err = d.project.Start(
			context.Background(),
			services...,
		)
	case "stop":
		err = d.project.Stop(
			context.Background(),
			10, // timeout in seconds
			services...,
		)
	default:
		err = errors.New("unrecognized action")
	}
	return err
}
