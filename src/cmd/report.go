/*
Copyright © 2024 chmikata <chmikata@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/chmikata/gh-report-cli/internal/application"
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
		org, _ := rootCmd.PersistentFlags().GetString("org")
		repo, _ := rootCmd.PersistentFlags().GetString("repo")
		token, _ := rootCmd.PersistentFlags().GetString("token")
		title, _ := rootCmd.PersistentFlags().GetString("title")
		input, _ := rootCmd.PersistentFlags().GetString("input")
		label, _ := rootCmd.PersistentFlags().GetString("label")
		labels := strings.Split(label, ",")
		search, _ := rootCmd.PersistentFlags().GetString("search")
		reporter := application.NewReporter(org, repo, token)
		issue, err := reporter.Report(title, input, search, labels)
		if err != nil {
			return err
		}
		v, err := json.Marshal(issue)
		if err != nil {
			return err
		}
		fmt.Println(string(v))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(reportCmd)
}
