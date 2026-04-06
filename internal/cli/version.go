package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/PlatformStackPulse/go-template/internal/logger"
	"github.com/PlatformStackPulse/go-template/pkg/version"
)

// NewVersionCommand creates the version command
func NewVersionCommand(log *logger.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Long:  `Print detailed version information including commit and build time.`,
		Run: func(cmd *cobra.Command, args []string) {
			versionInfo := version.Get()
			fmt.Println(versionInfo.String())
			log.Info("Version command executed", "version", versionInfo.Version)
		},
	}

	return cmd
}
