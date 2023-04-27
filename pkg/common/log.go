/*
Copyright © 2023 PRIYANSHU LANJEWAR @ PRIYANSHU.LANJEWAR@YAHOO.COM
*/package common

func LogErr(Error error) {
	if Error != nil {
		Log.Fatal(Error.Error())
	}
}

func LogMsg(Message string) {
	Log.Info(Message)
}