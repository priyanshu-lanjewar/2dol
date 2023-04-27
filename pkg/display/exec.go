package display

import (
	"os"
	"path/filepath"

	"github.com/priyanshu-lanjewar/2dol/pkg/common"
)

func Execute() {
	c :=0
	common.Error = filepath.Walk(common.UserToDoListDirectoryPath,func(path string, info os.FileInfo,err error)error{
		if err != nil{
			return err
		}

		if !info.IsDir(){
			if info.Name() != ".conf"{
				common.LogMsg(info.Name()[1:])
				c+=1
			}
		}
		return nil
	})

	common.LogErr(common.Error)
	if c==0{
		common.LogMsg("No List Exists")
	}
}