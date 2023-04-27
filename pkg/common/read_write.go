package common

import (
	"encoding/hex"
	"io/ioutil"
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