package main

import (
	"context"
	"log"
	"mitra-kirim-be-mgmt/http/rest"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := rest.New()
	if err != nil {
		return err
	}
	err = server.Run(ctx)
	return err
}
