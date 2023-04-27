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
		common.ListName = viper.GetString("name")
		purge.Execute()
	},
}

func init() {
	rootCmd.AddCommand(purgeCmd)
	purgeCmd.Flags().String("name", "", "Name of TO DO List")
	purgeCmd.MarkFlagRequired("name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// purgeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// purgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
