// Package cli provides command-line interface commands and initialization.
package cli

import (
	"github.com/spf13/cobra"

	"github.com/PlatformStackPulse/go-template/internal/feature"
	"github.com/PlatformStackPulse/go-template/internal/logger"
)

// NewRootCommand creates the root command
func NewRootCommand(log *logger.Logger, fm *feature.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "go-template",
		Short: "Production-ready Go template application",
		Long: `A production-ready Go template with CI/CD, DevSecOps,
and architectural best practices built-in.`,
	}

	cmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		log.Debug("Command execution started", "command", cmd.Name())
	}

	cmd.PersistentPostRun = func(cmd *cobra.Command, args []string) {
		log.Debug("Command execution completed", "command", cmd.Name())
	}

	return cmd
}
