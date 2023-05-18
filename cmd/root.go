/**
 * @Author: kwens
 * @Date: 2023-05-18 13:23:02
 * @Description:
 */
package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:                   "rancher-tool",
	Short:                 "rancher-tool is a tool for rancher api",
	Long:                  "rancher-tool is a tool for rancher api",
	DisableFlagsInUseLine: true,
	DisableAutoGenTag:     true,
	DisableSuggestions:    true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(UpdateCmd)
}
