/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package common

import (
	"os"

	"github.com/apex/log"
)

var UserHomeDirectoryPath string

var UserToDoListDirectoryPath string

var UserToDoListConfigPath string

var UserToDoListConfigFile *os.File

var ListName string

var ListFilePath string

var ListFile *os.File

var IsListSelected bool

var Error error

var Log log.Logger