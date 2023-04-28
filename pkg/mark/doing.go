package mark

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Doing() {
	tasks := common.ReadList()

	if len(tasks.ToBeDone) + len(tasks.Done) == 0 {
		common.LogMsg("Nothing to mark as Doing.")
		return
	}

	ListItemInp := append(tasks.ToBeDone, tasks.Done...)
	ListItemOut := make([]string,0)

	prompt := &survey.MultiSelect{
		Message: "Select Items to mark as Doing:",
		Options: ListItemInp,
	}
	
	survey.AskOne(prompt,&ListItemOut)

	for _, item := range ListItemOut{
		for i, task := range tasks.ToBeDone{
			if task == item{
				tasks.ToBeDone = append(tasks.ToBeDone[:i], tasks.ToBeDone[i+1:]...)
				tasks.Doing = append(tasks.Doing, task)
				break
			}
		}
		for i, task := range tasks.Done{
			if task == item{
				tasks.Done = append(tasks.Done[:i], tasks.Done[i+1:]...)
				tasks.Doing = append(tasks.Doing, task)
				break
			}
		}
	}
	common.WriteList(tasks)
	common.LogMsg("Tasks Marked as Doing Successfully.")
}