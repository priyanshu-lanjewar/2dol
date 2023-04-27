/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/back"
	"github.com/spf13/cobra"
)

// backCmd represents the back command
var backCmd = &cobra.Command{
	Use:   "back",
	Short: "Goes Back to List View",
	Aliases: []string{"b","B"},
	Run: func(cmd *cobra.Command, args []string) {
		back.Execute()
	},
}

func init() {
	
}
