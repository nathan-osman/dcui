package server

// Config provides configuration for the server.
type Config struct {
	Addr     string `multiarg:"server address"`
	Username string `multiarg:"HTTP basic auth username"`
	Password string `multiarg:"HTTP basic auth password"`
	TLSCert  string `multiarg:"TLS certificate"`
	TLSKey   string `multiarg:"TLS key"`
}
