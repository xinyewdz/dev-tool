package tool

import (
	"fmt"
	"time"
)

func Unix2Date(timestamp int64) (timeStr string) {
	if timestamp == 0 {
		timestamp = time.Now().Unix()
	}
	var dateTime = time.Unix(timestamp, 0)
	timeStr = dateTime.Format("2006-01-02 15:04:05")
	return timeStr
}

func Date2Unix(date string) (timestamp int64) {
	local, _ := time.LoadLocation("Local")
	timeObj, _ := time.ParseInLocation("2006-01-02 15:04:05", date, local)
	timestamp = timeObj.Unix()
	return timestamp
}

func PrintUnix() {
	fmt.Println(time.Now().Unix())
}
