package services

import (
	"fmt"
	"logger/internal/redis"
	"logger/internal/telegram"
	logmodel "logger/src/models/log_model"
	"time"
)

func StartLoggerService() {
	ticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for {
			<-ticker.C
			for i := 0; i < 5; i++ {
				go log()
			}
		}
	}()
}

func log() {
	data := redis.Rpop("logs")
	l := logmodel.ReportableLog{
		ENV:     data["env"],
		Message: data["message"],
		File:    data["file"],
		Line:    data["line"],
		Phone:   data["phone"],
		Ip:      data["ip"],
		Url:     data["url"],
		Time:    data["time"],
	}
	if data != nil {
		telegram.ReportError(l)
	}
	fmt.Println("checking redis")
}
