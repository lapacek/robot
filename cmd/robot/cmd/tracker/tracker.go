package tracker

import (
	"fmt"

	"github.com/spf13/cobra"

	"robot/cmd/robot/cmd/tracker/app"
)

func NewTrackerCommand() *cobra.Command {

	tracker := app.NewManager("Tracker Manager")

	cmd := &cobra.Command{
		Use:  "tracker",
		Long: `The Tracker controller.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			tracker.Run()
			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	return cmd
}
