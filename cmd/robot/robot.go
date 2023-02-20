package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"robot/cmd/robot/cmd/tracker"
)

func NewRobotCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:  "robot",
		Long: `The robot system control components.`,
		RunE: func(cmd *cobra.Command, args []string) error {

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

	cmd.AddCommand(tracker.NewTrackerCommand())

	return cmd
}
