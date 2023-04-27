package selects

import (
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	c := 0
	ListItems := make([]string,0)
	common.Error = filepath.Walk(common.UserToDoListDirectoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if info.Name() != ".conf" {
				ListItems = append(ListItems, info.Name()[1:])
				c += 1
			}
		}
		return nil
	})

	common.LogErr(common.Error)
	if c == 0 {
		common.LogMsg("No List Exists")
	} else {
		qs := []*survey.Question{{
			Name: "list",
			Prompt: &survey.Select{
				Message: "Select A List From Below Options: ",
				Options: ListItems,
			},
		}}

		response := struct{
			List string `survey:"list"`
		}{}

		common.Error = survey.Ask(qs,&response)
		common.LogErr(common.Error)
		common.UpdateConfListName(response.List)
		common.IsListSelected = true
	}
}