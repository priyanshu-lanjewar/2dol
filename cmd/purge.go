/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/purge"
	"github.com/spf13/cobra"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Deletes the TO-DO list",
	Aliases: []string{"p","P"},
	Run: func(cmd *cobra.Command, args []string) {
		purge.Execute()
	},
}

func init() {
	
}
