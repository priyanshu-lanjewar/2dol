/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/
package common

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadConf() string {
	
	Msg, Error := ioutil.ReadFile(UserToDoListConfigPath)
	LogErr(Error)
	decodedMsg, Error := hex.DecodeString(string(Msg))
	LogErr(Error)
	return string(decodedMsg)
}

func WriteConf(Message string) {
	Msg := []byte(Message)
	encodedMsg := hex.EncodeToString(Msg)
	Error = ioutil.WriteFile(UserToDoListConfigPath,[]byte(encodedMsg),0644)
	LogErr(Error)
}

func UpdateConfListName(List string){
	FileData := ReadConf()
	DalaList := strings.Split(FileData,",")
	DalaList[1] = List
	FileData = strings.Join(DalaList,",")
	WriteConf(FileData)
}

func IncConfCount() {
	FileData := ReadConf()
	DalaList := strings.Split(FileData,",")
	i,_ := strconv.Atoi(DalaList[0])
	DalaList[0] = fmt.Sprintf("%d",i+1)
	FileData = strings.Join(DalaList,",")
	WriteConf(FileData)
}

func DecConfCount() {
	FileData := ReadConf()
	DalaList := strings.Split(FileData,",")
	i,_ := strconv.Atoi(DalaList[0])
	DalaList[0] = fmt.Sprintf("%d",i-1)
	FileData = strings.Join(DalaList,",")
	WriteConf(FileData)
	if DalaList[0] == "0"{
		UpdateConfListName("")
	}
}