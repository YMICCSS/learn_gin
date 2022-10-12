package models

import (
	"fmt"
	"time"
)

// 時間截記轉換成日期
func UnixToTime(timestamp int) string{
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp),0)
	return t.Format("2006-01-02 15:04:05")
}

func GetDay() string{
	template :="20060102"
	return time.Now().Format(template)
}

func GetUnix() int64{
	return time.Now().Unix()
}


