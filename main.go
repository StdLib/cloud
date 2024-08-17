package main

import (
	"context"
	"os"

	"server/internal"
	"server/internal/environment"
)

func run(
	_ context.Context,
	config internal.ServerConfig,
) error {
	tlsCert := config.TLSCert()
	tlsKey := config.TLSKey()
	server := internal.NewServer(config)
	return server.ListenAndServeTLS(tlsCert, tlsKey)
}

func main() {
	ctx := context.Background()
	err := run(ctx, environment.ServerConfig)
	if err != nil {
		os.Exit(1)
	}
}
