package server

// Config provides configuration for the server.
type Config struct {
	Addr     string `json:"addr"`
	Username string `json:"username"`
	Password string `json:"password"`
	TLSCert  string `json:"tls_cert"`
	TLSKey   string `json:"tls_key"`
}
