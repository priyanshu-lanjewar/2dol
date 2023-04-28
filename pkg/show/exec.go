package show

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/list"
	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func PrinlList(content string) {
	for _, line := range strings.Split(content, "\n") {
		fmt.Printf("%s\n", line)
	}
	fmt.Println()
}

func ParseAsInterfaceList(tasks []string) []interface{} {
	iList := []interface{}{}

	if len(tasks) == 0 {
		iList = append(iList, "No Tasks")
		return iList
	}

	for _, task := range tasks {
		if task != "" {
			iList = append(iList, task)
		}
	}
	
	return iList
}

func Execute() {
	tasks := common.ReadList()
	l := list.NewWriter()
	l.SetStyle(list.StyleConnectedRounded)
	l.AppendItem(fmt.Sprintf("List: %s", common.ListName))
	l.Indent()
	l.AppendItem("To Be Done")
	l.Indent()
	l.AppendItems(ParseAsInterfaceList(tasks.ToBeDone))
	l.UnIndent()
	l.AppendItem("Doing")
	l.Indent()
	l.AppendItems(ParseAsInterfaceList(tasks.Doing))
	l.UnIndent()
	l.AppendItem("Done")
	l.Indent()
	l.AppendItems(ParseAsInterfaceList(tasks.Done))
	PrinlList( l.Render())

}