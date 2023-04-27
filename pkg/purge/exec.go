package purge

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	c :=0
	ListItemInp := make([]string,0)
	ListItemOut := make([]string,0)
	common.Error = filepath.Walk(common.UserToDoListDirectoryPath,func(path string, info os.FileInfo,err error)error{
		if err != nil{
			return err
		}

		if !info.IsDir(){
			if info.Name() != ".conf"{
				ListItemInp= append(ListItemInp,info.Name()[1:])
				c+=1
			}
		}
		return nil
	})

	common.LogErr(common.Error)
	if c==0{
		common.LogMsg("No List Exists")
	} else {
		prompt := &survey.MultiSelect{
			Message: "Select Lists to Delete:",
			Options: ListItemInp,
		}
		survey.AskOne(prompt,&ListItemOut)
	}

	for _, item := range ListItemOut{
		common.ListName = item
		common.ListFilePath = fmt.Sprintf("%s/.%s", common.UserToDoListDirectoryPath, common.ListName)
		if _, common.Error = os.Stat(common.ListFilePath); os.IsNotExist(common.Error) {
			common.LogMsg(fmt.Sprintf("The List with Name %s doesnt Exists.",common.ListName))
		} else {
			common.Error = os.Remove(common.ListFilePath)
			common.LogErr(common.Error)
			common.LogMsg(fmt.Sprintf("List Deleted Successfully: %s",common.ListName))
			common.DecConfCount()
		}
	}
}