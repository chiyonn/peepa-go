package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/chiyonn/peepa-go/internal/core"
	"github.com/chiyonn/peepa-go/internal/client"
	"github.com/chiyonn/peepa-go/internal/router"
)

func main() {
	pcfg := &client.PeepaConfig{
		Host: core.MustReadSecret("ERESA_HOST"),
		AuthHost: core.MustReadSecret("ERESA_AUTH_HOST"),
		ClientID: core.MustReadSecret("ERESA_CLIENT_ID"),
		RefreshToken: core.MustReadSecret("ERESA_CLIENT_SECRET"),
	}

	pcli, err := client.NewPeepaClient(pcfg)
	if err != nil {
		fmt.Errorf("failed to inizialize peepa client: %w", err)
		os.Exit(0)
	}

	r := router.NewRouter(pcli)

	http.ListenAndServe(":8080", r)
}
