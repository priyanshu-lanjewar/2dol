/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/mark"
	"github.com/spf13/cobra"
)

// tbdCmd represents the tbd command
var tbdCmd = &cobra.Command{
	Use:   "todo",
	Short: "Mark Existing Task as To-Do",
	Aliases: []string{"tbd","td","t","TODO","TBD","TD","T"},
	Run: func(cmd *cobra.Command, args []string) {
		mark.ToBeDone()
	},
}

func init() {
	markCmd.AddCommand(tbdCmd)
}
