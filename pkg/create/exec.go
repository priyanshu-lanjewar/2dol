package create

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	qs := []*survey.Question{
		{
			Name: "list",
			Prompt: &survey.Input{
				Message: "Enter Name of New List:",
			},
			Validate: survey.Required,
			Transform: survey.Title,
		},
	}

	rs := struct{
		List string `survey:"list"`
	}{}

	common.Error = survey.Ask(qs,&rs)
	common.LogErr(common.Error)
	common.ListName = rs.List
	common.ListFilePath = fmt.Sprintf("%s/.%s",common.UserToDoListDirectoryPath,common.ListName)
	if _, common.Error = os.Stat(common.ListFilePath); os.IsNotExist(common.Error) {
		common.ListFile, common.Error = os.Create(common.ListFilePath) 
		common.LogErr(common.Error)
		common.LogMsg(fmt.Sprintf("List %s created successfully.",common.ListName))
		common.ListFile.Close()
		common.IncConfCount()
		common.UpdateConfListName(common.ListName)
		common.LogMsg(fmt.Sprintf("Selected List %s.",common.ListName))
		common.WriteList(common.Tasks{})
	} else {
		common.LogMsg("A List With Same Name Already Exists.")
	}
}