/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// backCmd represents the back command
var backCmd = &cobra.Command{
	Use:   "back",
	Short: "Goes Back to List View",
	Aliases: []string{"b","B"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("back called")
	},
}

func init() {
	rootCmd.AddCommand(backCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// backCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// backCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
