/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
	"github.com/priyanshu-lanjewar/2dol/pkg/purge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Deletes the TO-DO list",
	Aliases: []string{"p","P"},
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())
		common.ListName = args[0]
		purge.Execute()
	},
}

func init() {
	
	}
