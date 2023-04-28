package enlist

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	
	qs := []*survey.Question{
		{
			Name: "task",
			Prompt: &survey.Input{
				Message: "Enter Name of Task:",
			},
			Validate: survey.Required,
			Transform: survey.Title,
		},
	}

	rs := struct{
		Task string `survey:"task"`
	}{}

	common.Error = survey.Ask(qs,&rs)
	common.LogErr(common.Error)

	tasks := common.ReadList()
	tasks.ToBeDone = append(tasks.ToBeDone, rs.Task)
	common.WriteList(tasks)
	common.LogMsg("Task Enlisted Successfully.")
}