/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/show"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		show.Execute()
	},
}

func init() {
	
}
