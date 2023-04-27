/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
	"github.com/priyanshu-lanjewar/2dol/pkg/create"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates A new To-Do List",
	Aliases: []string{"c","C"},
	
	Run: func(cmd *cobra.Command, args []string) {
		viper.BindPFlags(cmd.Flags())
		common.ListName = args[0]
		create.Execute()
	},
}

func init() {
	
}
