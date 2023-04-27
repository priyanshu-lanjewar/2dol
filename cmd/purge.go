/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Deletes the TO-DO list",
	Aliases: []string{"p","P"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("purge called")
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// purgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// purgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
