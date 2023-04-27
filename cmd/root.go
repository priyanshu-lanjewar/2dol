/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package cmd

import (
	"os"

	"github.com/priyanshu-lanjewar/2dol/pkg/common"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "2dol",
	Short: "To do list Commandline App",
}

func Execute() {
	common.Init()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
 rootCmd.CompletionOptions.DisableDefaultCmd = true


}


