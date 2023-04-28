package mark

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Done() {
	tasks := common.ReadList()

	if len(tasks.ToBeDone) + len(tasks.Doing) == 0 {
		common.LogMsg("Nothing to mark as Done.")
		return
	}

	ListItemInp := append(tasks.ToBeDone, tasks.Doing...)
	ListItemOut := make([]string,0)

	prompt := &survey.MultiSelect{
		Message: "Select Items to mark as Done:",
		Options: ListItemInp,
	}

	survey.AskOne(prompt,&ListItemOut)
	
	for _, item := range ListItemOut{
		for i, task := range tasks.ToBeDone{
			if task == item{
				tasks.ToBeDone = append(tasks.ToBeDone[:i], tasks.ToBeDone[i+1:]...)
				tasks.Done = append(tasks.Done, task)
				break
			}
		}
		for i, task := range tasks.Doing{
			if task == item{
				tasks.Doing = append(tasks.Doing[:i], tasks.Doing[i+1:]...)
				tasks.Done = append(tasks.Done, task)
				break
			}
		}
	}
	common.WriteList(tasks)
	common.LogMsg("Tasks Marked as Done Successfully.")
}