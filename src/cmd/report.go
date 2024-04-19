/*
Copyright © 2024 chmikata <chmikata@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Register an issue.",
	Long:  "Register report information in the issue.",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		input, _ := rootCmd.PersistentFlags().GetString("input")
		_, err := os.Stat(input)
		if os.IsNotExist(err) {
			return fmt.Errorf("input file does not exist: %s", input)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("report called")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
