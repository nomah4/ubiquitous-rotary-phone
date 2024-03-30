package cmd

import (
	"context"
	"log"

	"github.com/GoSeoTaxi/dh8_msg/internal/client"
	"github.com/GoSeoTaxi/dh8_msg/internal/config"
	"github.com/fatih/color"
)

func RunService(ctx context.Context) error {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("parcing config err = %v", err)
	}

	log.Printf(color.HiRedString("Starting version %v"), cfg.AppVersion)

	return client.StartWorker(ctx, cfg)
}
