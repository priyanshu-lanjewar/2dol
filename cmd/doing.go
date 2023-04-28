/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/mark"
	"github.com/spf13/cobra"
)

// doingCmd represents the doing command
var doingCmd = &cobra.Command{
	Use:   "doing",
	Short: "Mark Existing Task as Doing",
	Run: func(cmd *cobra.Command, args []string) {
		mark.Doing()
	},
}

func init() {
	markCmd.AddCommand(doingCmd)
}
