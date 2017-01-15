package config

import (
	"github.com/nathan-osman/dcui/server"
)

// Config provides storage for the DCUI configuration.
type Config struct {
	Server server.Config `json:"server"`
}
