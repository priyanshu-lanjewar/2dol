/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Displays all the Existing To Do List",
	Aliases: []string{"d","D"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("display called")
	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
