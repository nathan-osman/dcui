package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nathan-osman/dcui/docker"
	"github.com/nathan-osman/dcui/server"
	"github.com/nathan-osman/go-multiarg"
)

// Config provides storage for the application.
type Config struct {
	Docker *docker.Config
	Server *server.Config
}

func main() {

	// Default configuration
	config := &Config{
		Docker: &docker.Config{},
		Server: &server.Config{
			Addr: ":8000",
		},
	}

	// Parse arguments
	if ok, _ := multiarg.Load(config, &multiarg.Config{}); !ok {
		os.Exit(1)
	}

	// Start the server
	s, err := server.New(config.Server)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// Wait for SIGINT or SIGTERM
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
}
