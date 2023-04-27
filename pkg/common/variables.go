package common

import (
	"os"

	"github.com/apex/log"
)

var UserHomeDirectoryPath string

var UserToDoListDirectoryPath string

var UserToDoListConfigPath string

var UserToDoListConfigFile *os.File

var Error error

var Log log.Logger