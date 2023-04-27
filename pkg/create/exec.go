package create

import (
	"fmt"
	"os"

	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	common.ListFilePath = fmt.Sprintf("%s/.%s",common.UserToDoListDirectoryPath,common.ListName)
	if _, common.Error = os.Stat(common.ListFilePath); os.IsNotExist(common.Error) {
		common.ListFile, common.Error = os.Create(common.ListFilePath) 
		common.LogErr(common.Error)
		common.LogMsg(fmt.Sprintf("List %s created successfully.",common.ListName))
		common.ListFile.Close()
		common.IncConfCount()
		common.UpdateConfListName(common.ListName)
	} else {
		common.LogMsg("A List With Same Name Already Exists.")
	}
}