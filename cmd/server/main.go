package main

import (
	"os"
	"fmt"
	"net/http"

	"github.com/chiyonn/peepa-go/internal/client"
	"github.com/chiyonn/peepa-go/internal/router"
)

func main() {
	pcfg := &client.PeepaConfig{
		Host:         os.Getenv("ERESA_HOST"),
		AuthHost:     os.Getenv("ERESA_AUTH_HOST"),
		ClientID:     os.Getenv("ERESA_CLIENT_ID"),
		RefreshToken: os.Getenv("ERESA_REFRESH_TOKEN"),
	}

	pcli, err := client.NewPeepaClient(pcfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize peepa client: %v\n", err)
		os.Exit(1)
	}

	r := router.NewRouter(pcli)

	http.ListenAndServe(":8080", r)
}
