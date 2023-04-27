package create

import (
	"fmt"
	"os"

	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	common.ListFilePath = fmt.Sprintf("%s/.%s",common.UserToDoListDirectoryPath,common.ListName)
	fmt.Println(common.ListFilePath)
	if _, common.Error = os.Stat(common.ListFilePath); os.IsNotExist(common.Error) {
		common.ListFile, common.Error = os.Create(common.ListFilePath) 
		common.LogErr(common.Error)
		common.ListFile.Close()
	} else {
		common.LogErr(common.Error)
		common.LogMsg("A List With Same Name Already Exists.")
	}
}