package common

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

func Init() {

	Log = log.Logger{
		Handler: cli.New(os.Stdout),
		Level: 1,
	}

	UserHomeDirectoryPath, Error= os.UserHomeDir()
	LogErr(Error)

	UserToDoListDirectoryPath = fmt.Sprintf("%s/.2dol",UserHomeDirectoryPath)
	UserToDoListConfigPath = fmt.Sprintf("%s/.conf",UserToDoListDirectoryPath)
	
	_, Error = os.Stat(UserToDoListDirectoryPath)
	if ok := os.IsNotExist(Error); ok {
		Error = os.Mkdir(UserToDoListDirectoryPath,0700)
		LogErr(Error)
		UserToDoListConfigFile,Error = os.Create(UserToDoListConfigPath)
		LogErr(Error)
		WriteConf("0,")
		UserToDoListConfigFile.Close()
	} else {
		UserToDoListConfigFile, Error = os.OpenFile(UserToDoListConfigPath,os.O_RDWR,0777)
		UserToDoListConfigFile.Close()
	}
}