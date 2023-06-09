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

func CleanSlice(Slice []string) []string {
	var CleanedSlice []string
	for _, Element := range Slice {
		if Element != "" {
			CleanedSlice = append(CleanedSlice, Element)
		}
	}
	return CleanedSlice
}

func ReadList() Tasks {
	Data, Error := ioutil.ReadFile(ListFilePath)
	LogErr(Error)
	DecodedData, Error := hex.DecodeString(string(Data))
	LogErr(Error)
	DalaList := strings.Split(string(DecodedData),"|")

	return Tasks{
		ToBeDone: CleanSlice(strings.Split(DalaList[0],",")),
		Doing: CleanSlice(strings.Split(DalaList[1],",")),
		Done:CleanSlice(strings.Split(DalaList[2],",")),
	}
}

func WriteList(Tasks Tasks) {
	Msg := []byte(strings.Join(Tasks.ToBeDone,",")+"|"+strings.Join(Tasks.Doing,",")+"|"+strings.Join(Tasks.Done,","))
	encodedMsg := hex.EncodeToString(Msg)
	Error = ioutil.WriteFile(ListFilePath,[]byte(encodedMsg),0644)
	LogErr(Error)
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