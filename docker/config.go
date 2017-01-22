package docker

// Config provides configuration for connecting to Docker.
type Config struct {
	Filename    string `multiarg:"compose filename"`
	ProjectName string `multiarg:"project name"`
}
