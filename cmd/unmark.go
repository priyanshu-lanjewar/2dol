/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// unmarkCmd represents the unmark command
var unmarkCmd = &cobra.Command{
	Use:   "unmark",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("unmark called")
	},
}

func init() {
	rootCmd.AddCommand(unmarkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unmarkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unmarkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
