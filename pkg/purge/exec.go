package purge

import (
	"fmt"
	"os"

	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	common.ListFilePath = fmt.Sprintf("%s/.%s", common.UserToDoListDirectoryPath, common.ListName)
	if _, common.Error = os.Stat(common.ListFilePath); os.IsNotExist(common.Error) {
		common.LogMsg(fmt.Sprintf("The List with Name %s doesnt Exists.",common.ListName))
	} else {
		common.Error = os.Remove(common.ListFilePath)
		common.LogErr(common.Error)
		common.LogMsg("List Deleted Successfully")
		common.DecConfCount()
	}
}