/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/enlist"
	"github.com/spf13/cobra"
)

// enlistCmd represents the enlist command
var enlistCmd = &cobra.Command{
	Use:   "enlist",
	Short: "Enlist a task",
	Run: func(cmd *cobra.Command, args []string) {
		enlist.Execute()
	},
}

func init() {
	
}
