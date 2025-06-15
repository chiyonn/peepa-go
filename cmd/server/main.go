package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/chiyonn/peepa-go/internal/client"
	"github.com/chiyonn/peepa-go/internal/router"
	"github.com/chiyonn/peepa-go/internal/service"
)

func main() {
	pcfg := &client.PeepaConfig{
		Host:         os.Getenv("ERESA_HOST"),
		AuthHost:     os.Getenv("ERESA_AUTH_HOST"),
		ClientID:     os.Getenv("ERESA_CLIENT_ID"),
		RefreshToken: os.Getenv("ERESA_REFRESH_TOKEN"),
	}

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	logger := slog.New(handler)

	pcli, err := client.NewPeepaClient(pcfg, logger)
	if err != nil {
		logger.Error("failed to initialize peepa client", slog.Any("error", err))
		os.Exit(1)
	}

	psrv := service.NewProductService(pcli, logger)

	r := router.NewRouter(psrv, logger)

	logger.Info("Starting server", slog.String("addr", ":8080"))
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("Server failed", slog.Any("error", err))
		os.Exit(1)
	}
}
