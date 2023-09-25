package main

import (
	"os"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	err := cmd.RunE(cmd, os.Args)
	if err != nil {
		os.Exit(1)
	}
}
