/*
Copyright © 2024 chmikata <chmikata@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-report-cli",
	Short: "gh-report-cli is a CLI tool for manipulating GitHub Issues.",
	Long: `gh-report-cli is a CLI tool for manipulating GitHub Issues.

You can use this tool to report CI results to GitHub issues.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("org", "o", "", "Organization name")
	rootCmd.PersistentFlags().StringP("repo", "r", "", "Repository name")
	rootCmd.PersistentFlags().StringP("token", "t", "", "Token for authentication")
	rootCmd.PersistentFlags().StringP("title", "T", "", "Issue title")
	rootCmd.PersistentFlags().StringP("input", "i", "", "Issue text")
	rootCmd.PersistentFlags().StringP("label", "l", "", "Label to be assigned to the issue")

	rootCmd.MarkPersistentFlagRequired("org")
	rootCmd.MarkPersistentFlagRequired("repo")
	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("title")
	rootCmd.MarkPersistentFlagRequired("input")
}
