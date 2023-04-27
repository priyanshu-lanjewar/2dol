/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select Particular list to work on",
	Aliases: []string{"s","S"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("select called")
	},
}

func init() {
	rootCmd.AddCommand(selectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// selectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// selectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
