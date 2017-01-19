package main

import (
	"os"

	"github.com/nathan-osman/dcui/server"
	"github.com/nathan-osman/go-multiarg"
)

// Config provides storage for the application.
type Config struct {
	Server *server.Config
}

func main() {
	config := &Config{
		Server: &server.Config{
			Addr: ":8000",
		},
	}
	if ok, _ := multiarg.Load(config, &multiarg.Config{}); !ok {
		os.Exit(1)
	}
}
