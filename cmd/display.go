/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/display"
	"github.com/spf13/cobra"
)

// displayCmd represents the display command
var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "Displays all the Existing To Do List",
	Aliases: []string{"d","D"},
	Run: func(cmd *cobra.Command, args []string) {
		display.Execute()
	},
}

func init() {

}
