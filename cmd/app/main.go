package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/PlatformStackPulse/go-template/internal/cli"
	"github.com/PlatformStackPulse/go-template/internal/config"
	"github.com/PlatformStackPulse/go-template/internal/feature"
	"github.com/PlatformStackPulse/go-template/internal/logger"
	"github.com/PlatformStackPulse/go-template/pkg/version"
)

func main() {
	// Initialize configuration
	cfg := config.Load()

	// Initialize logger
	log := logger.NewLogger(cfg.Debug)

	// Initialize feature flags
	fm := feature.NewManager()

	// Create root command
	cmd := cli.NewRootCommand(log, fm)

	// Add version command
	cmd.AddCommand(cli.NewVersionCommand(log))

	// Add hello command
	cmd.AddCommand(cli.NewHelloCommand(log, fm))

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Info("Received signal", "signal", sig)
		cancel()
	}()

	// Execute command
	if err := cmd.ExecuteContext(ctx); err != nil {
		log.Error("Command execution failed", "error", err)
		log.Info("Application started successfully", "version", version.Version)
		os.Exit(1) //nolint:gocritic
		return
	}

	log.Info("Application started successfully", "version", version.Version)
}
