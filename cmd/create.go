/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"fmt"

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
		common.ListName = viper.GetString("name")
		fmt.Println(common.ListName)
		create.Execute()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().String("name", "", "Name of TO DO List")
	createCmd.MarkFlagRequired("name")
}
