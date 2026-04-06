// Package cli provides command-line interface commands and initialization.
package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/PlatformStackPulse/go-template/internal/feature"
	"github.com/PlatformStackPulse/go-template/internal/logger"
)

// NewHelloCommand creates the hello command
func NewHelloCommand(log *logger.Logger, fm *feature.Manager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "Print a hello message",
		Long:  `Print a hello message, optionally with a name.`,
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")

			if fm.IsEnabled(feature.FeatureHello) {
				log.Info("Feature flag enabled", "feature", feature.FeatureHello)
			}

			if name != "" {
				fmt.Printf("Hello, %s!\n", name)
				log.Info("Hello command executed", "name", name)
			} else {
				fmt.Println("Hello, World!")
				log.Info("Hello command executed")
			}
		},
	}

	cmd.Flags().StringP("name", "n", "", "Name to greet")

	return cmd
}
