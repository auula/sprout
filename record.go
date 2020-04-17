// Copyright (c) 2020 HigKer
// Open Source: MIT License
// Author: SDing <deen.job@qq.com>
// Date: 2020/4/16 - 10:59 下午

package logker

import "github.com/fatih/color"

const (
	//[INFO] 2006-01-02 13:05.0006 MP - Position: test.go|main.test:21 - Message: news
	FORMAT = "[%s] - Date: %s  %s - Message: %s"
)

// Logging record
type LogRecord interface {
	OutPutMessage(v string)
}

func (c *console) OutPutMessage(model Level, v string) {
	switch model.ToStr() {
	case DEBUG.ToStr():
		// blue color of log message.
		// format log message output console.
		color.Blue(FORMAT, DEBUG.ToStr(), c.tz.NowTimeStr(), BuildCallerStr(SKIP), v)
	case INFO.ToStr():
		color.Green(FORMAT, INFO.ToStr(), c.tz.NowTimeStr(), BuildCallerStr(SKIP), v)
	case WARNING.ToStr():
		color.Yellow(FORMAT, WARNING.ToStr(), c.tz.NowTimeStr(), BuildCallerStr(SKIP), v)
	case ERROR.ToStr():
		color.Red(FORMAT, ERROR.ToStr(), c.tz.NowTimeStr(), BuildCallerStr(SKIP), v)
	default:
		// Log Level Type Error
		// Program automatically set to debug
		color.Red("-----------------------------------------------------------------")
		color.Red("！！！Log Level Type Error,Program automatically set to debug！！！")
		color.Red("-----------------------------------------------------------------")
		c.logLevel = DEBUG
		// recursion
		c.OutPutMessage(model, v)
	}
}
