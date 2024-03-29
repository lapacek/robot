package main

import (
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
	}

	cmd.AddCommand(tracker.NewTrackerCommand())

	return cmd
}
