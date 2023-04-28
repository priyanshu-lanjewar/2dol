package mark

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func ToBeDone() {
	tasks := common.ReadList()
	if len(tasks.Doing) + len(tasks.Done) == 0 {
		common.LogMsg("Nothing to mark as To Be Done.")
		return
	}

	ListItemInp := append(tasks.Doing, tasks.Done...)
	ListItemOut := make([]string,0)

	prompt := &survey.MultiSelect{
		Message: "Select Items To mark as To-Do:",
		Options: ListItemInp,
	}
	survey.AskOne(prompt,&ListItemOut)

	for _, item := range ListItemOut{
		for i, task := range tasks.Doing{
			if task == item{
				tasks.Doing = append(tasks.Doing[:i], tasks.Doing[i+1:]...)
				tasks.ToBeDone = append(tasks.ToBeDone, task)
				break
			}
		}
		for i, task := range tasks.Done{
			if task == item{
				tasks.Done = append(tasks.Done[:i], tasks.Done[i+1:]...)
				tasks.ToBeDone = append(tasks.ToBeDone, task)
				break
			}
		}
	}
	common.WriteList(tasks)
	common.LogMsg("Tasks Marked as To Be Done Successfully.")
}