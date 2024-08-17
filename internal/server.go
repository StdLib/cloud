package internal

import (
	"net"
	"net/http"

	"server/internal/route"
)

type ServerConfig interface {
	Host() string
	Port() string
	TLSCert() string
	TLSKey() string
}

func NewServer(config ServerConfig) *http.Server {
	addr := net.JoinHostPort(config.Host(), config.Port())
	handler := route.Router.Mux()
	return &http.Server{Addr: addr, Handler: handler}
}
