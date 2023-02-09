/*
Copyright © 2023 UlinoyaPed

*/
package cmd

import (
	//"fmt"

	"github.com/UlinoyaPed/gcm/commit"

	//"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Git提交",
	Long:  `Git提交。`,
	Run: func(cmd *cobra.Command, args []string) {
		commit.ExecCommand("git", "commit", "-m "+commit.Header())
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// commitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// commitCmd.Flags().String("toggle", "t", false, "Help message for toggle")
}
