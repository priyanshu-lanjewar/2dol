/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/mark"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark Existing Task as Done",
	Run: func(cmd *cobra.Command, args []string) {
		mark.Done()
	},
}

func init() {
	markCmd.AddCommand(doneCmd)
}
