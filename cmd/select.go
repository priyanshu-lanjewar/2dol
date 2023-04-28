/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	selects "github.com/priyanshu-lanjewar/2dol/pkg/select"
	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select Particular list to work on",
	Aliases: []string{"s","S"},
	Run: func(cmd *cobra.Command, args []string) {
		selects.Execute()
	},
}

func init() {
	
}
