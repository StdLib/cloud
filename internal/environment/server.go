package environment

type serverConfig struct{}

var ServerConfig *serverConfig

func (*serverConfig) Host() string {
	return "0.0.0.0"
}

func (*serverConfig) Port() string {
	return "10443"
}

func (*serverConfig) TLSCert() string {
	return "server.crt"
}

func (*serverConfig) TLSKey() string {
	return "server.key"
}
